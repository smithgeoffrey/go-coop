package ui

import (
        "net/http"

        "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	// HTML method of Context to render a template
	c.HTML(http.StatusOK, "index.html",
		// pass data to the template
		gin.H{
			"title": "Home Page",
		},
	)
}

func Param(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello passed param in route: %s", name)
}
