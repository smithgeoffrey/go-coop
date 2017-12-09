package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("ui/templates/*")

	// from ~/routes.go
	initializeRoutes(router)

	router.Run(":8080")
}
