package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// Define a struct for the XML data
type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	// Open the XML file
	file, err := os.Open("person.xml")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Create a new XML decoder
	decoder := xml.NewDecoder(file)

	// Create a variable to store the decoded data
	var p Person

	// Decode the XML file into the struct
	err = decoder.Decode(&p)
	if err != nil {
		log.Fatal("Error decoding XML:", err)
	}

	// Print the decoded struct
	fmt.Printf("Decoded Person struct: %+v\n", p)
}
