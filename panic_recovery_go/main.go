package main

import "fmt"

// Function that triggers a panic
func triggerPanic() {
	fmt.Println("About to panic!")
	panic("Something went wrong!")
}

// Function that recovers from the panic
func safeFunction() {
	// Deferring the recovery function to catch the panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	triggerPanic()
}

func main() {
	// Calling safeFunction, which will handle the panic
	fmt.Println("Starting safeFunction...")
	safeFunction()
	fmt.Println("Execution continues after recovery.")
}
