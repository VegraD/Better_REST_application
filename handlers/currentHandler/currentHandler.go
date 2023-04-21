package currentHandler

import (
	"assignment-2/constants"
	"net/http"
	"path"
)

/*
RenewablesCurrentHandler is the handler to get current information about countries renewable energy
*/
func CurrentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		handleRenewablesCurrentGetRequest(w, r)
	} else {
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

/*
handleRenewablesCurrentGetRequest handles the get request for the current renewable energy information
*/
func handleRenewablesCurrentGetRequest(w http.ResponseWriter, r *http.Request) {
	//url := url.parse(r.URL.String())
	pathBase := path.Base(r.URL.Path)
	if pathBase == "current" {
		//Find information for all countries
	} else if checkForQueryParams(r) {
		//Find information for neighbours

	} else {
		//Find information for country
		buildCountryUrl(pathBase)
	}
}

/*
checkForQueryParams checks if the query param "neighbours" is true
*/
func checkForQueryParams(r *http.Request) bool {
	q := r.URL.Query()
	if q.Get("country") == "true" {
		return true
	}
	return false
}

/*
getCountryInformation gets the information about the country
*/
func buildCountryUrl(country string) string {
	// checks if the country is three letters
	if len(country) == 3 {
		return constants.CountryApi + constants.CountryAlpha + country
	} else {
		return constants.CountryApi + constants.CountryFullText + country
	}
}
