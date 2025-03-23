package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel to receive heartbeat signals
	heartbeat := make(chan bool)
	// Timeout duration
	timeout := 6 * time.Second
	timer := time.NewTimer(timeout)

	// Start a goroutine that sends heartbeat signals every second
	go sendHeartbeat(heartbeat)

	// Wait for a heartbeat or timeout
	for {
		select {
		case <-heartbeat:
			// If heartbeat received within timeout
			fmt.Println("Received heartbeat signal.")
		case <-timer.C:
			// If no heartbeat received within the timeout
			fmt.Println("Timeout: No heartbeat received.")
			return
		}
	}
}

// sendHeartbeat sends a heartbeat signal every second
func sendHeartbeat(heartbeat chan bool) {
	for {
		// Simulate doing some work by sleeping for 1 second
		time.Sleep(1 * time.Second)
		// Send a heartbeat signal
		heartbeat <- true
	}
}
