package historicalHandler

import (
	"assignment-2/constants"
	"assignment-2/handlers/renewableHandlers/renewableUtils"
	"assignment-2/utils"
	"net/http"
)

// HistoricalHandler is the handler to get historical information about the countries renewable energy production.
func HistoricalHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getHistoricalData(w, r)
	default:
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+"!", http.StatusMethodNotAllowed)
	}
}

// getHistoricalData gets the historical data from the csv file and returns it to the user.
func getHistoricalData(w http.ResponseWriter, r *http.Request) {
	params, err := utils.GetHistoricalDataParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If no country is specified, return data for all countries, else return data for specified country
	params.EndPoint = constants.History // Set the endpoint to historical for selecting the correct data.
	if params.Country == "" || params.Country == constants.NullString {
		renewableUtils.AllCountriesResponse(w, params)
	} else {
		renewableUtils.SpecifiedCountryResponse(w, params)
	}
}
