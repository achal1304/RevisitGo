package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct to represent data in JSON format
type Response struct {
	Message string `json:"message"`
}

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header content type
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// Handler for the about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header content type
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is the About Page!")
}

// Handler for the contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header content type
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is the Contact Page!")
}

// Handler with URL variables (for dynamic routing)
func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get userID from URL variables
	vars := mux.Vars(r)
	userID := vars["userID"]

	// Set the response header content type
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	// Respond with user profile information
	fmt.Fprintf(w, "User Profile for User ID: %s", userID)
}

// Handler for JSON response with query parameter
func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Get query parameter 'name'
	name := r.URL.Query().Get("name")

	// Prepare the response struct
	response := Response{
		Message: "Hello, " + name,
	}

	// Encode and send the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// Main function to start the HTTP server with Gorilla Mux router
func main() {
	// Start the HTTP server on port 8080
	fmt.Println("Server is running on http://localhost:8080/")
	err := http.ListenAndServe(":8080", router())
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func router() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", homeHandler).Methods("Get")             // Route for "/"
	r.HandleFunc("/about", aboutHandler)                      // Route for "/about"
	r.HandleFunc("/contact", contactHandler)                  // Route for "/contact"
	r.HandleFunc("/user/{userID:[0-9]+}", userProfileHandler) // Dynamic route with userID (only digits)
	r.HandleFunc("/json", jsonResponseHandler)                // Route for JSON response

	return r
}
