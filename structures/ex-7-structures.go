package main

import (
	"log"
)

func StructFunc() {
	// Initialisation and assignment (value) in a single step
	data := &Device{ID: "dev-001", Reading: 50.0}
	data.addBias(10.0) // With associated methods, the code reads better
	log.Printf("Device ID: %s\t Reading: %f\n", data.ID, data.Reading)
}

// The below simply means:
// .. the type Device (struct) is a receiver of a method called AddBias
// Tip: If a method is NOT altering or changing struct variables declare
// them as generic methods.
func (d *Device) addBias(bias float64) {
	d.Reading += bias
}
