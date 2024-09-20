package setup

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"stori/src/types"

	"github.com/gocarina/gocsv"
)

var filename = ""

func ReadFile() *os.File {

	log.Println("READING FILE: " + filename)

	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)

	if err != nil || file == nil {
		log.Fatal("[ERROR] Searching CSV File ", err)
	}

	return file

}

func CSV(id int) chan types.TTransaction {
	if id == 12345 {
		filename = "./src/setup/transactions.csv"
	} else {
		filename = "./src/setup/transactions2.csv"
	}

	readChannel := make(chan types.TTransaction, 1)
	file := ReadFile()
	getCSVRecords(file, readChannel)
	return readChannel
}

func getCSVRecords(file *os.File, c chan types.TTransaction) {
	gocsv.SetCSVReader(func(r io.Reader) gocsv.CSVReader {
		reader := csv.NewReader(r)
		reader.Comma = ','
		reader.LazyQuotes = true
		reader.FieldsPerRecord = -1
		return reader
	})

	go func() {
		err := gocsv.UnmarshalToChan(file, c)
		if err != nil {
			log.Fatal("[ERROR] Reading CSV File Records ", err)
		}
	}()
}
