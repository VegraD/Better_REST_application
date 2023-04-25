package currentHandler

import (
	"assignment-2/constants"
	"assignment-2/handlers/renewableHandlers/renewableUtils"
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"fmt"
	"net/http"
	"strings"
)

// CurrentHandler is the handler to get current information about countries renewable energy
func CurrentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		handleRenewablesCurrentGetRequest(w, r)
	} else {
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

// handleRenewablesCurrentGetRequest handles the get request for the current renewable energy information
func handleRenewablesCurrentGetRequest(w http.ResponseWriter, r *http.Request) {
	//url := url.parse(r.URL.String())
	//pathBase := path.Base(r.URL.Path)
	//if pathBase == "current" {
	//	//Find information for all countries
	//	findAllCountriesInformation(w, r)
	//} else if strings.EqualFold(r.URL.Query().Get("neighbours"), "true") {
	//	//Find information for neighbours
	//	findCountryNeighbours(w, pathBase)
	//
	//} else {
	//	//Find information for country
	//	singleCountry := findSingleCountryInformation(w, pathBase)
	//	json_coder.PrettyPrint(w, singleCountry)
	//}

	// Get the parameters from the URL
	params, err := utils.GetCurrentDataParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set flags for getting the correct data
	params.EndPoint = constants.Current
	params.BeginYear, params.EndYear = constants.CurrentYear, constants.CurrentYear

	if params.Country != "" && params.Country != constants.NullString { // If a country is specified
		if params.Neighbours {
			renewableUtils.NeighboursResponse(w, params)
			//findCountryNeighbours(w, "PLACEHOLDER")
		} else {
			renewableUtils.SpecifiedCountryResponse(w, params)
		}
	} else { // If no country is specified, find all countries
		renewableUtils.AllCountriesResponse(w, params)
	}

}

// findSingleCountryInformation find the renewable information on the single country that have been specified
func findSingleCountryInformation(w http.ResponseWriter, pathBase string) []structs.CountryInfo {
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Filter the countries by the parameters specified in the URL
	var singleCountry []structs.CountryInfo
	for _, c := range countries {
		if strings.EqualFold(c.Country, pathBase) || strings.EqualFold(c.IsoCode, pathBase) {
			if len(singleCountry) > 0 && singleCountry[0].Year >= c.Year {
				continue
			}
			singleCountry = singleCountry[:0] // Clear the slice
			singleCountry = append(singleCountry, c)
		}
	}

	// No countries with the specified parameters were found
	if len(singleCountry) == 0 {
		http.Error(w, "Country not found", http.StatusNotFound)
		return nil
	}

	return singleCountry

}

// findAllCountriesInformation finds the current information about all countries
func findAllCountriesInformation(w http.ResponseWriter, r *http.Request) {
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var currentYearCountries []structs.CountryInfo
	for _, c := range countries {
		// Skip countries that don't have an iso code as they're not countries
		if c.IsoCode == "" {
			continue
		}
		// Append c if the slice is empty or the last country in the slice is not the same as c
		if len(currentYearCountries) == 0 || currentYearCountries[len(currentYearCountries)-1].Country != c.Country {
			currentYearCountries = append(currentYearCountries, c)
			// If the last country in the slice is the same as c, replace it if c is newer
		} else if currentYearCountries[len(currentYearCountries)-1].Year < c.Year {
			currentYearCountries = currentYearCountries[:len(currentYearCountries)-1]
			currentYearCountries = append(currentYearCountries, c)
		}
	}

	json_coder.PrettyPrint(w, currentYearCountries)

}

// finne land so jeg er ute etter.
// print dette landet
// bruke landets isokode til å finne den i json
// legge naboer inn i liste
// for hvert element i listen, print det
func findCountryNeighbours(w http.ResponseWriter, pathBase string) {
	//country := findSingleCountryInformation(w, pathBase)
	var allCountries []structs.CountryInfo
	originalCountry := findSingleCountryInformation(w, pathBase)
	allCountries = append(allCountries, originalCountry...)

	// get the information about the neighbours from the country api
	apiLink := constants.CountryApi
	if len(pathBase) == 3 {
		apiLink += constants.CountryAlpha + pathBase
	} else {
		apiLink += constants.CountryFullTextName + pathBase + constants.CountryFullText
	}

	neighbours, err := http.Get(apiLink)
	if err != nil {
		fmt.Print(err.Error())
	}

	// decode the information about the countries from the country api
	var countryApi = json_coder.DecodeCountryNeighbour(neighbours)
	for _, neighbour := range countryApi {
		for _, neighbour := range neighbour.Borders {
			countryInfo := findSingleCountryInformation(w, neighbour)
			allCountries = append(allCountries, countryInfo...)

		}
	}

	json_coder.PrettyPrint(w, allCountries)

}
