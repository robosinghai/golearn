package main

import "testing"

func TestSensorReading(t *testing.T) {
	want := 3.141592653589793
	if got := SensorReading(); got != want {
		t.Errorf("SensorReading() = %f, want %f", got, want)
	}
}
