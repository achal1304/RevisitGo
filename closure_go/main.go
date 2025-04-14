package main

import "fmt"

// Function to generate a sequence of numbers
func generateSequence(start int) func() int {
	// Variable to store the current number
	current := start

	// Closure that generates the next number in the sequence
	return func() int {
		current++
		return current
	}
}

func main() {
	// Create a sequence generator starting from 10
	seq := generateSequence(10)

	// Generate the next 5 numbers
	for i := 0; i < 5; i++ {
		fmt.Println(seq()) // Output: 11, 12, 13, 14, 15
	}
}
