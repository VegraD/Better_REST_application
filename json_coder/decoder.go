package json_coder

import (
	"assignment-2/structs"
	"encoding/json"
	"log"
	"net/http"
)

/*
DecodeCountryInfo is a method that takes a http request and decodes the json body
*/
func DecodeCountryNeighbour(httpResponse *http.Response) []structs.Border {
	var neighbours []structs.Border

	// Create a new decoder
	decoder := json.NewDecoder(httpResponse.Body)

	// Decode the json body into the struct
	if err := decoder.Decode(&neighbours); err != nil {
		log.Print(err, http.StatusNoContent)
	}

	return neighbours
}
