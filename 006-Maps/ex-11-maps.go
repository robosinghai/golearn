package main

import (
	"golearn/devices"
	"log"
	"strconv"
)

func main() {
	dev := make([]*devices.Device, 5, 8) // Note: slice is 5
	// user-defined function generates dummy devices
	devices.GenDummyDevices(dev) // generate dummy devices

	reg := make(map[string]devices.Device, 8)
	// Registry is a map
	for i := range dev {
		reg[string("Device-00"+strconv.Itoa(i))] = *dev[i]
	}

	log.Printf("Number of devices in registry: %d",
		len(reg)) // prints 5
	log.Printf("--------------------------------------------------------")
	log.Printf("Device Name\tDevice ID\tReading\t\tDevice Type")
	log.Printf("--------------------------------------------------------")
	for devName, details := range reg {
		log.Printf("%s\t%s\t\t%f\t%s", devName,
			details.ID, details.Reading, details.Type)
	}
	log.Printf("--------------------------------------------------------")

}
