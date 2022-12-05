package main // Package name

import (
	"fmt"  // Standard library
	"math" // Math library
)

func main() {

	// MaxInt64 is an exported name
	fmt.Println("Max value of int64: ", int64(math.MaxInt64))

	// Phi is an exported name
	fmt.Println("Value of Phi (ϕ): ", math.Phi)

	// pi starts with a small letter, so it is not exported
	// fmt.Println("Value of Pi (𝜋): ", math.pi) // ERROR
	fmt.Println("Value of Pi (𝜋): ", math.Pi) // CORRECT
}

func SensorReading() float64 {
	return math.Pi
}
