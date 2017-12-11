package main

import (
	"github.com/gin-gonic/gin"

	"github.com/smithgeoffrey/go-coop/api"
	"github.com/smithgeoffrey/go-coop/ui"
)

func initializeRoutes(router *gin.Engine) {
	r1 := router.Group("/api/v1")
	{
		r1.GET("/sensor/door", api.GetDoor)
		r1.GET("/sensor/temp", api.GetTemp)
	}

	r2 := router.Group("")
	{
		r2.GET("/", ui.Home)
	}
}
