package main

import (
	"fmt"
	"time"
)

// longRunningTask simulates a background process that listens for cancellation
func longRunningTask(id int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Printf("Goroutine %d received done signal. Exiting.\n", id)
			return
		default:
			fmt.Printf("Goroutine %d is working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// Create a done channel to signal cancellation
	done := make(chan struct{})

	// Start two long-running goroutines
	go longRunningTask(1, done)
	go longRunningTask(2, done)

	// Let them run for 2 seconds
	time.Sleep(2 * time.Second)

	// Cancel all goroutines by closing the done channel
	fmt.Println("Main: Sending done signal...")
	close(done)

	// Give goroutines a moment to shut down
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Exiting")
}
