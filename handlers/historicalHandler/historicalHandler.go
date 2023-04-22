package historicalHandler

import (
	"assignment-2/constants"
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
)

// HistoricalHandler is the handler to get historical information about the countries renewable energy production.
// For now, it uses if statements to determine what data to get. Can be changed to a switch if needed for expansion.
func HistoricalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHistoricalData(w, r)
	} else {
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+"!", http.StatusMethodNotAllowed)
		return
	}
}

// getHistoricalData is the function that handles the get request for the historical data.
func getHistoricalData(w http.ResponseWriter, r *http.Request) {
	// Gets the historical data parameters from either the URL path or the query parameters as a struct.
	params, err := utils.GetHistoricalDataParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Switch for what data to get.
	switch {
	case params.Country == "": // No country specified
		switch {
		case params.BeginYear == "" && params.EndYear == "": // No time period specified
			allCountries(w, params) // Get all countries, only shows mean percentage
		case params.BeginYear != "" && params.EndYear != "": // {?begin=year&end=year?} only mean percentage value per country
			log.Println("The mean energy percentage for each country in the period between beginYear and endYear")
		case params.BeginYear != "" && params.EndYear == "": // {?begin=year}
			log.Println("Every year from beginYear to the most recent year")
		case params.BeginYear == "" && params.EndYear != "": // {?end=year}
			log.Println("Every year from the earliest year to endYear")
		}

	case params.Country != "": // Specified country
		switch {
		case params.BeginYear == "" && params.EndYear == "": // All years for specified country
			getSpecifiedCountry(w, params)
		case params.BeginYear != "" && params.EndYear != "": // {country?begin=year&end=year?}
			log.Println("Same as getSpecifiedCountry but only entries between beginYear and endYear")
		case params.BeginYear != "" && params.EndYear == "": // {country?begin=year}
			log.Println("Same as getSpecifiedCountry but only entries from beginYear to the most recent year")
		case params.BeginYear == "" && params.EndYear != "": // {country?end=year}
			log.Println("Same as getSpecifiedCountry but only entries from the earliest year to endYear")
		}
	}
}

func getSpecifiedCountry(w http.ResponseWriter, params structs.URLParams) {

	// Get the country data from the csv file
	csv, err := utils.ReadCsv(constants.HistoricalCsv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the csv file
	countries, err := parseCountriesCsv(csv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Find all the data for the specified country from the countries slice of countryInfo structs
	var countryData []structs.CountryInfo
	for _, c := range countries.Countries {
		if c.IsoCode == params.Country {
			countryData = append(countryData, structs.CountryInfo{
				Country:    c.Country,
				IsoCode:    c.IsoCode,
				Year:       c.Year,
				Percentage: c.Percentage,
			})
		}
	}

	// If no data was found for the specified country, return an error
	if len(countryData) == 0 {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	// Sort the data by percentage value if the sort parameter is set to true
	if params.SortByValue {
		sort.Slice(countryData, func(i, j int) bool {
			return countryData[i].Percentage < countryData[j].Percentage
		})
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json_coder.PrettyPrint(w, countryData)
}

// allCountries returns all the countries in the csv file. It's used when no country is specified in the URL.
func allCountries(w http.ResponseWriter, params structs.URLParams) {
	// Get the data from the csv file
	csv, err := utils.ReadCsv(constants.HistoricalCsv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the csv file
	countries, err := parseCountriesCsv(csv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Compute the mean of the percentages for each country
	countries = computeMean(countries)

	// Sort the data by percentage value if the sort parameter is set to true
	if params.SortByValue {
		sort.Slice(countries.Countries, func(i, j int) bool {
			return countries.Countries[i].Percentage < countries.Countries[j].Percentage
		})
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json_coder.PrettyPrint(w, countries.Countries)
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

		// TODO: find a better way to omit the year field in the json response.
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

// parseCountriesCsv parses the csv file and returns a slice of countryInfo structs.
func parseCountriesCsv(records [][]string) (structs.Countries, error) {
	var countries structs.Countries

	// Iterate through the records and populate the Countries struct
	for _, record := range records {
		// TODO: Are there better ways to handle the header row?
		// Skip the header row
		if record[0] == "Entity" {
			continue
		}

		y, err := strconv.Atoi(record[2])
		if err != nil {
			return countries, fmt.Errorf("error parsing year: %s", err)
		}

		p, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			return countries, fmt.Errorf("error parsing percentage: %s", err)
		}

		countryInfo := structs.CountryInfo{
			Country:    record[0],
			IsoCode:    record[1],
			Year:       y,
			Percentage: float32(p),
		}
		countries.Countries = append(countries.Countries, countryInfo)
	}

	return countries, nil
}
