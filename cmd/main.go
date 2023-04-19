/*
Package main contains the main function for the application. It is responsible for starting the server and handling
requests. It also extracts the port from the environment variables and sets it to the default port if it is not set.
*/
package main

import (
	"assignment-2/constants"
	"log"
	"net/http"
	"os"
)

func main() {

	// Extract port from env
	port := os.Getenv("PORT")

	// Set port to default if it has not been set
	if port == "" {
		log.Println("$PORT has not been set. Default: " + constants.DefaultPort + " will be used.")
		port = constants.DefaultPort
	}

	// TODO: Add handlers
	// E.g., http.HandleFunc("/path", handlerpackage.handler)

	// Start http server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
