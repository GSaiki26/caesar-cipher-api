package main

import (
	"log"
	"net/http"

	"github.com/GSaiki26/go-caesar-cipher/internal/routes"
)

// Data
const ADDR = "0.0.0.0:8080"

// Functions
func main() {
	// Configure the server.
	mux := http.NewServeMux()
	mux.HandleFunc("POST /encode", routes.PostEncode)

	// Start the server.
	log.Println("Server started on:", ADDR)
	err := http.ListenAndServe(ADDR, mux)
	if err != nil {
		panic(err)
	}

}
