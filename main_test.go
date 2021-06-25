package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ExampleShowValues() {
	argFile := "data.csv"
	csvFile, err := os.Open(argFile)
	if err != nil {
		fmt.Errorf("error to open file: %v", err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	ShowValues(csvLines, 1)
	// Output:
	// Name: 1a
	// Surname: 2a
	// Lastname: 3a
	// Sex: 4a
	// Age: 5a
	// Address: 6a
	// City: 7a
	// Zipcode: 8a
	// County: 9a
	// CountryCode: 10a
	// Telephone: 11a
}