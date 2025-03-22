package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// Handler for the about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page!")
}

// Handler for the contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Contact Page!")
}

// Handler with URL variables (for dynamic routing)
func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	fmt.Fprintf(w, "User Profile for User ID: %s", userID)
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

	r.HandleFunc("/", homeHandler)                            // Route for "/"
	r.HandleFunc("/about", aboutHandler)                      // Route for "/about"
	r.HandleFunc("/contact", contactHandler)                  // Route for "/contact"
	r.HandleFunc("/user/{userID:[0-9]+}", userProfileHandler) // Dynamic route with userID (only digits)

	return r
}
