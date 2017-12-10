package api

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (d *Door) Get() {
	if d.DownSensor && !(d.UpSensor) {
		d.Status = "down"
	} else if d.UpSensor && !(d.DownSensor) {
		d.Status = "up"
	} else {
		d.Status = "jammed"
	}
}
