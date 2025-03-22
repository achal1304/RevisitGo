package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Getting the current time
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	// 2. Formatting the current time
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted current time:", formattedTime)

	// 3. Parsing a time string into a time.Time object
	timeStr := "2025-03-22 14:30:00"
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time:", parsedTime.Year(), parsedTime.Day())

	// 4. Performing time arithmetic (Adding/Subtracting time)
	// Adding 2 hours to the current time
	plusTwoHours := currentTime.Add(2 * time.Hour)
	fmt.Println("Current time + 2 hours:", plusTwoHours)

	// Subtracting 30 minutes from the current time
	minusThirtyMinutes := currentTime.Add(-30 * time.Minute)
	fmt.Println("Current time - 30 minutes:", minusThirtyMinutes)

	// 5. Calculating the difference between two times
	duration := plusTwoHours.Sub(currentTime)
	fmt.Println("Difference between the two times (2 hours later):", duration)

	// 6. Time-based operations using durations
	// Sleep for 2 seconds
	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Woke up after sleeping for 2 seconds!")
}
