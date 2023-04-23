package utils

import (
	"assignment-2/constants"
	"assignment-2/structs"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// pathToQueryHistParams converts the URL path into query parameters. Now a user can either enter the parameters in
// the URL path or in the query parameters. For example:
// http://localhost:8080/history/NOR/1965/1999/true or
// http://localhost:8080/history?country=NOR&begin=1965&end=1999&sortByValue=true
// If using the path format, the parameters must be in the following order: country, begin, end, sortByValue,
// and if only either begin or end is specified, the other must be specified as "null".
// like so: http://localhost:8080/history/NOR/1970/null/true
func pathToQueryHistParams(r *http.Request) (url.Values, error) {
	// Parse the request's URL into a *url.URL struct
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, err
	}
	// Get the URL path and trim the prefix
	path := strings.TrimPrefix(u.Path, constants.HistoryEP)
	// Get the URL query params
	queryParams := u.Query()

	// Convert the ULR path into query params
	pathParts := strings.Split(path, "/")

	// Regex to check if the country code is 3 letters
	re := regexp.MustCompile(`^$|^[a-zA-Z]{3}$|^null$`)

	if !queryParams.Has("country") { // no country in query params
		if len(pathParts) > 0 && re.MatchString(pathParts[0]) { // country in path
			queryParams.Set("country", pathParts[0])
		} else if (len(pathParts) > 1 && pathParts[1] != "" && pathParts[1] != "null") ||
			(len(pathParts) > 2 && pathParts[2] != "" && pathParts[2] != "null") { // begin or end in path
			queryParams.Set("country", "null") // set country to null
		} else {
			return nil, errors.New("only 3 letter country codes are allowed")
		}
	}
	if len(pathParts) > 1 && pathParts[1] != "" && pathParts[1] != "null" {
		queryParams.Set("begin", pathParts[1])
	}
	if len(pathParts) > 2 && pathParts[2] != "" && pathParts[2] != "null" {
		queryParams.Set("end", pathParts[2])
	}
	if len(pathParts) > 3 && pathParts[3] != "" {
		queryParams.Set("sortByValue", pathParts[3])
	}

	// Update the request's URL with the new query parameters
	u.RawQuery = queryParams.Encode()
	r.URL = u

	// Returns the query parameters in the URL
	return queryParams, nil
}

// GetHistoricalDataParams extracts the historical data parameters from the URL query parameters and populates a
// structs.URLParams struct with the values. If the URL path is used instead of the query parameters, the URL path
// is converted into query parameters and then the query parameters are extracted.
func GetHistoricalDataParams(r *http.Request) (structs.URLParams, error) {
	params := structs.URLParams{}

	// Convert the URL path into query parameters
	queryParams, err := pathToQueryHistParams(r)
	if err != nil {
		return structs.URLParams{}, err
	}

	// Set the params struct fields
	country := queryParams.Get("country")
	if country != "" {
		if strings.ToLower(country) == "null" {
			params.Country = "null"
		} else {
			params.Country = strings.ToUpper(country)
		}
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

	// Correct the order of the years if beginYear > endYear
	params = correctYearOrder(params)

	return params, nil
}

// correctYearOrder corrects the order of the years if beginYear > endYear
func correctYearOrder(params structs.URLParams) structs.URLParams {
	if params.BeginYear != "" && params.EndYear != "" {
		begin, err1 := strconv.Atoi(params.BeginYear)
		end, err2 := strconv.Atoi(params.EndYear)
		if err1 == nil && err2 == nil && begin > end || end < begin {
			temp := params.BeginYear
			params.BeginYear = params.EndYear
			params.EndYear = temp
		}
	}
	return params
}
