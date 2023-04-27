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

// pathToQueryHistParams converts the URL path into query parameters.
// Now a user can either enter the parameters in the URL path or in the query parameters. For example:
// http://localhost:8080/history/NOR/1965/1999/true or
// http://localhost:8080/history?country=NOR&begin=1965&end=1999&sortByValue=true
// If using the path format, the parameters must be in the following order: country, begin, end, sortByValue,
// and if only either begin or end is specified, the other must be specified as "null".
// like so: http://localhost:8080/history/NOR/1970/null/true
func pathToQueryHistParams(r *http.Request) (url.Values, error) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, err
	}

	path := strings.TrimPrefix(u.Path, constants.HistoryEP)
	queryParams := u.Query()
	pathParts := strings.Split(path, "/")

	err = histCountryRegex(queryParams, pathParts)
	if err != nil {
		return nil, err
	}

	validateBeginEndYear(queryParams, pathParts)
	setSortByYear(queryParams, pathParts)

	u.RawQuery = queryParams.Encode()
	r.URL = u

	// queryParams is on the form: map[country:[nor] begin:[2011] end:[2013] sortByValue:[true]]
	return queryParams, nil
}

// pathToQueryCurrParams converts the URL path into query parameters.
// {country?} refers to an optional country 3-letter code or the full name of a country.
// {?neighbours=bool?} refers to an optional parameter indicating whether neighbouring countries' values should be shown.
// That means a valid URL can be either: http://localhost:8080/current/{country?}?neighbours={bool?} or
// http://localhost:8080/current/country/bool
func pathToQueryCurrParams(r *http.Request) (url.Values, error) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, err
	}

	path := strings.TrimPrefix(u.Path, constants.CurrentEP)
	queryParams := u.Query()
	pathParts := strings.Split(path, "/")

	err = currCountryRegex(queryParams, pathParts)
	if err != nil {
		return nil, err
	}

	setNeighbours(queryParams, pathParts)

	u.RawQuery = queryParams.Encode()
	r.URL = u

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

	re := regexp.MustCompile(`^$|^[a-zA-Z\s]{3,50}$`)

	// Set the params struct fields
	country := queryParams.Get(constants.CountryString)
	if country != "" && country != constants.NullString { // Country is defined
		if !re.MatchString(country) { // Country name is invalid
			return structs.URLParams{}, errors.New("country name can either be empty or contain 3-50 letters")
		} else if len(country) > 3 { // Country name is valid and longer than 3 letters
			c, err := convertNameToCode(country)
			if err != nil {
				return structs.URLParams{}, err
			}
			params.Country = c
		} else { // Country name is valid and 3 letters --> Country name == country code
			params.Country = strings.ToUpper(country)
		}
	}

	beginYear := queryParams.Get(constants.BeginString)
	if beginYear != "" {
		params.BeginYear = beginYear
	}

	endYear := queryParams.Get(constants.EndString)
	if endYear != "" {
		params.EndYear = endYear
	}

	sortByValue := queryParams.Get(constants.SortByValueString) == constants.TrueString
	params.SortByValue = sortByValue

	// Correct the order of the years if beginYear > endYear
	params = correctYearOrder(params)

	return params, nil
}

// GetCurrentDataParams extracts the current data parameters from the URL query parameters and populates a
// structs.URLParams struct with the values. If the URL path is used instead of the query parameters, the URL path
// is converted into query parameters and then the query parameters are extracted.
func GetCurrentDataParams(r *http.Request) (structs.URLParams, error) {
	params := structs.URLParams{}

	// Convert the URL path into query parameters
	queryParams, err := pathToQueryCurrParams(r)
	if err != nil {
		return structs.URLParams{}, err
	}

	re := regexp.MustCompile(`^$|^[a-zA-Z\s]{3,50}$`)

	// Set the params struct fields
	country := queryParams.Get(constants.CountryString)
	if country != "" {
		if !re.MatchString(country) {
			return structs.URLParams{}, errors.New("country name can either be empty or contain 3-50 letters")
		} else if len(country) > 3 {
			c, err := convertNameToCode(country)
			if err != nil {
				return structs.URLParams{}, err
			}
			country = c
		}
	}
	params.Country = strings.ToUpper(country)

	neighbours := queryParams.Get(constants.NeighboursString) == constants.TrueString
	params.Neighbours = neighbours

	return params, nil
}

// convertNameToCode converts the country name to the country code.
func convertNameToCode(countryName string) (string, error) {
	// Get the countries from the csv file
	countries, err := GetCountriesFromCsv()
	if err != nil {
		return "", err
	}

	// Get the country code
	for _, c := range countries {
		if strings.EqualFold(c.Country, countryName) {
			return c.IsoCode, nil
		}
	}

	return "", errors.New("no country with the specified country code was found")
}

// validateCountry validates the country parameter in the URL path.
func validateCountry(queryParams url.Values, pathParts []string, re *regexp.Regexp) error {
	if !queryParams.Has(constants.CountryString) {
		if len(pathParts) > 0 && re.MatchString(pathParts[0]) {
			queryParams.Set(constants.CountryString, pathParts[0])
		} else {
			return errors.New("invalid country parameter")
		}
	}
	return nil
}

// histCountryRegex validates the country parameter in the URL path for the historical endpoint.
// TODO: Combine the regex functions into one function.
func histCountryRegex(queryParams url.Values, pathParts []string) error {
	re := regexp.MustCompile(`^$|^[a-zA-Z\s]{3,50}$`)
	return validateCountry(queryParams, pathParts, re)
}

// currCountryRegex validates the country parameter in the URL path for the current endpoint.
func currCountryRegex(queryParams url.Values, pathParts []string) error {
	re := regexp.MustCompile(`^$|^[a-zA-Z\s]{3,50}$`)
	return validateCountry(queryParams, pathParts, re)
}

// validateBeginEndYear validates the begin and end year parameters in the URL path.
func validateBeginEndYear(queryParams url.Values, pathParts []string) {
	// TODO: Include other parameters in the validation other than just checking if they are specified.
	if len(pathParts) > 1 && pathParts[1] != "" && pathParts[1] != constants.NullString {
		queryParams.Set(constants.BeginString, pathParts[1])
	}
	if len(pathParts) > 2 && pathParts[2] != "" && pathParts[2] != constants.NullString {
		queryParams.Set(constants.EndString, pathParts[2])
	}
}

// setSortByYear sets the sortByValue query parameter to true if the sortByValue path parameter is "true".
func setSortByYear(queryParams url.Values, pathParts []string) {
	for _, p := range pathParts {
		if strings.ToLower(p) == constants.TrueString {
			queryParams.Set(constants.SortByValueString, constants.TrueString)
			break
		}
	}
}

// setNeighbours sets the neighbours query parameter to true if the neighbours path parameter is "true".
func setNeighbours(queryParams url.Values, pathParts []string) {
	for _, p := range pathParts {
		if strings.ToLower(p) == constants.TrueString {
			queryParams.Set(constants.NeighboursString, constants.TrueString)
			break
		}
	}
}

// correctYearOrder corrects the order of the years if beginYear > endYear
func correctYearOrder(params structs.URLParams) structs.URLParams {
	if params.BeginYear != "" && params.BeginYear != constants.NullString &&
		params.EndYear != "" && params.EndYear != constants.NullString {
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
