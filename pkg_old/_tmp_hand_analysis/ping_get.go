package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttudrej/hand_analysis/platform/newsfeed"
)

func PingGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
