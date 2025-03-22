package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// 1. Generate a random integer between 0 and 100 (inclusive)
	randomInt := rand.Intn(101)
	fmt.Println("Random integer between 0 and 100:", randomInt)

	// 2. Generate a random floating-point number between 0.0 and 1.0
	randomFloat := rand.Float64()
	fmt.Println("Random float between 0.0 and 1.0:", randomFloat)

	// 3. Generate a random integer between a range (e.g., between 50 and 100)
	randomIntRange := rand.Intn(51) + 50
	fmt.Println("Random integer between 50 and 100:", randomIntRange)
}
