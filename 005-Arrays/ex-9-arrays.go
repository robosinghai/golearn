package main

import "log"

type Device struct {
	ID      string
	Reading float64
	Type    string
}

type Registry struct {
	Devices []Device
}

type DeviceType int

const (
	WLAN DeviceType = iota
	TX
	Actuator
	Sensor
	NumOfDevTypes
)

func (dType DeviceType) String() string {
	return []string{"WLAN", "TX", "Actuator", "Sensor"}[dType]
}

func main() {
	dev := MakeArray(5) // make array of lenght 5
	// user-defined function generates dummy devices
	genDummyDevices(dev) // generate dummy devices

	// Add to registry
	reg := Registry{} // empty structure (method 1) or
	// reg := new(Registry) // method 2

	// Append each device to registry
	for i := range dev {
		reg.AddDeviceToRegistry(dev[i])
	}
	log.Printf("Number of devices in registry: %d",
		len(reg.Devices)) // prints 5
	for i := range reg.Devices {
		log.Printf("dev: %s\t%f\t%s", reg.Devices[i].ID,
			reg.Devices[i].Reading, reg.Devices[i].Type)
	}
}
