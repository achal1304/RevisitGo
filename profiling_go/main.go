package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

// Function that simulates CPU-intensive work
func doWork() {
	var sum int
	for i := 0; i < 10000000; i++ {
		sum += i
	}
	fmt.Println("Work done:", sum)
}

func doMemoryWork() []int {
	// Allocate a large slice and return it
	data := make([]int, 1000000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	return data
}

func main() {
	profFile, err := os.Create("cpu_profile.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer profFile.Close()

	memFile, err := os.Create("memory_profile.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer memFile.Close()

	if err := pprof.StartCPUProfile(profFile); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	if err := pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		doWork()
		doMemoryWork()
		time.Sleep(1 * time.Second)
	}
}
