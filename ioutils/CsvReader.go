package ioutils

import (
	"encoding/csv"
	"log"
	"os"
)

// ReadCsvFile reads a CSV file and returns an array of Ham/Spam and text
func ReadCsvFile(path string) (record [][]string) {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error while opening file: ", err)
		return
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	record, err = reader.ReadAll()
	if err != nil {
		log.Fatalln("Error while reading the file: ", err)
	}

	return
}
