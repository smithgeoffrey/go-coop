package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	door  Door
	temp  Temp
)

func GetDoor(c *gin.Context) {
	door.Get()
	c.JSON(http.StatusOK, gin.H{
		"UpSensor": door.UpSensor,
		"DownSensor": door.DownSensor,
		"Status": door.Status})
}

func GetTemp(c *gin.Context) {
	temp.Get()
	c.JSON(http.StatusOK, gin.H{
		"InsideSensor":  temp.InsideSensor,
		"OutsideSensor": temp.OutsideSensor})
}
