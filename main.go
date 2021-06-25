package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// PersonalInfo you can change this for what you need.
type PersonalInfo struct {
	Name        string
	Surname     string
	Lastname    string
	Sex         string
	Age         string
	Address     string
	City        string
	Zipcode     string
	County      string
	CountryCode string
	Telephone   string
}

// GetArguments obtain values of parameters passed for binary LINE FILE
func GetArguments() error {
	argLine := os.Args[1]
	// Get INT to read one line of file CSV
	numLine, _ := strconv.ParseInt(argLine, 0, 64)

	// Get PATH of file
	argFile := os.Args[2]

	ConvertCSVtoStruct(argFile, numLine)

	return nil
}

// ConvertCSVtoStruct get information of file and convert to Struct for Go can understand
func ConvertCSVtoStruct(argFile string, numLine int64) error {
	csvFile, err := os.Open(argFile)
	if err != nil {
		fmt.Errorf("error to open file: %v", err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Errorf("error to read file: %v", err)
	}
	fmt.Printf("\nValor: %v \t Tipo: %T \n", csvLines, csvLines)

	ShowValues(csvLines, numLine)
	return nil
}

// ShowValues show values of conversion
func ShowValues(csvLines [][]string, numLine int64) (string, error) {
	// Use ; for delimiter information
	var b [20]string
	for i, j := range strings.Split(csvLines[numLine][0], ";") {
		b[i] = j
	}
	// Get 11 positions of .csv
	s := PersonalInfo{b[0], b[1], b[2],
		b[3], b[4], b[5], b[6],
		b[7], b[8], b[9], b[10],
	}

	v := reflect.ValueOf(s)

	typeOfS := v.Type()
	// print each value
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
	return "ok", nil
}

func main() {
	GetArguments()
}
