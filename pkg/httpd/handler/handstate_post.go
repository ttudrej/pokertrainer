package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// #################################################################
type HandStatePostRequest struct {
	Pot     int `json:"pot"`
	CallAmt int `json:"call_amt"`
}

// #################################################################

// HandStatePost receives the hand state payload
func HandStatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := HandStatePostRequest{}
		c.Bind(&requestBody) // Bind the request to HandStatePostRequest struct

		fmt.Println("Received POT     value : ", requestBody.Pot)
		fmt.Println("Received CallAmt value : ", requestBody.CallAmt)

		err := updateHandProgress(requestBody)
		if err != nil {
			fmt.Print("handle error condition : ", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"masgreceived": "true",
		})
	}
}

// #################################################################

// updateHandProgress updates the hand_state data structure to reflect current status.
func updateHandProgress(rb HandStatePostRequest) (err error) {

	return nil
}
