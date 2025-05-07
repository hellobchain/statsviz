package statsviz

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hellobchain/statsviz/internal/static"
)

// RegisterDefault registers statsviz HTTP handlers on the default serve mux.
//
// Note this is not advised on a production server, unless it only serves on
// localhost.
func RegisterDefault(opts ...OptionFunc) error {
	return Register(http.DefaultServeMux, opts...)
}

// Root sets the root path of statsviz handlers.
func Root(root string) OptionFunc {
	return func(s *server) error {
		s.root = root
		return nil
	}
}

// SendFrequency defines the frequency at which statistics are sent from the
// application to the HTML page.
func SendFrequency(freq time.Duration) OptionFunc {
	return func(s *server) error {
		if freq <= 0 {
			return fmt.Errorf("frequency must be a positive integer")
		}
		s.freq = freq
		return nil
	}
}

func Fs(fs embed.FS) OptionFunc {
	return func(s *server) error {
		s.fs = fs
		return nil
	}
}

// An OptionFunc is a server configuration option.
type OptionFunc func(s *server) error

const (
	defaultRoot          = "/debug/statsviz"
	defaultSendFrequency = time.Second
)

// Register registers statsviz HTTP handlers on the provided mux.
func Register(mux *http.ServeMux, opts ...OptionFunc) error {
	s := &server{
		mux:  mux,
		root: defaultRoot,
		freq: defaultSendFrequency,
		fs:   static.Assets,
	}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return err
		}
	}

	s.register()
	return nil
}

type server struct {
	mux  *http.ServeMux
	freq time.Duration
	root string
	fs   embed.FS
}

func (s *server) register() {
	s.mux.Handle(s.root+"/", IndexAtRootWithFs(s.root, s.fs))
	s.mux.HandleFunc(s.root+"/ws", NewWsHandler(s.freq))
}

func RouteRegister(rg gin.IRouter, opts ...OptionFunc) error {
	s := &server{
		root: defaultRoot,
		freq: defaultSendFrequency,
		fs:   static.Assets,
	}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return err
		}
	}
	prefixRouter := rg.Group(s.root)
	{
		prefixRouter.GET("/", gin.WrapF(IndexAtRootWithFs(s.root, s.fs)))
		prefixRouter.GET("/css/*any", gin.WrapF(IndexAtRootWithFs(s.root, s.fs)))
		prefixRouter.GET("/js/*any", gin.WrapF(IndexAtRootWithFs(s.root, s.fs)))
		prefixRouter.GET("/libs/*any", gin.WrapF(IndexAtRootWithFs(s.root, s.fs)))
		prefixRouter.GET("/index.html", gin.WrapF(IndexAtRootWithFs(s.root, s.fs)))
		prefixRouter.GET("/ws", gin.WrapH(NewWsHandler(s.freq)))
	}
	return nil
}
