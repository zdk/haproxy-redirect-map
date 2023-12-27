package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/url"
	"os"
)

// ReadFirstTwoColumns reads the first two columns from a CSV file and returns them, URL path-encoded.
func ReadFirstTwoColumns(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var data []string

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(record) >= 2 {
			// URL path-encode and combine the first two columns with a space
			requestPath := url.PathEscape(record[0])
			targetPath := url.PathEscape(record[1])
			combined := requestPath + " " + targetPath
			data = append(data, combined)
		}
	}

	return data, nil
}

// WriteToFile writes the given data to a file.
func WriteToFile(data []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range data {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func main() {
	// Read the data from the original CSV file
	source_filename := "data.csv"
	data, err := ReadFirstTwoColumns(source_filename)
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	} else {
		fmt.Println("Reading from file", source_filename)
	}

	// Write the data to a new file
	target_filename := "out/redirect-list.map"
	err = WriteToFile(data, target_filename)
	if err != nil {
		fmt.Println("Error writing to new file:", err)
		return
	}

	fmt.Println("Data successfully written to", target_filename)
}
