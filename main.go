package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// this struct need refact just like you need
type PersonalInfo struct {
	Name string
	Surname string
	Lastname string
	Sex string
	Age string
	Address string
	City string
	Zipcode string
	County string
	Countrycode string
	Tephone string
}

func main() {
	// Convert string to int
	// Parse this line passed with first argument
	argLine := os.Args[1]
	numLine, _ := strconv.ParseInt(argLine, 0, 64)
	fmt.Println("I will check line:", numLine)

	// Open file if exist
	// This is the second argument after number of line
	argFile := os.Args[2]
	fmt.Println("Start check, wait...")
	csvFile, err := os.Open(argFile)
	if err != nil {
		fmt.Println("Fail to open file", err)
	}

	// Close process after read .csv
	defer csvFile.Close()
	// read all lines with ReadAll()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Err to read file", err)
	}
	// Use ; for delimiter information
	var b [14]string
	for i, j := range strings.Split(csvLines[numLine][0], ";") {
		// fmt.Println(i)
		b[i] = j
	}
	// Get 11 positions of .csv
	s := PersonalInfo{b[0], b[1], b[2],
		b[3], b[4], b[5], b[6],
		b[7], b[8], b[9], b[10],
	}
	v := reflect.ValueOf(s)
	typeOfS := v.Type()
	// imprime cada valor coletado
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

