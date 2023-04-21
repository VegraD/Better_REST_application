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
func DecodeCountryInfo(httpResponse *http.Response) []structs.CountryInfo {
	var country []structs.CountryInfo

	// Create a new decoder
	decoder := json.NewDecoder(httpResponse.Body)

	// Decode the json body into the struct
	if err := decoder.Decode(&country); err != nil {
		log.Print(err, http.StatusNoContent)
	}

	return country
}
