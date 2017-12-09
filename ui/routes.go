package ui

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		// pass data to the template
		gin.H{
			"title": "Home Page",
		},
	)
}

