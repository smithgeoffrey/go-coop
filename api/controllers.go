package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Door struct {
	UpSensor  bool `json:"upsensor"`
	DownSensor  bool `json:"downsensor"`
	Status string `json:"status"`
}

type Temp struct {
	InsideSensor float32 `json:"inside"`
	OutideSensor float32 `json:"outside"`
}

type Video struct {
	Location string `json:"location"`
	Ip   string `json:"ip"`
	Url  string `json:"url"`
}

var door *Door

func GetDoor(c *gin.Context) {
	door.Get()
	c.JSON(http.StatusOK, gin.H{
		"message": door.Status,
	})
}

func GetTemp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("get temp sensor"),
	})
}

func GetVideo(c *gin.Context) {
	id := c.Param("id")
	if id != "run" {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Must pass param `run` not %s", id),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("get video %s", id),
		})
	}
}