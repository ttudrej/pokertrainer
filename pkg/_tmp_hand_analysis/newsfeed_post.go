package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttudrej/hand_analysis/platform/newsfeed"
)

type NewsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewsfeedPostRequest{}
		c.Bind(&requestBody) // Bind the request to NewsfeedPostRequest struct

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
