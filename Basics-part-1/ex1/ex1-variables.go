package main // Interface

import "fmt" // Standard library

func main() {
	//NOTE: By default, GO assigns a zero value to variables
	var name string = "robosingh" // Explicit assignment
	var tail_1 = ".co.uk"         // Implicit assignment (as string)
	fmt.Printf("Strings: %s%s\n", name, tail_1)

	// var keyword can be dropped. := infers the type automatically
	sens_val := getSensorReading() // Implicit assignment
	fmt.Printf("Sensor reading: %d\n", sens_val)
}

// Implicit assignment from a function (returning int)
func getSensorReading() int {
	// sensor reading process
	return 35
}
