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
			NeighboursResponse(w, params)
			//findCountryNeighbours(w, "PLACEHOLDER")
		} else {
			renewableUtils.SpecifiedCountryResponse(w, params)
		}
	} else { // If no country is specified, find all countries
		renewableUtils.AllCountriesResponse(w, params)
	}

}

// NeighboursResponse gives a response with the data for the specified country and its neighbours.
func NeighboursResponse(w http.ResponseWriter, params structs.URLParams) {
	// Get all countries from csv file
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the data for each neighbouring country
	neighbours, err := getNeighbours(countries, params)
	if err != nil {
		return
	}

	// Get current year data for each neighbour
	currentNeighbours := getCurrentYearData(neighbours)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the current year data for the specified country and prepend it to the neighbours
	params.BeginYear, params.EndYear = constants.CurrentYear, constants.CurrentYear
	currentCountry, err := renewableUtils.GetSpecifiedCountry(countries, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	currentNeighbours = append(currentCountry, currentNeighbours...)

	// Write the response
	json_coder.PrettyPrint(w, currentNeighbours)
}

// getNeighbours returns the data for the neighbouring countries of the specified country.
func getNeighbours(countries []structs.CountryInfo, params structs.URLParams) ([]structs.CountryInfo, error) {
	// Get the country code from the params
	countryCode := params.Country

	// Get the border data from the JSON file
	borders, err := getBorderDataFromApi(countryCode)
	if err != nil {
		return nil, err
	}

	// Get the data for the neighbouring countries
	var neighbourData []structs.CountryInfo
	for _, c := range countries {
		for _, b := range borders {
			if c.IsoCode == b {
				neighbourData = append(neighbourData, c)
			}
		}
	}

	// return the neighbour data as a JSON response.
	//E.g. {"country": "Canada", "iso_code": "CAN", "year": 2018, "value": 0.0}
	return neighbourData, err
}

// getBorderDataFromApi gets the border data for the given country code from the JSON file.
// TODO: Improve error handling
func getBorderDataFromApi(countryCode string) ([]string, error) {
	// Create the array to store the border data
	var neighbourArray []string

	// Create the api link by adding the country code to the api link
	apiLink := constants.CountryApi + constants.CountryAlpha + countryCode

	// Get the data from the api
	neighbours, err := http.Get(apiLink)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Decode the data from the api to the struct
	var countryApi = json_coder.DecodeCountryNeighbour(neighbours)
	// Get the border data from the struct
	for _, neighbour := range countryApi {
		neighbourArray = neighbour.Borders
	}

	// Return the border data as a slice of strings e.g, ["USA", "CAN"]
	return neighbourArray, nil

}

// getCurrentYearData returns the data for the current year.
func getCurrentYearData(countries []structs.CountryInfo) []structs.CountryInfo {
	// Get max year from the data
	_, maxYear := renewableUtils.GetMinMaxYear(countries)

	var lastYearData []structs.CountryInfo
	for _, c := range countries {
		if c.Year == maxYear {
			lastYearData = append(lastYearData, c)
		}
	}

	return lastYearData
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

// findSingleCountryInformation find the renewable information on the single country that have been specified
// find specified country -> renewables methods
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
// resten -> get all countrie -> Renewables methods
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
