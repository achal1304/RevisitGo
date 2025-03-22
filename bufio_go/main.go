package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// bufio package usage example
	// Reading input from the user

	// Creating a new reader to read from standard input (os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	// Prompting the user
	fmt.Print("Enter your name: ")

	// Reading the input from the user (reads until a newline)
	name, _ := reader.ReadString('\n')

	// Trimming the newline character from the input
	name = strings.TrimSpace(name)

	// Displaying the user's input
	fmt.Printf("Hello, %s! Welcome to Go.\n", name)

	// os package usage examples

	// Checking if a file exists using os.Stat
	filename := "sample.txt"
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		fmt.Printf("The file %s does not exist.\n", filename)
	} else if err != nil {
		// Handle any other error
		fmt.Printf("Error checking file: %v\n", err)
	} else {
		fmt.Printf("The file %s exists.\n", filename)
	}

	// Creating and writing to a file using os.Create
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Writing data to the file using fmt.Fprintln
	fmt.Fprintln(file, "This is a test file created by Go.")

	// Using bufio.Writer to write to the file with buffering
	writer := bufio.NewWriter(file)
	writer.WriteString("This is some buffered data.\n")
	writer.Flush() // Ensure all data is written to the file

	// Reading a file using bufio
	fileToRead, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file for reading:", err)
		return
	}
	defer fileToRead.Close()

	// Creating a buffered reader to read the file
	bufReader := bufio.NewReader(fileToRead)

	// Reading the first line of the file
	line, err := bufReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("First line of file:", line)
	}
}
