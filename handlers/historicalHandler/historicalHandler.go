package historicalHandler

import (
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"fmt"
	"net/http"
	"sort"
	"strconv"
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
	if params.Country == "" || params.Country == "null" {
		getAllCountries(w, params)
	} else {
		getSpecifiedCountry(w, params)
	}

}

// getSpecifiedCountry returns the historical data for the specified country.
func getSpecifiedCountry(w http.ResponseWriter, params structs.URLParams) {
	// Get all countries from csv file
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter the countries by the parameters specified in the URL
	countryData, err := filterCountriesByParams(countries.Countries, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// No countries with the specified parameters were found
	if len(countryData) == 0 {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	// Sort the data by percentage of renewable energy production
	countriesToSort := structs.Countries{Countries: countryData}
	sortByValue(params.SortByValue, countriesToSort)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json_coder.PrettyPrint(w, countryData)
}

// getAllCountries returns all the countries in the csv file. It's used when no country is specified in the URL.
func getAllCountries(w http.ResponseWriter, params structs.URLParams) {

	// TODO: Evaluate the http status codes in this function.

	// Get the countries from the csv file
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter the data by the specified year range
	filteredCountriesSlice, err := filterCountriesByParams(countries.Countries, params)
	if err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Convert from CountryInfo slice to Countries struct
	filteredCountries := structs.Countries{Countries: filteredCountriesSlice}

	// Compute the mean for each country
	countries = computeMean(filteredCountries)

	// Sort in alphabetical order of country name
	sort.Slice(countries.Countries, func(i, j int) bool {
		return countries.Countries[i].Country < countries.Countries[j].Country
	})

	// Sort the data by percentage if sortByValue is true
	sortByValue(params.SortByValue, countries)

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json_coder.PrettyPrint(w, countries.Countries)
}

// filterCountriesByParams only returns the countries that match the parameters specified in the URL.
func filterCountriesByParams(countries []structs.CountryInfo, params structs.URLParams) ([]structs.CountryInfo, error) {

	beginYear, endYear, err := convertYearToInt(params)
	if err != nil {
		return nil, err
	}

	var filteredCountries []structs.CountryInfo
	for _, c := range countries {
		if params.Country != "" && params.Country != "null" && c.IsoCode != params.Country {
			continue
		}
		if beginYear == -1 && endYear == -1 { // If both beginYear and endYear are -1, return only the latest year
			if len(filteredCountries) > 0 && filteredCountries[0].Year >= c.Year {
				continue
			}
			filteredCountries = filteredCountries[:0] // Clear the slice
		} else if (beginYear != 0 && c.Year < beginYear) || (endYear != 0 && c.Year > endYear) {
			continue
		}
		filteredCountries = append(filteredCountries, c)
	}
	return filteredCountries, nil
}

// convertYearToInt converts the beginYear and endYear parameters from the URL to int.
func convertYearToInt(params structs.URLParams) (int, int, error) {
	// Convert beginYear and endYear to int
	var beginYear int
	if params.BeginYear != "" {
		var err error
		beginYear, err = strconv.Atoi(params.BeginYear)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid begin year")
		}
	}
	var endYear int
	if params.EndYear != "" {
		var err error
		endYear, err = strconv.Atoi(params.EndYear)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid end year")
		}
	}
	return beginYear, endYear, nil
}

// sortByValue sorts the countries slice by percentage of renewable energy if the sort parameter from the URL is set to
// true. The countries will be sorted in ascending order.
func sortByValue(sbv bool, countries structs.Countries) {
	if sbv {
		sort.Slice(countries.Countries, func(i, j int) bool {
			return countries.Countries[i].Percentage < countries.Countries[j].Percentage
		})
	}
}

// computeMean computes the mean of the percentages for each country and returns a new slice of countries with the mean.
// Will map the iso codes to a slice of percentages. The order of the countries will not be guaranteed.
func computeMean(countries structs.Countries) structs.Countries {

	// Map the iso codes
	isoCodeMap := make(map[string][]float32)
	for _, country := range countries.Countries {
		isoCodeMap[country.IsoCode] = append(isoCodeMap[country.IsoCode], country.Percentage)
	}

	// Create a new slice of countries with the mean
	var newCountries structs.Countries
	for isoCode, percentages := range isoCodeMap {
		mean := 0.0
		for _, percentage := range percentages {
			mean += float64(percentage)
		}
		mean /= float64(len(percentages))

		// TODO: better way to omit the year field in the json response?
		// Find the country with the iso code and set the percentage to the mean
		for _, country := range countries.Countries {
			if country.IsoCode == isoCode {
				country.Percentage = float32(mean)
				country.Year = 0 // Set Year field to 0, so it's omitted in the json response.
				newCountries.Countries = append(newCountries.Countries, country)
				break
			}
		}
	}

	return newCountries
}
