package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strings"
)

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

	for _, i:= range strings.Split(csvLines[0][0], ";") {
		fmt.Println(i)
	}
}

