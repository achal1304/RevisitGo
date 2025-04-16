package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware function to log request details
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request method and URL
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call the next handler (the actual route handler)
		next.ServeHTTP(w, r)

		// Log the time it took to process the request
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// Middleware function to check for authentication
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate checking for an authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer valid-token" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Call the next handler (the actual route handler)
		next.ServeHTTP(w, r)
	})
}

// Sample handler to serve a welcome message
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the secure API!")
}

func main() {
	// Create the main handler
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcomeHandler)

	// Apply the middleware to the main handler
	// Using both logging and auth middleware
	handler := loggingMiddleware(authMiddleware(mux))

	// Start the server with the wrapped handler
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", handler)
}
