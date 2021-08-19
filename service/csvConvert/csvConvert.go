package csvConvert

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

// OpenCSV read file csv and return slice of string
func OpenCSV(file string) ([]string, error) {
	// open file
	csvFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			log.Panic(err)
		}
	}(csvFile)

	// Read line by line
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	var labels []string

	//fmt.Println(csvLines)
	for _, lines := range csvLines {
		//fmt.Println(i, lines)
		for _, col := range lines {
			//fmt.Println(j, col)
			for _, v := range strings.Split(col, ";") {
				//fmt.Println(k, v)
				labels = append(labels, v)
			}
		}
	}
	return labels, nil
}
