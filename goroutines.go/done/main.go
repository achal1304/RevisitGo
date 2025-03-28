package main

import (
	"fmt"
	"time"
)

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
	done := make(chan struct{})

	go longRunningTask(1, done)
	go longRunningTask(2, done)

	time.Sleep(2 * time.Second)

	fmt.Println("Main: Sending done signal...")
	close(done)

	time.Sleep(1 * time.Second)
	fmt.Println("Main: Exiting")
}
