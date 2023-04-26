package json_coder

import (
	"assignment-2/constants"
	"encoding/json"
	"log"
	"net/http"
)

// PrettyPrint gotten from 02-JSON-demo and modified slightly.
// Using an interface so that no extra method is needed.
func PrettyPrint(w http.ResponseWriter, in interface{}) {

	// Set content type to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	output, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		log.Println(constants.MarshallingErr + err.Error())
		http.Error(w, constants.MarshallingErr, http.StatusInternalServerError)
		return
	}

	//_, err = fmt.Fprintf(w, string(output)) // TODO: Why the Fprintf?
	_, err = w.Write(output)
	if err != nil {
		log.Println(constants.PrettyPrintErr + err.Error())
		http.Error(w, constants.PrettyPrintErr, http.StatusInternalServerError)
		return
	}

}
