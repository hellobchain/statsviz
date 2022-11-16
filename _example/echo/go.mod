module example/echo

go 1.16

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hellobchain/statsviz v0.5.1
	github.com/labstack/echo/v4 v4.9.0
	github.com/mattn/go-colorable v0.1.13 // indirect
)

replace github.com/hellobchain/statsviz v0.5.1 => ../../../statsviz
