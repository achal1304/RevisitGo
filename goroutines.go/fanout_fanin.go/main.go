package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 34, 5, 1, 1234, 34523, 5212}

	var wg sync.WaitGroup
	chIn := make(chan int, len(nums))
	chOut := make(chan int, len(nums))
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go processor(chIn, chOut, &wg)
	}

	go func() {
		for output := range chOut {
			fmt.Println(output)
		}
		done <- struct{}{}
	}()

	for _, num := range nums {
		chIn <- num
	}

	close(chIn)
	wg.Wait()
	close(chOut)
	<-done
}

func processor(chIn, chOut chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for ele := range chIn {
		chOut <- ele * 2
	}
}
