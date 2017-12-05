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
	"net/http"
)

type DoorSensor struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Value    bool   `json:"value"`
}

type TempSensor struct {
	Name     string  `json:"name"`
	Location string  `hson`
	Value    float32 `json:"inside"`
}

type Video struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Url  string `json:"url"`
}

func main() {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		// api
		api.GET("/sensor", apiListSensors)
		api.GET("/sensor/door", apiListDoorSensors)
		api.GET("/sensor/door/:id", apiGetDoorSensor)
		api.GET("/sensor/temp", apiListTempSensors)
		api.GET("/sensor/temp/:id", apiGetTempSensor)
		api.GET("/video", apiListVideo)
		api.GET("/video/:id", apiGetVideo)
	}

	ui := router.Group("")
	{
		ui.GET("/", uiHome)
		ui.GET("/param/:name", uiParam)
	}

	fmt.Println("Starting :8080 listener")
	router.Run()
}

// compare
func uiHome(c *gin.Context) {
	// name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", "chickens home page")
}

func uiParam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello passed param in route: %s", name)
}

// Todo: fill in these api funcs
func apiListSensors(c *gin.Context)     {}
func apiListDoorSensors(c *gin.Context) {}
func apiGetDoorSensor(c *gin.Context)   {}
func apiListTempSensors(c *gin.Context) {}
func apiGetTempSensor(c *gin.Context)   {}
func apiListVideo(c *gin.Context)       {}
func apiGetVideo(c *gin.Context)        {}
