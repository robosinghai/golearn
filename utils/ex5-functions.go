package utils

import (
	"log"
)

// Syntax to declare functions in Go
/*
func functionname(parametername type) returntype {
	//function body
   }
*/

var foo = "" // Global variable

// NOTE: functions declared with capital letters are EXPORT functions
func DemoFunction() { // This function can be called from other pkgs
	setString("May ") // NOTE: private/local functions
	//Function that assumes type automatically (value, exists)
	value, exists := setAnotherString("the force be with")

	if !exists { //TIP: do not compare value to a bool const!
		log.Printf("Value is missing")
	} else { // NOTE: function call and return from another function
		log.Println("-----------------------------------------------")
		log.Printf("%s you!", concatString(foo, value))
		log.Println("-----------------------------------------------")
	}
}
func setString(message string) { // Private/Local function
	foo = message
}

func concatString(x string, y string) string { // Private/Local function
	a, b := x, y // String concatenation
	return a + b
}

// Private/Local function
func setAnotherString(name string) (string, bool) {
	return name, true //NOTE: return multiple types
}
