package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func csv_go() {
	// Open the CSV file
	file, err := os.Open("people.csv")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all rows from the CSV file
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
	}

	// Process each row
	for _, row := range rows {
		// Assuming each row has two columns: Name and Age
		name := row[0]
		age := row[1]
		fmt.Printf("Name: %s, Age: %s\n", name, age)
	}
}
