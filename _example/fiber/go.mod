module example/fiber

go 1.16

require (
	github.com/gofiber/adaptor/v2 v2.1.27
	github.com/gofiber/fiber/v2 v2.38.1
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hellobchain/statsviz v0.5.1
	github.com/soheilhy/cmux v0.1.5
)

replace github.com/hellobchain/statsviz v0.5.1 => ../../../statsviz
