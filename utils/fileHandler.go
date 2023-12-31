package utils

import (
	"assignment-2/constants"
	"assignment-2/structs"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// OpenFile is a wrapper function for opening a file. Will return an error if the file does not exist or if it could
// not be opened.
func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist"+"%s", err)
		} else {
			return nil, fmt.Errorf("failed to open file"+"%s", err)
		}
	}
	return file, nil
}

// CloseFile is a wrapper function for closing a file. Could be used with defer for cleaner code.
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Printf("failed to close file: "+"%s", err)
	}
}

// readCsv reads a CSV file and returns the rows as a 2D string array. The first row is the header and the rest are the
// data. Will return an error if the file could not be opened or read.
func readCsv(path string) ([][]string, error) {
	// Open the CSV file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: "+"%s", err)
	}
	defer CloseFile(file)

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %s", err)
	}

	return records, nil
}

// parseCountriesCsv parses the csv file into a Countries struct which is a slice of CountryInfo structs.
// Will return an error if the year or percentage could not be parsed.
func parseCountriesCsv(records [][]string) ([]structs.CountryInfo, error) {
	var countries []structs.CountryInfo

	// Iterate through the records and populate the Countries struct
	for _, record := range records {
		// TODO: Are there better ways to handle the header row?
		// Skip the header row and  all geographical regions
		if record[0] == "Entity" || len(record[1]) != 3 {
			continue
		}

		// Parse the year and percentage
		year, err := strconv.Atoi(record[2])
		if err != nil {
			return countries, fmt.Errorf("error parsing year: %s", err)
		}

		percentage, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			return countries, fmt.Errorf("error parsing percentage: %s", err)
		}

		// Create a CountryInfo struct and append it to the Countries slice
		countryInfo := structs.CountryInfo{
			Country:    record[0],
			IsoCode:    record[1],
			Year:       year,
			Percentage: float32(percentage),
		}
		countries = append(countries, countryInfo)
	}

	return countries, nil
}

// GetCountriesFromCsv reads the historical CSV file and returns a Countries struct.
func GetCountriesFromCsv() ([]structs.CountryInfo, error) {
	// Read the CSV file
	csvData, err := readCsv(constants.HistoricalCsv)
	if err != nil {
		return []structs.CountryInfo{}, err
	}

	// Parse the CSV file
	countries, err := parseCountriesCsv(csvData)
	if err != nil {
		return []structs.CountryInfo{}, err
	}

	return countries, nil
}

// GetCountriesFromJson reads the JSON file and returns a Countries struct.
func GetBordersFromJson() ([]structs.Border, error) {
	// Open the JSON file
	file, err := OpenFile(constants.CountriesJSON)
	if err != nil {
		return []structs.Border{}, err
	}
	defer CloseFile(file)

	// Decode the JSON data
	var borders []structs.Border
	err2 := json.NewDecoder(file).Decode(&borders)
	if err2 != nil {
		return []structs.Border{}, err2
	}

	return borders, nil
}
