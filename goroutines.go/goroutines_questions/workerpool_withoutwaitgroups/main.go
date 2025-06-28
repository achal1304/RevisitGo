package main

import (
	"fmt"
	"time"
)

// Worker pool pattern without using waitgroups

func main() {
	inputCh := make(chan int)
	maxGoroutines := 10
	done := make(chan struct{}, maxGoroutines)

	go producer(inputCh, 1000)

	for data := range inputCh {
		done <- struct{}{}
		go consumer(data, done)
	}
}

func consumer(num int, done chan struct{}) {
	fmt.Println("Consuming num", num)
	time.Sleep(4 * time.Second)
	fmt.Println("Consumption completed for num ", num)
	defer func() {
		<-done
	}()
}

func producer(inputCh chan int, maxVal int) {
	for i := 0; i < maxVal; i++ {
		inputCh <- i
	}
}
