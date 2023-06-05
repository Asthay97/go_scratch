package main

import (
	"log"
	"net/http"

	"github.com/Asthay27/go_scratch/src/router"
)

func main() {

	router := router.SetupRouter()

	// Start the server
	log.Println("Server started on http://localhost:3000")
	// Start the server
	log.Fatal(http.ListenAndServe(":3000", router))
}
