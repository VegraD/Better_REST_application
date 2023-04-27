package renewableUtils

import (
	"assignment-2/constants"
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"assignment-2/webhooks"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

// SpecifiedCountryResponse gives a response with data for the specified country.
func SpecifiedCountryResponse(w http.ResponseWriter, params structs.URLParams) {
	// Get all countries from csv file
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get data for the specified country
	countryData, err := GetSpecifiedCountry(countries, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Filter data by year range
	countryData, err = getDataInYearRange(countryData, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sort the data by percentage of renewable energy production if sortByValue is true
	sortByValue(params.SortByValue, countryData)

	// Errors are handled in the function
	webhooks.InvokeWebhook(params.Country)

	// Write response
	json_coder.PrettyPrint(w, countryData)
}

// AllCountriesResponse gives a response with data for all countries.
func AllCountriesResponse(w http.ResponseWriter, params structs.URLParams) {

	// Get the countries  from the csv file
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter data by year range
	filteredCountries, err := getDataInYearRange(countries, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Compute the mean for each country
	filteredCountriesMean := computeMean(filteredCountries, params)

	// Sort in alphabetical order of country name
	sort.Slice(filteredCountriesMean, func(i, j int) bool {
		return filteredCountriesMean[i].Country < filteredCountriesMean[j].Country
	})

	// Sort the data by percentage if sortByValue is true
	sortByValue(params.SortByValue, filteredCountriesMean)

	// Errors are handled in the function
	webhooks.InvokeWebhook(params.Country)

	// Write the response
	json_coder.PrettyPrint(w, filteredCountriesMean)
}

// getDataInYearRange returns the data for the specified year range.
func getDataInYearRange(countries []structs.CountryInfo, params structs.URLParams) ([]structs.CountryInfo, error) {
	// Get the begin and end year as integers
	beginYear, endYear, err := convertYearToInt(params)
	if err != nil {
		return nil, err
	}

	// Getting current year
	if params.BeginYear == constants.CurrentYear && params.EndYear == constants.CurrentYear {
		_, maxYear := GetMinMaxYear(countries)
		beginYear = maxYear
		endYear = maxYear
	}

	// Get the data for the specified year range
	var yearRangeData []structs.CountryInfo
	for _, c := range countries {
		if (params.BeginYear == "" || params.BeginYear == constants.NullString || c.Year >= beginYear) &&
			(params.EndYear == "" || params.EndYear == constants.NullString || c.Year <= endYear) {
			yearRangeData = append(yearRangeData, c)
		}
	}

	return yearRangeData, nil
}

// GetSpecifiedCountry returns the data for the specified country.
// TODO: Have made function public. Consider splitting this function to have lower cohesion in code
func GetSpecifiedCountry(countries []structs.CountryInfo, params structs.URLParams) ([]structs.CountryInfo, error) {
	// Get the country code from the params
	countryQuery := params.Country

	_, max := GetMinMaxYear(countries)

	// Get the data for the specified country
	var country []structs.CountryInfo
	for _, c := range countries {
		if c.IsoCode == countryQuery || c.Country == countryQuery {
			// Get the current year data if specified
			if params.BeginYear == constants.CurrentYear && params.EndYear == constants.CurrentYear {
				if c.Year == max {
					country = append(country, c)
				}
			} else {
				country = append(country, c)
			}
		}
	}

	// Check if the country was found
	if len(country) == 0 {
		return nil, errors.New("no country with the specified country code was found")
	}

	return country, nil
}

// convertYearToInt converts the beginYear and endYear parameters from the URL to int.
func convertYearToInt(params structs.URLParams) (int, int, error) {
	var beginYear, endYear int
	var err error
	if params.BeginYear != "" {
		beginYear, err = strconv.Atoi(params.BeginYear)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid begin year")
		}
	}
	if params.EndYear != "" {
		endYear, err = strconv.Atoi(params.EndYear)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid end year")
		}
	}
	return beginYear, endYear, nil
}

// GetMinMaxYear finds the minimum and maximum year in the dataset.
// TODO: Have made public. Consider splitting this function to have lower cohesion in code
func GetMinMaxYear(countries []structs.CountryInfo) (int, int) {
	// sort the countries by year, then return the first and last element
	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Year < countries[j].Year
	})
	return countries[0].Year, countries[len(countries)-1].Year
}

// sortByValue sorts the countries slice by percentage of renewable energy if the sort parameter from the URL is set to
// true. The countries will be sorted in ascending order.
func sortByValue(sbv bool, countries []structs.CountryInfo) {
	if sbv {
		sort.Slice(countries, func(i, j int) bool {
			return countries[i].Percentage < countries[j].Percentage
		})
	}
}

// computeMean computes the mean of the percentages for each country and returns a new slice of countries with the mean.
// Will map the iso codes to a slice of percentages. The order of the countries will not be guaranteed.
func computeMean(countries []structs.CountryInfo, params structs.URLParams) []structs.CountryInfo {

	// Map the iso codes
	isoCodeMap := make(map[string][]float32)
	for _, country := range countries {
		isoCodeMap[country.IsoCode] = append(isoCodeMap[country.IsoCode], country.Percentage)
	}

	// Create a new slice of countries with the mean
	var newCountries []structs.CountryInfo
	for isoCode, percentages := range isoCodeMap {
		mean := 0.0
		for _, percentage := range percentages {
			mean += float64(percentage)
		}
		mean /= float64(len(percentages))

		// Find the country with the iso code and set the percentage to the mean
		for _, country := range countries {
			if country.IsoCode == isoCode {
				country.Percentage = float32(mean)
				country.Year = 0 // Set Year field to 0, so it's omitted in the json response.
				newCountries = append(newCountries, country)
				break
			}
		}
	}

	return newCountries
}
