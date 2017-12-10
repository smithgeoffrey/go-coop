package api

import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (d *Door) Get() {

	// mock the sensors for now
	d.DownSensor = true
	d.UpSensor = true

	if d.DownSensor && !(d.UpSensor) {
		d.Status = "down"
	} else if d.UpSensor && !(d.DownSensor) {
		d.Status = "up"
	} else {
		d.Status = "jammed"
	}
}
