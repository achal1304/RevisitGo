package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID     int
	Number int
}

type Result struct {
	JobID  int
	Output int
}

func worker(id int, result chan Result, jobs chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("started worker ", id)
		time.Sleep(time.Millisecond * 500)
		result <- Result{JobID: job.ID, Output: job.Number * 2}
		fmt.Println("stopped worker ", id)
	}
}

func main() {
	jobs := make(chan Job, 10)
	result := make(chan Result, 10)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		i = i
		wg.Add(1)
		go worker(i, result, jobs, &wg)
	}

	go func() {
		for i := 0; i < 10; i++ {
			jobs <- Job{ID: i, Number: i * 3}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println("Result from job ", res)
	}
}
