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
// Blank1___ -> Name string
// Blank2___ -> Years int
// Blank3___ -> Address string 
type myData struct {
	Blank1______ string
	Blank2______ string
	Blank3______ string
	Blank4o_____ string
	Blank5o_____ string
	Blank6o_____ string
	Blank7______ string
	Blank8______ string
	Blank9o_____ string
	Blank10_____ string
	Blank11o____ string
	Blank12o____ string
	Blank13_____ string
	Blank14_____ string
}

func main() {
	// Converte o parametro de texto para inteiro
	argLine := os.Args[1]
	numLine, _ := strconv.ParseInt(argLine, 0, 64)
	fmt.Println("I will check line:", numLine)

	// Abre o arquivo caso exista
	argFile := os.Args[2]
	fmt.Println("Start check, wait...")
	csvFile, err := os.Open(argFile)
	if err != nil {
		fmt.Println("Fail to open file", err)
	}

	// garante sucesso para encerrar após leitura
	defer csvFile.Close()
	// le todas as linhas do arquivo
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Err to read file", err)
	}
	// alimenta a array com os dados do csv, delimitado por ";"
	var b [14]string
	for i, j := range strings.Split(csvLines[numLine][0], ";") {
		// fmt.Println(i)
		b[i] = j
	}
	// pega os 14 campos do csv
	s := myData{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13]}
	v := reflect.ValueOf(s)
	typeOfS := v.Type()
	// imprime cada valor coletado
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}

}
