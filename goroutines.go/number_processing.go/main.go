package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 4, 3, 6, 8, 111, 2323, 5555}
	ch := make(chan int, len(nums))
	done := make(chan struct{})
	var wg sync.WaitGroup

	for _, n := range nums {
		wg.Add(1)
		go square(n, ch, &wg)
	}

	go func() {
		for ele := range ch {
			fmt.Println(ele)
		}
		done <- struct{}{}
	}()
	wg.Wait()
	close(ch)
	<-done
}

func square(n int, ch chan int, wg *sync.WaitGroup) {
	ch <- n * n
	wg.Done()
}
