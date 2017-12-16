package api

import "github.com/smithgeoffrey/go-coop/config"

type Temp struct {
	InsideSensor  string `json:"insidesensor"`
	OutsideSensor string `json:"outsidesensor"`
}

func (t *Temp) Get() {
	// mock the sensors for now
	t.InsideSensor = config.MOCK_TEMP_INSIDE_SENSOR
	t.OutsideSensor = config.MOCK_TEMP_OUTSIDE_SENSOR
}
