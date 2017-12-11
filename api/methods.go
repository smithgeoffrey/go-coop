package api

import (
	"github.com/smithgeoffrey/go-coop/config"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (d *Door) Get() {
	// mock the sensors for now
	d.DownSensor = config.MOCK_DOOR_DOWN_SENSOR
	d.UpSensor = config.MOCK_DOOR_UP_SENSOR

	if d.DownSensor && !(d.UpSensor) {
		d.Status = "down"
	} else if d.UpSensor && !(d.DownSensor) {
		d.Status = "up"
	} else {
		d.Status = "error"
	}
}

func (t *Temp) Get() {
	// mock the sensors for now
	t.InsideSensor = config.MOCK_TEMP_INSIDE_SENSOR
	t.OutsideSensor = config.MOCK_TEMP_OUTSIDE_SENSOR
}

func (v *Video) Get() {
	v.Location = "run"
	v.Url = config.VIDEO_URL
}
