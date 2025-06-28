package main

import (
	"fmt"
)

func main() {
	n := 1000
	batchSize := 30
	maxGoroutines := 4

	chIn := make(chan int)
	done := make(chan struct{}, maxGoroutines)
	donePrinting := make(chan struct{})

	// Printer goroutine
	go func() {
		for prime := range chIn {
			fmt.Println("prime number", prime)
		}
		donePrinting <- struct{}{}
	}()

	batches := (n + batchSize - 1) / batchSize

	for i := 0; i < batches; i++ {
		// Block if maxGoroutines are in use
		done <- struct{}{}

		start := i * batchSize
		end := min(start+batchSize, n)

		go processBatch(start, end, chIn, done)
	}

	// Fill done buffer to ensure all goroutines complete
	for i := 0; i < min(maxGoroutines, batches); i++ {
		done <- struct{}{}
	}

	close(chIn)
	<-donePrinting
	fmt.Println("All numbers have been printed")
}

func processBatch(start, end int, ch chan int, done chan struct{}) {
	defer func() { <-done }()
	for i := start; i < end; i++ {
		if checkPrime(i) {
			ch <- i
		}
	}
}

func checkPrime(num int) bool {
	if num < 2 {
		return false
	}
	for j := 2; j*j <= num; j++ {
		if num%j == 0 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
