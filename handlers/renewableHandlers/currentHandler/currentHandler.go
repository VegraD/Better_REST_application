package currentHandler

import (
	"assignment-2/constants"
	"assignment-2/handlers/renewableHandlers/renewableUtils"
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"fmt"
	"net/http"
)

// CurrentHandler is the handler to get current information about countries renewable energy
func CurrentHandler(w http.ResponseWriter, r *http.Request) {

	// Switch for the different http methods
	switch r.Method {
	case http.MethodGet:
		handleRenewablesCurrentGetRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
	}
}

// handleRenewablesCurrentGetRequest handles the get request for the current renewable energy information
func handleRenewablesCurrentGetRequest(w http.ResponseWriter, r *http.Request) {

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
		if params.Neighbours { // If neighbours is true, get the neighbours
			NeighboursResponse(w, params)
		} else { // If neighbours is false, get the specified country
			renewableUtils.SpecifiedCountryResponse(w, params)
		}
	} else { // If no country is specified, find all countries
		renewableUtils.AllCountriesResponse(w, params)
	}
}

// NeighboursResponse gives a response with the data for the specified country and its neighbours.
func NeighboursResponse(w http.ResponseWriter, params structs.URLParams) {
	// Get all countries from the csv file
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
// return the neighbour data as a JSON response.
// E.g. {"country": "Canada", "iso_code": "CAN", "year": 2018, "value": 0.0}
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

	return neighbourData, err
}

// getBorderDataFromApi gets the border data for the given country code from the JSON file.
// TODO: Improve error handling
func getBorderDataFromApi(countryCode string) ([]string, error) {
	// Create the array to store the border data
	var neighbourArray []string

	// Create the api link by adding the country code to the api link
	apiLink := constants.CountryApi + constants.CountryApiVersion + constants.CountryAlpha + countryCode
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
