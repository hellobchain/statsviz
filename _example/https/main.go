package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hellobchain/statsviz"
	example "github.com/hellobchain/statsviz/_example"
)

func main() {
	// Force the GC to work to make the plots "move".
	go example.Work()

	const (
		// Use your own certificates and key files.
		certFile = "./cert.pem"
		keyFile  = "./key.pem"
	)

	// Create a serve mux and register statsviz handlers.
	mux := http.NewServeMux()
	if err := statsviz.Register(mux); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Point your browser to https://localhost:8087/debug/statsviz/")
	log.Fatal(http.ListenAndServeTLS(":8087", certFile, keyFile, mux))
}
