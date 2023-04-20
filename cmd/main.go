/*
Package main contains the main function for the application. It is responsible for starting the server and handling
requests. It also extracts the port from the environment variables and sets it to the default port if it is not set.
*/
package main

import (
	"assignment-2/constants"
	"assignment-2/handlers"
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

	// Register handlers
	http.HandleFunc("/", handlers.DefaultHandler)
	http.HandleFunc(constants.StatusEP, handlers.StatusHandler)

	// Start http server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
