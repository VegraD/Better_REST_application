package json_coder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
PrettyPrint gotten from 02-JSON-demo
Using an interface so that no extra method is needed.
*/
func PrettyPrint(w http.ResponseWriter, in interface{}) {

	// Set content type to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	output, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		log.Println("Error during pretty printing of output: " + err.Error())
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(output))
}
