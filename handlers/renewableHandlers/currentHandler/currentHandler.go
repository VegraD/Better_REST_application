package currentHandler

import (
	"net/http"
	"path"
)

// CurrentHandler is the handler to get current information about countries renewable energy
func CurrentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		handleRenewablesCurrentGetRequest(w, r)
	} else {
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

// handleRenewablesCurrentGetRequest handles the get request for the current renewable energy information
func handleRenewablesCurrentGetRequest(w http.ResponseWriter, r *http.Request) {
	//url := url.parse(r.URL.String())
	pathBase := path.Base(r.URL.Path)
	if pathBase == "current" {
		//Find information for all countries
	} else if r.URL.Query().Get("country") == "true" {
		//Find information for neighbours

	} else {
		//Find information for country
		findSingleCountryInformation(w, pathBase)
	}
}

// findSingleCountryInformation finds the information about a single country
func findSingleCountryInformation(w http.ResponseWriter, pathBase string) {
	url := buildCountryUrl(pathBase)

	country, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}

	// decode the information about the countries from the country api
	var countryApi = json_coder.DecodeCountryInfo(country)
	json_coder.PrettyPrint(w, countryApi)

	csvResp := FetchCsvData()
	csvData, err := DecodeCsvData(csvResp)
	if err != nil {
		fmt.Print(err.Error())
	}
	GetCountryRenewables(csvData, pathBase)

}

// buildCountryUrl builds the url for the country api
func buildCountryUrl(country string) string {
	// checks if the country is three letters
	if len(country) == 3 {
		return constants.CountryApi + constants.CountryAlpha + country
	} else {
		return constants.CountryApi + constants.CountryFullTextName + country + constants.CountryFullText
	}
}

// FetchCsvData fetches the csv data from the url
func FetchCsvData() *http.Response {
	url := constants.RenewablesApi
	// get the information from the apis
	csvDataResp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}
	//defer csvData.Body.Close()

	return csvDataResp
}

// DecodeCsvData decodes the csv data
func DecodeCsvData(csvDataResp *http.Response) ([]structs.Renewables, error) {
	// decode the csv data

	reader := csv.NewReader(csvDataResp.Body)
	reader.TrimLeadingSpace = true
	csvData, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var csvDataDecoded = csv_coder.DecodeRenewables(csvData)
	return csvDataDecoded, nil
}

// GetCountryRenewables gets the latest renewable energy information for the specified country
func GetCountryRenewables(csvData []structs.Renewables, country string) (int, float64) {
	var latestYear int
	var renewablePercantage float64

	for _, csvData := range csvData {
		if csvData.Entity == country && csvData.Year > latestYear {
			latestYear = csvData.Year
			renewablePercantage = csvData.Renewables
		}
	}

	return latestYear, renewablePercantage
}
