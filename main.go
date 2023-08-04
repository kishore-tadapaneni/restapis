package main

import (
	"log"
	"net/http"
	"restapi/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
