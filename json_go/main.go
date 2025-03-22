package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Define a custom type for handling dates
type Date struct {
	Day   int
	Month int
	Year  int
}

// Implement the MarshalJSON method for custom encoding
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%02d-%02d-%04d", d.Day, d.Month, d.Year))
}

// Implement the UnmarshalJSON method for custom decoding
func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	// Parse the date string (assumes format DD-MM-YYYY)
	_, err := fmt.Sscanf(dateStr, "%02d-%02d-%04d", &d.Day, &d.Month, &d.Year)
	return err
}

func main() {
	// Example 1: Encoding a struct to JSON
	p := Person{Name: "Alice", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encoded Person struct:", string(jsonData))

	// Example 2: Pretty printing JSON
	jsonPretty, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nPretty Printed JSON:")
	fmt.Println(string(jsonPretty))

	// Example 3: Decoding JSON back into a struct
	jsonDataToDecode := []byte(`{"name": "Bob", "age": 25}`)
	var decodedPerson Person
	err = json.Unmarshal(jsonDataToDecode, &decodedPerson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nDecoded Person struct: %+v\n", decodedPerson)

	// Example 4: Working with custom types (Date)
	d := Date{Day: 25, Month: 12, Year: 2022}
	dateJSON, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nEncoded Date:", string(dateJSON))

	// Example 5: Decoding custom JSON Date format
	var newDate Date
	err = json.Unmarshal([]byte(`"25-12-2022"`), &newDate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nDecoded Date struct: %+v\n", newDate)

	// Example 6: Handling JSON with Null values
	jsonDataWithNull := []byte(`{"name": "Alice", "age": null}`)
	var personWithNullAge struct {
		Name string `json:"name"`
		Age  *int   `json:"age"` // Using pointer to handle null values
	}
	err = json.Unmarshal(jsonDataWithNull, &personWithNullAge)
	if err != nil {
		log.Fatal(err)
	}
	if personWithNullAge.Age == nil {
		fmt.Println("\nAge is null")
	} else {
		fmt.Printf("Age: %d\n", *personWithNullAge.Age)
	}
}
