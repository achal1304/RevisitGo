package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open (or create) the CSV file for writing
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write the header row
	header := []string{"Name", "Age", "City"}
	err = writer.Write(header)
	if err != nil {
		log.Fatal("Error writing header:", err)
	}

	// Write some data rows
	data := [][]string{
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
		{"Charlie", "35", "Chicago"},
	}

	// Write multiple records at once
	err = writer.WriteAll(data)
	if err != nil {
		log.Fatal("Error writing data:", err)
	}

	// Flush the data to the file
	writer.Flush()

	// Check for any error encountered during the flush
	if err := writer.Error(); err != nil {
		log.Fatal("Error flushing data to file:", err)
	}

	fmt.Println("Data written to CSV successfully!")
}
