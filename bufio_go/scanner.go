package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

func scanner_terminate() {
	// Create a new scanner for reading from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Set up a signal channel to catch Ctrl+C (SIGINT)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	// Prompt the user to enter something
	fmt.Println("Please enter some text (Ctrl+C to exit):")

	// Start a goroutine to listen for signals
	go func() {
		<-sigs
		fmt.Println("\nExiting program...")
		os.Exit(0)
	}()

	// Use scanner to read from stdin
	for scanner.Scan() {
		// Read the line
		input := scanner.Text()

		// If input is a non-empty line, print it
		if len(input) > 0 {
			fmt.Printf("You entered: %s\n", input)
		} else {
			// Exit loop if input is empty (optional)
			fmt.Println("Exiting program...")
			break
		}
	}

	// Check for errors during the scanning process
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
	}
}
