package utils

import (
	"assignment-2/constants"
	"assignment-2/structs"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// convertPathToQueryParams converts the URL path into query parameters. Now a user can either enter the parameters in
// the URL path or in the query parameters. For example:
// http://localhost:8080/history/NOR/1965/1999/true or
// http://localhost:8080/history?country=NOR&begin=1965&end=1999&sortByValue=true
func convertPathToQueryParams(r *http.Request) {
	// Parse the request's URL into a *url.URL struct
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return
	}
	// Get the URL path and trim the prefix
	path := strings.TrimPrefix(u.Path, constants.HistoryEP)
	// Get the URL query params
	queryParams := u.Query()

	// Convert the ULR path into query params
	pathParts := strings.Split(path, "/")
	if len(pathParts) > 0 && pathParts[0] != "" {
		queryParams.Set("country", pathParts[0])
	}
	if len(pathParts) > 1 && pathParts[1] != "" {
		queryParams.Set("begin", pathParts[1])
	}
	if len(pathParts) > 2 && pathParts[2] != "" {
		queryParams.Set("end", pathParts[2])
	}
	if len(pathParts) > 3 && pathParts[3] != "" {
		queryParams.Set("sortByValue", pathParts[3])
	}

	// Update the request's URL with the new query parameters
	u.RawQuery = queryParams.Encode()
	r.URL = u
}

// GetHistoricalDataParams extracts the historical data parameters from the URL query parameters and populates a
// structs.URLParams struct with the values. If the URL path is used instead of the query parameters, the URL path
// is converted into query parameters and then the query parameters are extracted.
func GetHistoricalDataParams(r *http.Request) (structs.URLParams, error) {
	params := structs.URLParams{}

	// TODO: This is for debugging, remove it later
	fmt.Println("URL: ", r.URL.String())

	// Convert the URL path into query parameters
	convertPathToQueryParams(r)

	// Extract the query parameters from the URL and set the struct fields
	queryParams := r.URL.Query()

	country := queryParams.Get("country")
	if country != "" {
		params.Country = strings.ToUpper(country)
	}

	beginYear := queryParams.Get("begin")
	if beginYear != "" {
		params.BeginYear = beginYear
	}

	endYear := queryParams.Get("end")
	if endYear != "" {
		params.EndYear = endYear
	}

	sortByValue := queryParams.Get("sortByValue") == "true"
	params.SortByValue = sortByValue

	return params, nil
}
