package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Basic arithmetic functions
	fmt.Println("Square Root of 16:", math.Sqrt(16))         // Square root
	fmt.Println("Power of 2^3:", math.Pow(2, 3))             // Exponentiation
	fmt.Println("Absolute value of -3.14:", math.Abs(-3.14)) // Absolute value
	fmt.Println("Ceiling of 3.14:", math.Ceil(3.14))         // Round up to nearest integer
	fmt.Println("Floor of 3.14:", math.Floor(3.14))          // Round down to nearest integer

	// 2. Trigonometric functions
	// math.Sin, math.Cos, math.Tan expects radians as input
	fmt.Println("Sine of Pi/2:", math.Sin(math.Pi/2))    // Sin function
	fmt.Println("Cosine of Pi/2:", math.Cos(math.Pi/2))  // Cosine function
	fmt.Println("Tangent of Pi/4:", math.Tan(math.Pi/4)) // Tangent function

	// 3. Constants
	fmt.Println("Pi:", math.Pi)              // Pi constant
	fmt.Println("Euler's constant:", math.E) // Euler's number

	// 4. Logarithmic functions
	fmt.Println("Natural Log of 10:", math.Log(10))     // Natural logarithm (base e)
	fmt.Println("Log base 2 of 8:", math.Log2(8))       // Logarithm base 2
	fmt.Println("Log base 10 of 100:", math.Log10(100)) // Logarithm base 10

	// 5. Rounding functions
	fmt.Println("Round 3.14159:", math.Round(3.14159))      // Rounding to the nearest integer
	fmt.Println("Truncating 3.14159:", math.Trunc(3.14159)) // Removing the fractional part

	// 6. Hypotenuse (Pythagorean theorem)
	fmt.Println("Hypotenuse of sides 3 and 4:", math.Hypot(3, 4)) // Hypotenuse of a right triangle

	// 7. Conversion between radians and degrees
	degrees := 90.0
	radians := degrees * math.Pi / 180
	fmt.Printf("%v degrees in radians: %v\n", degrees, radians)
}
