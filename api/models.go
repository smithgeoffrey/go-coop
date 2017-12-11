package api

import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (d *Door) Get() {

	// mock the sensors for now
	d.DownSensor = true
	d.UpSensor = false

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
	t.InsideSensor = 35.1
	t.OutsideSensor = 28.4
}

func (v *Video) Get() {

	// mock the sensors for now
	v.Location = "run"
	v.Url = "http://172.16.1.128/video.mjpg"
}
