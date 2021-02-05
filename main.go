package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Student struct {
	Blank1_____ string
	Blank2_____ string
	Blank3_____ string
	Blank4_____ string
	Blank5_____ string
	Blank6_____ string
	Blank7_____ string
	Blank8_____ string
	Blank9_____ string
	Blank10____ string
	Blank11____ string
	Blank12____ string
	Blank13____ string
	Blank14____ string
}

func main() {

	fmt.Println("Começando")
	csvFile, err := os.Open("lista.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Sucesso ao abrir o arquivo")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var j [14]string
	for b, i := range strings.Split(csvLines[1][0], ";") {
		// fmt.Println(i)
		j[b] = i
	}
	s := Student{j[0], j[1], j[2], j[3], j[4], j[5], j[6], j[7], j[8], j[9], j[10], j[11], j[12], j[13]}
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}

}
