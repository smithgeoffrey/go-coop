package api

import "github.com/smithgeoffrey/go-coop/config"

type Temp struct {
	InsideSensor  float32 `json:"inside"`
	OutsideSensor float32 `json:"outside"`
}

func (t *Temp) Get() {
	// mock the sensors for now
	t.InsideSensor = config.MOCK_TEMP_INSIDE_SENSOR
	t.OutsideSensor = config.MOCK_TEMP_OUTSIDE_SENSOR
}
