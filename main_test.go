package main

import (
	"os"
	"strings"
	"testing"
)

func TestReadFirstTwoColumns(t *testing.T) {
	expected := []string{
		"Column1Row1 Column2Row1",
		"Column1Row2 Column2Row2",
		"%E0%B8%97%E0%B8%94%E0%B8%AA%E0%B8%AD%E0%B8%9A %E0%B9%84%E0%B8%97%E0%B8%A2",
	}
	result, err := ReadFirstTwoColumns("test_data/data.csv")
	if err != nil {
		t.Fatalf("Error reading CSV file: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d records, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %q, got %q at index %d", expected[i], v, i)
		}
	}
}

func TestWriteToFile(t *testing.T) {
	data := []string{"data1", "data2"}
	filename := "testoutput.txt"

	// Write data to file
	err := WriteToFile(data, filename)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

	// Read back the data
	file, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error reading back file: %v", err)
	}

	// Convert read data to string and split by newline
	readData := strings.Split(string(file), "\n")

	// The file read might have an extra empty line at the end
	if len(readData) > len(data) {
		readData = readData[:len(data)]
	}

	// Compare with original data
	for i, line := range readData {
		if line != data[i] {
			t.Errorf("Expected %q, got %q at line %d", data[i], line, i)
		}
	}

	// Clean up: delete the test output file
	os.Remove(filename)
}
