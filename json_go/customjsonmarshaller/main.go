package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Person struct
type Person struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

// Custom MarshalJSON method
func (p *Person) MarshalJSON() ([]byte, error) {
	// Custom format for DateOfBirth field: YYYY-MM-DD
	type Alias Person // Create an alias to avoid infinite recursion
	return json.Marshal(&struct {
		DateOfBirth string `json:"date_of_birth"`
		*Alias
	}{
		DateOfBirth: p.DateOfBirth.Format("2006-01-02"), // Format the time
		Alias:       (*Alias)(p),
	})
}

// Custom UnmarshalJSON method
func (p *Person) UnmarshalJSON(data []byte) error {
	// Define a custom struct with DateOfBirth as string
	type Alias Person
	aux := &struct {
		DateOfBirth string `json:"date_of_birth"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	// Unmarshal into the auxiliary struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Custom date format: accepting MM-DD-YYYY
	parsedDate, err := time.Parse("01-02-2006", aux.DateOfBirth)
	if err != nil {
		// Fall back to a default format (YYYY-MM-DD)
		parsedDate, err = time.Parse("2006-01-02", aux.DateOfBirth)
	}
	if err != nil {
		return err
	}

	p.DateOfBirth = parsedDate
	return nil
}

func main() {
	// Create a Person struct
	person := &Person{
		Name:        "Alice",
		DateOfBirth: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
	}

	// Marshal the struct into JSON with custom formatting
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("JSON Marshal:", string(data))

	// Example JSON with a custom date format for unmarshaling
	jsonData := `{"name":"Alice","date_of_birth":"03-15-1990"}`

	// Unmarshal the JSON data back into a Person struct
	var person2 Person
	if err := json.Unmarshal([]byte(jsonData), &person2); err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	// Print the unmarshaled Person struct
	fmt.Printf("Person struct: %+v\n", person2)
}
