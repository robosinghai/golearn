package main

import (
	"golearn/devices"
	"log"
)

// Demo - Different ways to make arrays in Golang
func MakeArray(arr_length int) []*devices.Device {
	// Four ways to initalise an array

	// Method 1: Meth_1 is assigned an array of type string with values{""}
	// meth_1 := []string{"dev-001", "dev-002", "dev-003"}

	// Method 2: Make an array of type bool and length 5 and assign it to meth_2
	// meth_2 := make([]bool, 5)

	// Method 3: Declars a variable devs array [] of type string
	// var devs []string

	// Method 4: Make an array of type *Device with a slice length of 0 and capacity of 10
	devices := make([]*devices.Device, 5, arr_length) // Note: slice is 0...
	// devices[5] = "dev-005" 		// ... hence this statement will throw an error
	log.Printf("Length of the array when slice is 0: %d",
		len(devices)) // returns 0

	// Reset the slice
	devices = devices[0:5]
	log.Printf("Length of the array after the array is sliced: %d", len(devices)) // returns 5

	return devices
}
