package main

import "fmt"

// Fibonacci function - iterative implementation
func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Fibonacci function - recursive implementation
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func stringConcatenate(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "A" // This creates new memory allocations each time
	}
	return result
}


func main() {
	n := 10

	// Test both implementations
	fmt.Printf("Fibonacci (Iterative) of %d: %d\n", n, fibonacciIterative(n))
	fmt.Printf("Fibonacci (Recursive) of %d: %d\n", n, fibonacciRecursive(n))
}
