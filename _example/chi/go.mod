module example/chi

go 1.16

require (
	github.com/go-chi/chi v1.5.4
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hellobchain/statsviz v0.5.1
)

replace github.com/hellobchain/statsviz v0.5.1 => ../../../statsviz
