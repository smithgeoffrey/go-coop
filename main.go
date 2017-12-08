package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/smithgeoffrey/go-coop/api"
	"github.com/smithgeoffrey/go-coop/ui"
)

func main() {
	router := gin.Default()

	r1 := router.Group("/api/v1")
	{
		r1.GET("/sensor", api.ListSensors)
		r1.GET("/sensor/door", api.ListDoorSensors)
		r1.GET("/sensor/door/:id", api.GetDoorSensor)
		r1.GET("/sensor/temp", api.ListTempSensors)
		r1.GET("/sensor/temp/:id", api.GetTempSensor)
		r1.GET("/video", api.ListVideo)
		r1.GET("/video/:id", api.GetVideo)
	}

	r2 := router.Group("")
	{
		r2.GET("/", ui.Home)
		r2.GET("/param/:name", ui.Param)
	}

	fmt.Println("Starting :8080 listener")
	router.Run()
}

