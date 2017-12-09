package api

import (
        "fmt"
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
        c.JSON(http.StatusOK, gin.H{
                "message": "list sensors",
        })
}

func ListDoorSensors(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
                "message": "list door sensors",
        })
}

func GetDoorSensor(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
                "message": fmt.Sprintf("get door sensor %s", id),
        })

}

func ListTempSensors(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
                "message": "list temp sensors",
        })
}

func GetTempSensor(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
                "message": fmt.Sprintf("get temp sensor %s", id),
        })
}

func ListVideo(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
                "message": "list video",
        })
}

func GetVideo(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
                "message": fmt.Sprintf("get video %s", id),
        })
}
