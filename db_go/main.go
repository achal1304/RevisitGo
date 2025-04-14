package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Define connection parameters
	connStr := "user=yourusername dbname=yourdbname sslmode=disable password=yourpassword host=localhost port=5432"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Check if the connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Connection successful
	fmt.Println("Successfully connected to the database!")

	// Example query
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	// Iterate over rows and display data
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Check for any error after iterating over the rows
	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}
}
