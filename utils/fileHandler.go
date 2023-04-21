package utils

import (
	"assignment-2/constants"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// CloseFile is a wrapper function for closing a file. Could be used with defer for cleaner code.
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Printf(constants.CloseFileFail+"%s", err)
	}
}

// ReadCsv reads a CSV file and returns the rows as a 2D string array. The first row is the header and the rest are the
// data. Will return an error if the file could not be opened or read.
func ReadCsv(path string) ([][]string, error) {
	// Open the CSV file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(constants.OpenFileFail+"%s", err)
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
