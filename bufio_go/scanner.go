package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanner_go() {
	// bufio.Scanner usage example

	// Open a file to read (for demonstration purposes)
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Reading the file line by line
	fmt.Println("Reading file line by line using bufio.Scanner:")
	for scanner.Scan() {
		// Get each line from the scanner
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check if there was an error during the scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Using bufio.Scanner to read from standard input (os.Stdin)
	fmt.Print("\nEnter something: ")
	scanner = bufio.NewScanner(os.Stdin)

	// Read a line from stdin
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Println("You entered:", input)
	}

	// Check if there was an error during the scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
