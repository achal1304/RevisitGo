package main

import "fmt"

func main() {
	max := 1000
	evenTurn := make(chan struct{})
	oddTurn := make(chan struct{})
	done := make(chan struct{})

	// Even printer
	go func() {
		for i := 0; i < max; i += 2 {
			<-evenTurn
			fmt.Println(i)
			oddTurn <- struct{}{}
		}
	}()

	// Odd printer
	go func() {
		for i := 1; i < max; i += 2 {
			<-oddTurn
			fmt.Println(i)
			if i+1 < max {
				evenTurn <- struct{}{}
			}
		}
		done <- struct{}{}
	}()

	// Start with even
	evenTurn <- struct{}{}
	<-done // wait for odd to finish
}
