package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	door  Door
	temp  Temp
	video Video
)

func GetDoor(c *gin.Context) {
	door.Get()
	c.JSON(http.StatusOK, gin.H{
		"status": door.Status,
	})
}

func GetTemp(c *gin.Context) {
	temp.Get()
	c.JSON(http.StatusOK, gin.H{
		"inside":  temp.InsideSensor,
		"outside": temp.OutsideSensor,
	})
}

func GetVideo(c *gin.Context) {
	id := c.Param("id")
	video.Get()
	if id == "run" {
		c.JSON(http.StatusOK, gin.H{
			"location": video.Location,
			"url":      video.Url,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Must pass `run` not %s", id),
		})
	}
}
