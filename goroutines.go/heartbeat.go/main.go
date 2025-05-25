package main

import (
	"fmt"
	"time"
)

func main() {
	heartbeat := make(chan bool)
	timer := time.NewTimer(6 * time.Second)

	go sendHeartbeat(heartbeat)

	for {
		select {
		case <-heartbeat:
			fmt.Println("received heartbeat")
		case <-timer.C:
			fmt.Println("time is up, stopping the code")
			close(heartbeat)
			return
		}
	}
}

// sendHeartbeat sends a heartbeat signal every second
func sendHeartbeat(heartbeat chan bool) {
	for {
		time.Sleep(1 * time.Second)
		heartbeat <- true
	}
}
