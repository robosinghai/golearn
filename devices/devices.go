package devices

import (
	"fmt"
	"math/rand"
	"time"
)

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

// Demo - Expand a slice of an array and add append to it
// Note - The slice is expanded in the MakeArray func ^ ^
func (registry *Registry) AddDeviceToRegistry(dev *Device) {

	registry.Devices = append(registry.Devices, *dev)
}

// A simple private func to generate random data for type Device
func GenDummyDevices(devices []*Device) []*Device {
	min_val := 17.0
	max_val := 24.0
	rand.Seed(time.Now().Unix())

	for i := range devices {
		// allocate memory of type Device
		tmp_dev := new(Device)
		// initialise ID. Note: %03d (zero-padding)
		tmp_dev.ID = fmt.Sprintf("d-%03d", i)
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

func AddBias(d *Device) {
	bias := 10.5
	d.Reading += bias
}

func (d *Device) AddBias() {
	bias := 10.5
	d.Reading += bias
}
