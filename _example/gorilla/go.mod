module example/gorilla

go 1.16

require (
	github.com/hellobchain/statsviz v0.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0 // indirect
)

replace (
	github.com/hellobchain/statsviz v0.5.1 => ../../../statsviz
)
