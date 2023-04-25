package responses

import (
	"assignment-2/utils"
	"bytes"
	"encoding/json"
	"io"
	"os"
)

// FormatJSON This is just for formatting the JSON data in the responses folder. It is not used in the program.
func FormatJSON() error {
	// Read JSON data from file
	file, err := utils.OpenFile("responses/countries.json")
	if err != nil {
		panic(err)
	}
	defer utils.CloseFile(file)

	// Read the JSON data
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var formattedData bytes.Buffer

	// Format the JSON data with indentation
	err = json.Indent(&formattedData, data, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("responses/countries_formatted.json", formattedData.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	return nil
}
