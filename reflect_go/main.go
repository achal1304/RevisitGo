package main

import (
	"fmt"
	"reflect"
)

// Define two sample structs for comparison
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

// Function to compare two structs using reflect
func compareStructs(a, b interface{}) bool {
	// Use reflect.DeepEqual to compare values
	return reflect.DeepEqual(a, b)
}

func main() {
	// Example 1: Identifying the type of a value
	var x int = 42
	var y float64 = 3.14
	var z string = "hello"

	fmt.Println("Type of x:", reflect.TypeOf(x)) // int
	fmt.Println("Type of y:", reflect.TypeOf(y)) // float64
	fmt.Println("Type of z:", reflect.TypeOf(z)) // string

	// Example 2: Comparing two structs (using reflect.DeepEqual)
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Alice", Age: 30}
	p3 := Person{Name: "Bob", Age: 25}

	fmt.Println("p1 and p2 are equal:", compareStructs(p1, p2)) // true
	fmt.Println("p1 and p3 are equal:", compareStructs(p1, p3)) // false

	// Example 3: Accessing struct fields dynamically
	e := Employee{Name: "John", Age: 40, Salary: 50000.0}

	// Get the value of the struct (Employee) using reflection
	val := reflect.ValueOf(e)

	// Iterate over struct fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Printf("Field Name: %s, Value: %v, Type: %s\n", val.Type().Field(i).Name, field.Interface(), field.Type())
	}

	// Example 4: Modifying a struct field dynamically (using reflect)
	// Since `e` is not a pointer, this will cause an error. So, we'll work with a pointer to `e`.
	ptr := reflect.ValueOf(&e) // Get a pointer to the Employee struct
	field := ptr.Elem().FieldByName("Salary")

	if field.CanSet() {
		// Modify the Salary field dynamically
		field.SetFloat(55000.0)
	}

	// After modification
	fmt.Printf("Updated Employee: %+v\n", e) // Output: Updated Employee: {Name:John Age:40 Salary:55000}
}
