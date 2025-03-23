package main

import (
	"fmt"
)

func sendData(ch chan int, data []int) {
	for _, num := range data {
		ch <- num
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendData(ch1, []int{1, 2, 3})
	go sendData(ch2, []int{4, 5, 6})

	for {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println("Received from ch1:", msg1)
			}
		case msg2, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println("Received from ch2:", msg2)
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
