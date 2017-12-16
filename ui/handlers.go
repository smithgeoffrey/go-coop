package ui

import (
    "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smithgeoffrey/go-coop/config"
	"github.com/smithgeoffrey/go-coop/api"
)

func Home(c *gin.Context) {

	door := api.Door{}
	temperature := api.Temp{}

	response, err := http.Get(config.TEMP_URL)
	if err != nil {
		panic(err.Error())
	}
	json.NewDecoder(response.Body).Decode(&temperature)

	response, err = http.Get(config.DOOR_URL)
	if err != nil {
		panic(err.Error())
	}
	json.NewDecoder(response.Body).Decode(&door)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
		"videoUrl": config.VIDEO_URL,
		"doorStatus": door.Status,
		"tempOutside": temperature.OutsideSensor,
		"tempInside": temperature.InsideSensor })

}
