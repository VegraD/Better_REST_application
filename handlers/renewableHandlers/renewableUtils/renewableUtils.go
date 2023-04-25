package renewableUtils

import (
	"assignment-2/constants"
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)

// GetSpecifiedCountry returns the historical data for the specified country.
func GetSpecifiedCountry(w http.ResponseWriter, params structs.URLParams) {
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
	} else if len(countryData) == 0 {
		http.Error(w, "No countries with the specified parameters were found", http.StatusNotFound)
		return
	}

	// Sort the data by percentage of renewable energy production
	countriesToSort := structs.Countries{Countries: countryData}
	sortByValue(params.SortByValue, countriesToSort)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json_coder.PrettyPrint(w, countryData)
}

// GetAllCountries returns all the countries in the csv file. It's used when no country is specified in the URL.
func GetAllCountries(w http.ResponseWriter, params structs.URLParams) {

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
		// TODO: error handling.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Convert from CountryInfo slice to Countries struct
	filteredCountries := structs.Countries{Countries: filteredCountriesSlice}

	// Compute the mean for each country
	countries = computeMean(filteredCountries, params)

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

func FindCountryNeighbours(w http.ResponseWriter, params structs.URLParams) {
	// Get the country code from the params
	countryCode := params.Country

	// Get the border data from the JSON file
	borders, err := getBorderDataFromFile(countryCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get all countries from csv file
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the data for each neighbouring country
	var neighbourData []structs.CountryInfo
	var self structs.CountryInfo
	for _, country := range countries.Countries {
		for _, border := range borders {
			if country.IsoCode == border {
				neighbourData = append(neighbourData, country)
			} else if country.IsoCode == countryCode {
				// save the country itself in a variable
				self = country
			}
		}
	}

	// Filter the data by the specified year range
	filteredNeighbourData, err := filterCountriesByParams(neighbourData, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if len(filteredNeighbourData) == 0 {
		http.Error(w, "No neighbours with the specified parameters were found", http.StatusNotFound)
		return
	}

	// Prepend the country itself to the neighbour data
	filteredNeighbourData = append([]structs.CountryInfo{self}, filteredNeighbourData...)

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json_coder.PrettyPrint(w, filteredNeighbourData)
}

func getBorderDataFromFile(countryCode string) ([]string, error) {
	// Open and read the JSON file
	file, err := utils.OpenFile(constants.CountriesJSON)
	if err != nil {
		return nil, err
	}
	defer utils.CloseFile(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data
	var countries []struct {
		Cca3    string   `json:"cca3"`
		Borders []string `json:"borders"`
	}
	err = json.Unmarshal(data, &countries)
	if err != nil {
		return nil, err
	}

	// Get the border data for the given country code
	var borderData []string
	for _, country := range countries {
		if country.Cca3 == countryCode {
			borderData = country.Borders
			break
		}
	}
	if len(borderData) == 0 {
		return nil, errors.New("no border data found for the given country code")
	}

	return borderData, nil
}

func filterCountriesByParams(countries []structs.CountryInfo, params structs.URLParams) ([]structs.CountryInfo, error) {
	beginYear, endYear, err := convertYearToInt(params)
	if err != nil {
		return nil, err
	}

	// Find the min and max year in the dataset.
	minYear, maxYear := findMinAndMaxYear(countries)

	// Filter the countries by the parameters specified in the URL
	var filteredCountries []structs.CountryInfo
	for _, c := range countries {
		if c.IsoCode == "" { // Skip regions with empty IsoCode
			continue
		}
		if !params.Neighbours && params.Country != "" && params.Country != constants.NullString &&
			c.IsoCode != params.Country {
			continue // If the country is not the one specified in the URL, skip it
		}
		if beginYear == -1 && endYear == -1 { // If both beginYear and endYear are -1, return only the latest year
			if c.Year != maxYear {
				continue // If the year is not the latest, skip it
			}
		} else if (beginYear != 0 && c.Year < beginYear) || (endYear != 0 && c.Year > endYear) {
			continue // If the data for the country's year is outside the range, skip it
		}
		filteredCountries = append(filteredCountries, c)
	}

	// If no data was found for the given year or year range, return an error
	if len(filteredCountries) == 0 {
		if (params.BeginYear != "" && params.BeginYear != "-1" && beginYear < minYear) ||
			(params.EndYear != "" && params.EndYear != "-1" && endYear > maxYear) ||
			(params.BeginYear != "" && params.BeginYear != "-1" && beginYear > maxYear) ||
			(params.EndYear != "" && params.EndYear != "-1" && endYear < minYear) {
			return nil, errors.New("no data found for the given year or range of years")
		}
	}

	return filteredCountries, nil
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

// findMinAndMaxYear finds the minimum and maximum year in the dataset.
func findMinAndMaxYear(countries []structs.CountryInfo) (int, int) {
	// sort the countries by year, then return the first and last element
	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Year < countries[j].Year
	})
	return countries[0].Year, countries[len(countries)-1].Year
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
func computeMean(countries structs.Countries, params structs.URLParams) structs.Countries {

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
				if params.EndPoint == constants.History {
					country.Year = 0 // Set Year field to 0, so it's omitted in the json response.
				}
				newCountries.Countries = append(newCountries.Countries, country)
				break
			}
		}
	}

	return newCountries
}