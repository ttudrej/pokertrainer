package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandStatePostRequest struct {
	Pot     int `json:"pot"`
	CallAmt int `json:"call_amt"`
}

// HandStatePost receives the hand state payload
func HandStatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := HandStatePostRequest{}
		c.Bind(&requestBody) // Bind the request to HandStatePostRequest struct

		// values := math.HandStatePostRequest{
		// 	Pot:     requestBody.Pot,
		// 	CallAmt: requestBody.CallAmt,
		// }

		// result := values.PotOdds()
		result := 1.0

		c.JSON(http.StatusOK, result)
	}
}
