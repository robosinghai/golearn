package main

import (
	"log"
)

// Check out golearn/devices package for a better way to declear
// structs and associated functions.
type Device struct {
	ID      string
	Reading float64
}

// Generate the exe:
//	from within the folder structures, run:
//	go build . && ./structures

func main() {
	data := &Device{ID: "dev-001", Reading: 50.0}
	log.Printf("Device ID: %s\t Reading: %f\n", data.ID, data.Reading)

	AddBias(data)
	log.Printf("Device ID: %s\t Reading: %f\n", data.ID, data.Reading)

	//calling ex-7 see ex-7-structures.go
	StructFunc()

	//calling ex-8 see ex-8-new-function
	NewFuncExample()

}

func AddBias(d *Device) {
	bias := 10.5
	d.Reading += bias
}

func (d *Device) AddBias() {
	bias := 10.5
	d.Reading += bias
}
