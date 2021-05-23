package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttudrej/pokertrainer/pkg/calculating"
)

type PotOddsPostRequest struct {
	Pot     int `json:"pot"`
	CallAmt int `json:"call_amt"`
}

func PotOddsPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := PotOddsPostRequest{}
		c.Bind(&requestBody) // Bind the request to PotOddsPostRequest struct

		values := calculating.PotOddsPostRequest{
			Pot:     requestBody.Pot,
			CallAmt: requestBody.CallAmt,
		}

		result := values.PotOdds()

		c.JSON(http.StatusOK, result)
	}
}
