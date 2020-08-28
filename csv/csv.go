package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

func Lookup(file *os.File, fieldName string, targetValue string) ([][]string, error) {
	var fieldIndex int = -1
	var returnValues [][]string

	csvReader := csv.NewReader((io.Reader)(file))

	record, err := csvReader.Read()
	if err != nil {
		panic(err)
	}

	// find index of `fieldName`, using first record
	for i := 0; i < len(record); i++ {
		if fieldName == record[i] {
			fieldIndex = i
		}
	}

	// make sure fieldIndex exists
	if fieldIndex == -1 {
		return nil, errors.New("Name of field not found")
	}

	records, err := csvReader.ReadAll()

	// iterate over all records, return ones that match the value
	for _, record := range records {
		if targetValue == record[fieldIndex] {
			returnValues = append(returnValues, record)
		}
	}

	if len(returnValues) == 0 {
		return nil, errors.New("csv.go: Value not found")
	}

	return returnValues, nil
}
