package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttudrej/hand_analysis/platform/math"
	"github.com/ttudrej/pokertrainer/gameobjects"
)

// ############################################################
type HandOddsPostRequest struct {
	H1 gameobjects.CardString `json:"h1"`
	H2 gameobjects.CardString `json:"h2"`
	F1 gameobjects.CardString `json:"f1"`
	F2 gameobjects.CardString `json:"f2"`
	F3 gameobjects.CardString `json:"f3"`
	FT gameobjects.CardString `json:"ft"`
	FR gameobjects.CardString `json:"fr"`
}

// ############################################################
func HandOddsFPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := HandOddsPostRequest{}
		c.Bind(&requestBody) // Bind the request to NewsfeedPostRequest struct

		values := math.HandOddsPostRequest{
			H1: requestBody.H1,
			H2: requestBody.H2,
			F1: requestBody.F1,
			F2: requestBody.F2,
			F3: requestBody.F3,
			FT: requestBody.FT,
			FR: requestBody.FR,
		}

		result := values.HandOdds()

		c.JSON(http.StatusOK, result)
	}
}
