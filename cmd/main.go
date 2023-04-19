package cmd

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// extract port from env
	port := os.Getenv("PORT")

	// set to default if not set
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// TODO: Add handlers
	// http.HandleFunc("/", handlerpackage.handler)

	// start http server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
