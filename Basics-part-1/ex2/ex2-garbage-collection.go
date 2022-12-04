package main // Interface

import "fmt"          // Standard library
var name string = ""  // Explicit assignment (Global)
var tail string = ""  // Explicit assignment (Global)
var tail_1 = ".co.uk" // Implicit assignment as string (Global)

// Variable assignemnt
func main() {
	var name = "robosingh"
	tail = ".com"
	fmt.Printf("Hello %s%s\n", name, tail)
	variableRedefined()
} // Garbage collected

func variableRedefined() {
	var name1 = "robosingh"
	fmt.Printf("Hello again %s%s\n", name, tail_1)  // Global pref
	fmt.Printf("Hello again %s%s\n", name1, tail_1) // Local pref
} // Garbage collected
