package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 34, 5, 1, 1234, 34523, 5212}

	var wg sync.WaitGroup
	chIn := make(chan int)
	chOut := make(chan int)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go processor(chIn, chOut, &wg)
	}

	go func() {
		for output := range chOut {
			fmt.Println(output)
		}
	}()

	for _, num := range nums {
		chIn <- num
	}

	close(chIn)
	wg.Wait()
	close(chOut)
}

func processor(chIn, chOut chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for ele := range chIn {
		chOut <- ele * 2
	}
}
