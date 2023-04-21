package historicalHandler

import (
	"assignment-2/constants"
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"fmt"
	"net/http"
	"strconv"
)

func HistoricalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHistoricalData(w, r)
	} else {
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+"!", http.StatusMethodNotAllowed)
		return
	}
}

func getHistoricalData(w http.ResponseWriter, r *http.Request) {

	// Get the path
	path := r.URL.Path

	switch path {
	case constants.HistoryEP:
		allCountries(w, r)
	case constants.HistoryEP + "/?country={country?}":
		// Send a test response, just a string
		fmt.Fprintf(w, "You requested data for a specific country")
		specifiedCountry(w, r)
	case constants.HistoryEP + "/{country?}" + "{?begin=year&end=year?}":
		// Find information on a specific country and a specific time period
	case constants.HistoryEP + "/{country?}" + "{?begin=year&end=year?}" + "{?sortByValue=bool?}":
		// Find information on a specific country and a specific time period and sort by value
	default:
		// Default needed?
	}
}

func specifiedCountry(w http.ResponseWriter, r *http.Request) {
	// example request: /energy/v1/renewables/history/nor
	// example request: /energy/v1/renewables/history/nor?country=nor
	/*
		Example response:
		[
		    {
		        "name": "Norway",
		        "isoCode": "NOR",
		        "year": "1965",
		        "percentage": 67.87996
		    },
		    {
		        "name": "Norway",
		        "isoCode": "NOR",
		        "year": "1966",
		        "percentage": 65.3991
		    },
		    ...
					    {
		        "name": "Norway",
		        "isoCode": "NOR",
		        "year": "2023",     <- Present year
		        "percentage": 65.3991
		    },
		]

	*/

	// Get the country query parameter
	country := r.URL.Query().Get("country")
	if country == "" {
		http.Error(w, "Missing country query parameter", http.StatusBadRequest)
		return
	}

	// Get the csv
	csv, err := utils.ReadCsv(constants.HistoricalCsv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the csv
	countries, err := parseCountriesCsv(csv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Find the country matching the query parameter and put all the years in a slice
	var years []structs.CountryInfo
	for _, c := range countries.Countries {
		if c.IsoCode == country {
			years = append(years, c)
		}
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json_coder.PrettyPrint(w, years)

}

func allCountries(w http.ResponseWriter, r *http.Request) {

	// If there are no country query parameters, the mean of each country is returned.
	// Example request: /energy/v1/renewables/history/ (no query parameters)
	if r.URL.Query().Get("country") == "" {
		// Get all countries
		csv, err := utils.ReadCsv(constants.HistoricalCsv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		countries, err := parseCountriesCsv(csv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Compute the mean
		countries = computeMean(countries)

		// Write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json_coder.PrettyPrint(w, countries)
	}

}

func computeMean(countries structs.Countries) structs.Countries {
	// Map the iso codes to a slice of percentages. Using map to avoid duplicates and for faster lookup, but the
	// order of the countries will be lost and likely be different for each request.
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

		// Find the country with the iso code and set the percentage to the mean
		for _, country := range countries.Countries {
			if country.IsoCode == isoCode {
				country.Percentage = float32(mean)
				country.Year = 0 // Set Year field to 0, so it's omitted in the response.
				newCountries.Countries = append(newCountries.Countries, country)
				break
			}
		}
	}

	return newCountries
}

func parseCountriesCsv(records [][]string) (structs.Countries, error) {
	var countries structs.Countries

	// Iterate through the records and populate the Countries struct
	for _, record := range records {
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
