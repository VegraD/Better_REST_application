package historicalHandler

import (
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
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
	if params.Country == "" {
		getAllCountries(w, params)
	} else {
		getSpecifiedCountry(w, params)
	}

}

func getSpecifiedCountry(w http.ResponseWriter, params structs.URLParams) {
	// Get all countries from csv file
	countries, err := utils.GetCountriesFromCsv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert beginYear and endYear to int
	var beginYear int
	if params.BeginYear != "" {
		beginYear, err = strconv.Atoi(params.BeginYear)
		if err != nil {
			http.Error(w, "Invalid begin year", http.StatusBadRequest)
			return
		}
	}
	var endYear int
	if params.EndYear != "" {
		endYear, err = strconv.Atoi(params.EndYear)
		if err != nil {
			http.Error(w, "Invalid end year", http.StatusBadRequest)
			return
		}
	}

	// Get the data for the specified country with the specified time period
	var countryData []structs.CountryInfo
	for _, c := range countries.Countries {
		if c.IsoCode == params.Country {
			if params.BeginYear != "" && c.Year < beginYear {
				continue
			}
			if params.EndYear != "" && c.Year > endYear {
				continue
			}
			countryData = append(countryData, structs.CountryInfo{
				Country:    c.Country,
				IsoCode:    c.IsoCode,
				Year:       c.Year,
				Percentage: c.Percentage,
			})
		}
	}

	if len(countryData) == 0 {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	// Sort the data by percentage of renewable energy production
	countriesToSort := structs.Countries{Countries: countryData}
	sortByValue(params, countriesToSort)

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

	// Convert beginYear and endYear to int
	var beginYear int
	if params.BeginYear != "" {
		beginYear, err = strconv.Atoi(params.BeginYear)
		if err != nil {
			http.Error(w, "Invalid begin year", http.StatusBadRequest)
			return
		}
	}
	var endYear int
	if params.EndYear != "" {
		endYear, err = strconv.Atoi(params.EndYear)
		if err != nil {
			http.Error(w, "Invalid end year", http.StatusBadRequest)
			return
		}
	}

	// Filter the data by the specified year range
	var filteredCountries structs.Countries
	for _, c := range countries.Countries {
		if params.BeginYear != "" && c.Year < beginYear {
			continue
		}
		if params.EndYear != "" && c.Year > endYear {
			continue
		}
		filteredCountries.Countries = append(filteredCountries.Countries, c)
	}

	countries = computeMean(filteredCountries)

	// Sort the data
	sortByValue(params, countries)

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json_coder.PrettyPrint(w, countries.Countries)
}

// sortByValue sorts the countries slice by percentage of renewable energy if the sort parameter from the URL is set to
// true. The countries will be sorted in ascending order.
func sortByValue(params structs.URLParams, countries structs.Countries) {
	if params.SortByValue {
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
