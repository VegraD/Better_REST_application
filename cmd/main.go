/*
Package main contains the main function for the application. It is responsible for starting the server and handling
requests. It also extracts the port from the environment variables and sets it to the default port if it is not set.
*/
package main

import (
	"assignment-2/constants"
	"assignment-2/handlers/currentHandler"
	"assignment-2/handlers/defaultHandler"
	"assignment-2/handlers/historicalHandler"
	"assignment-2/handlers/statusHandler"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Extract port from env
	port := os.Getenv("PORT")

	// Set port to default if it has not been set
	if port == "" {
		fmt.Println()
		log.Println("$PORT has not been set. Default port: " + constants.DefaultPort + " will be used.")
		port = constants.DefaultPort
	}

	// Register handlers
	http.HandleFunc(constants.DefaultEP, defaultHandler.DefaultHandler)
	http.HandleFunc(constants.CurrentEP, currentHandler.CurrentHandler)
	http.HandleFunc(constants.HistoryEP, historicalHandler.HistoricalHandler)
	http.HandleFunc(constants.StatusEP, statusHandler.StatusHandler)

	// Start http server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
