package csv_coder

import (
	"assignment-2/structs"
	"strconv"
)

/*
DecodeRenewables is a method that takes a csv file and decodes the data into a struct
*/
func DecodeRenewables(csvData [][]string) []structs.Renewables {
	var data []structs.Renewables
	for _, value := range csvData[1:] {
		year, _ := strconv.Atoi(value[2])
		percent, _ := strconv.ParseFloat(value[3], 64)
		data = append(data, structs.Renewables{
			Entity:     value[0],
			Code:       value[1],
			Year:       year,
			Renewables: percent,
		})
	}
	return data
}
