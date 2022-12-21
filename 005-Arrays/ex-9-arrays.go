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
	// This function generates dummy devices
	genDummyDevices(dev) // generate dummy devices

	// Add to registry
	reg := Registry{}

	for i := range dev {
		reg.AddDeviceToRegistry(dev[i])
	}
	log.Printf("Number of devices in registyr: %d", len(reg.Devices))

}
