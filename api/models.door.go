package api

import "github.com/smithgeoffrey/go-coop/config"

type Door struct {
	UpSensor   bool   `json:"upsensor"`
	DownSensor bool   `json:"downsensor"`
	Status     string `json:"status"`
}

func (d *Door) Get() {
	// mock the sensors for now
	d.DownSensor = config.MOCK_DOOR_DOWN_SENSOR
	d.UpSensor = config.MOCK_DOOR_UP_SENSOR

	if d.DownSensor && !(d.UpSensor) {
		d.Status = "Down"
	} else if d.UpSensor && !(d.DownSensor) {
		d.Status = "Up"
	} else {
		d.Status = "Jam"
	}
}
