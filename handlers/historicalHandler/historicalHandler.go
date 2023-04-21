package historicalHandler

import (
	"assignment-2/constants"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func HistoricalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHistoricalData(r)
	} else {
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+"!", http.StatusMethodNotAllowed)
		return
	}
}

func getHistoricalData(r *http.Request) {
	switch url.QueryEscape(r.URL.Path) {
	case constants.HistoryEP:
		// Find information for all countries
	case constants.HistoryEP + "/{country?}":
		// Find information on a specific country
	case constants.HistoryEP + "/{country?}" + "{?begin=year&end=year?}":
		// Find information on a specific country and a specific time period
	case constants.HistoryEP + "/{country?}" + "{?begin=year&end=year?}" + "{?sortByValue=bool?}":
		// Find information on a specific country and a specific time period and sort by value
	default:
		// Default needed?
	}
}

// Validates the country query parameter.
func validateCountryCode(country string) bool {
	// TODO: Check if country is a valid country code and not just 3 random letters.
	if len(country) != 3 {
		return false
	}
	return true
}

// Validates begin and end query parameters. Year must be 4 digits and begin must be before end, and neither can not be
// in the future.
func validateBeginAndEnd(begin string, end string) bool {
	currentYear := time.Now().Year()

	beginQueryAsInt, err := strconv.Atoi(begin)
	if err != nil {
		return false
	}

	endQueryAsInt, err := strconv.Atoi(end)
	if err != nil {
		return false
	}

	if len(begin) < 4 || len(end) < 4 || endQueryAsInt > currentYear || beginQueryAsInt > currentYear {
		return false
	}

	if beginQueryAsInt > endQueryAsInt {
		temp := beginQueryAsInt
		beginQueryAsInt = endQueryAsInt
		endQueryAsInt = temp
	}

	return true
}
