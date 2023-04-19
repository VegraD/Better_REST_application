/*
Package main contains the main function for the application. It is responsible for starting the server and handling
requests. It also extracts the port from the environment variables and sets it to the default port if it is not set.
*/
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// Extract port from env
	port := os.Getenv("PORT")

	// Set to default if not set
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// TODO: Add handlers
	// http.HandleFunc("/path", handlerpackage.handler)

	// Start http server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
