package api

import (
        "net/http"

        "github.com/gin-gonic/gin"
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

func ListSensors(c *gin.Context) {
        c.String(http.StatusOK, "Hello %s", "list sensors")
}

func ListDoorSensors(c *gin.Context) {
        c.String(http.StatusOK, "Hello %s", "list Door Sensors")
}

func GetDoorSensor(c *gin.Context) {
        id := c.Param("id")
        c.String(http.StatusOK, "Hello %s: %s", "Get Door Sensors", id)
}

func ListTempSensors(c *gin.Context) {
        c.String(http.StatusOK, "Hello %s", "list Temp Sensors")
}

func GetTempSensor(c *gin.Context) {
        id := c.Param("id")
        c.String(http.StatusOK, "Hello %s: %s", "Get Temp Sensors", id)
}

func ListVideo(c *gin.Context) {
        c.String(http.StatusOK, "Hello %s", "list Video")
}

func GetVideo(c *gin.Context) {
        id := c.Param("id")
        c.String(http.StatusOK, "Hello %s: %s", "Get Video", id)
}
