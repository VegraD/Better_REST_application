package utils

import (
	"assignment-2/constants"
	"assignment-2/structs"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func GetCountriesFromCsv() (structs.Countries, error) {
	// Read the CSV file
	csvData, err := readCsv(constants.HistoricalCsv)
	if err != nil {
		return structs.Countries{}, err
	}

	// Parse the CSV file
	countries, err := parseCountriesCsv(csvData)
	if err != nil {
		return structs.Countries{}, err
	}

	return countries, nil
}
