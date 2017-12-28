package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"),
		"src/github.com/smithgeoffrey/go-coop/ui/templates/*"))

	// from ~/routes.go
	initializeRoutes(router)

	router.Run(":8080")
}