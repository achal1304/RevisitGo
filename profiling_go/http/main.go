package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import the pprof package for profiling
	"time"
)

func main() {
	// Start HTTP server in a goroutine
	go func() {
		log.Println("Starting pprof server on port 6060...")
		log.Fatal(http.ListenAndServe("localhost:6060", nil)) // pprof server runs on http://localhost:6060
	}()

	// Regular HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go Profiler!")
		time.Sleep(2 * time.Second) // Simulate some work
	})

	// Start the main HTTP server
	log.Println("Starting main server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
