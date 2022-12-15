package main

import "log"

func NewFuncExample() {
	sensor := new(Device) // or data := &Device() both are valid!
	sensor.ID = "temp-001"
	sensor.Reading = 17.0

	// or
	sensor2 := &Device{ID: "hum-001", Reading: 1.4}
	log.Printf("Device1 ID: %s\t Device2 ID: %s\n", sensor.ID, sensor2.ID)
}
