package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttudrej/pokertrainer/pkg/calculating"
	"github.com/ttudrej/pokertrainer/pkg/tableitems"
)

// ############################################################
type HandOddsPostRequest struct {
	H1 tableitems.CardString `json:"h1"`
	H2 tableitems.CardString `json:"h2"`
	F1 tableitems.CardString `json:"f1"`
	F2 tableitems.CardString `json:"f2"`
	F3 tableitems.CardString `json:"f3"`
	FT tableitems.CardString `json:"ft"`
	FR tableitems.CardString `json:"fr"`
}

// ############################################################
func HandOddsFPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := HandOddsPostRequest{}
		c.Bind(&requestBody) // Bind the request to NewsfeedPostRequest struct

		values := calculating.HandOddsPostRequest{
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