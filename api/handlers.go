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
		"status": door.Status})
}

func GetTemp(c *gin.Context) {
	temp.Get()
	c.JSON(http.StatusOK, gin.H{
		"inside":  temp.InsideSensor,
		"outside": temp.OutsideSensor})
}
