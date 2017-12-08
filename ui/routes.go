package ui

import (
        "net/http"

        "github.com/gin-gonic/gin"
)

// compare
func Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello %s", "chickens home page")
}

func Param(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello passed param in route: %s", name)
}
