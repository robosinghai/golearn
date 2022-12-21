package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Demo - Different ways to make arrays in Golang
func MakeArray(arr_length int) []*Device {
	// Four ways to initalise an array

	// Method 1: Meth_1 is assigned an array of type string with values{""}
	// meth_1 := []string{"dev-001", "dev-002", "dev-003"}

	// Method 2: Make an array of type bool and length 5 and assign it to meth_2
	// meth_2 := make([]bool, 5)

	// Method 3: Declars a variable devs array [] of type string
	// var devs []string

	// Method 4: Make an array of type *Device with a slice length of 0 and capacity of 10
	devices := make([]*Device, 0, arr_length) // Note: slice is 0...
	// devices[5] = "dev-005" 		// ... hence this statement will throw an error

	// Reset the slice
	devices = devices[0:5]
	log.Printf("Length of the array: %d", len(devices))

	return devices
}

// Demo - Expand a slice of an array and add append to it
// Note - The slice is expanded in the MakeArray func ^ ^
func (registry *Registry) AddDeviceToRegistry(dev *Device) {
	log.Printf("dev: %s\t%f\t%s", dev.ID, dev.Reading, dev.Type)
	registry.Devices = append(registry.Devices, *dev)
}

// A simple private func to generate random data for type Device
func genDummyDevices(devices []*Device) []*Device {
	min_val := 17.0
	max_val := 24.0
	rand.Seed(time.Now().Unix())

	for i := range devices {
		// allocate memory of type Device
		tmp_dev := new(Device)
		// initialise ID. Note: %03d (zero-padding)
		tmp_dev.ID = fmt.Sprintf("dev-%03d", i)
		// Generate random float64 number between a range
		tmp_dev.Reading = min_val + rand.Float64()*(max_val-min_val)
		// Generate random device types between min/max enum range
		tmp := rand.Intn(int(NumOfDevTypes)-0) + 0
		// DeviceType is an enum with an assoc. func String()
		tmp_dev.Type = DeviceType(tmp).String()
		devices[i] = tmp_dev
	}
	return devices
}
