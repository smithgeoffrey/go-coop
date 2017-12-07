// Geoff Smith, Dec 2017. Small REST api for ckicken coop automation.
//
// I'm trying my first web framework (gin) for this. 
// I started here: https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
//
// Frontend:
//     /        - home
//     /video   - ip cameras
//     /sensors - door and temp stats
//
// Api:
//     /api/v1/sensors         - list sensors
//     /api/v1/sensors/door    - list door sensors
//     /api/v1/sensors/door:id - get status of one of the door sensors
//     /api/v1/sensors/temp    - list temp sensors
//     /api/v1/sensors/temp:id - get status of one of the temp sensors
//     /api/v1/video           - list video sources
//     /api/v1/video:id        - get one of the video sources
//

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/smithgeoffrey/go-coop/api"
	"github.com/smithgeoffrey/go-coop/ui"
)

type DoorSensor struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Value    bool   `json:"value"`
}

type TempSensor struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Value    float32 `json:"inside"`
}

type Video struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Url  string `json:"url"`
}

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

