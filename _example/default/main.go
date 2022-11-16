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

	// Register statsviz handlers on the default serve mux.
	err := statsviz.RegisterDefault()
	if err != nil {
		return
	}

	fmt.Println("Point your browser to http://localhost:8080/debug/statsviz/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
