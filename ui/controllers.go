package ui

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/smithgeoffrey/go-coop/config"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		// pass data to the template
		gin.H{
			"title": "Home Page",
			"video_url": config.VIDEO_URL,
		},
	)
}
