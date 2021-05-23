package calculating

// ############################################################
import (
	"fmt"
	"math"

	"github.com/ttudrej/pokertrainer/pkg/tableitems"
)

// ############################################################
type PotOddsPostRequest struct {
	Pot     int
	CallAmt int
}

// ############################################################
type PotOddsCalculator interface {
	PotOdds(popr PotOddsPostRequest) float64
}

// ############################################################
func (popr PotOddsPostRequest) PotOdds() float64 {

	if popr.Pot == 0 {
		return 0
	}
	if popr.CallAmt == 0 {
		return 0
	}

	po := float64(popr.CallAmt) * 100 / (float64(popr.Pot) + float64(popr.CallAmt))
	// fmt.Println(math.Floor(x*100)/100) // 12.34 (round down)
	// fmt.Println(math.Round(x*100)/100) // 12.35 (round to nearest)
	// fmt.Println(math.Ceil(x*100)/100)  // 12.35 (round up)
	return math.Round(po*100) / 100
}

// ############################################################
type HandOddsPostRequest struct {
	H1 tableitems.CardString
	H2 tableitems.CardString
	F1 tableitems.CardString
	F2 tableitems.CardString
	F3 tableitems.CardString
	FT tableitems.CardString
	FR tableitems.CardString
}

// ############################################################
type HandOddsCalculator interface {
	HandOdds(hopr HandOddsPostRequest) float64
}

// ############################################################
// HandOdds calculates the odds of a hand to win, for either the Hero or Villain, based on variables provided.
func (hopr HandOddsPostRequest) HandOdds() float64 {

	ho := 0.0

	// Figure out which scenario we're working with
	switch {
	case hopr.FR != "":
		// fmt.Println("we're on the River")
		ho = 1
	case hopr.FT != "":
		// fmt.Println("we're on the Turn")
		ho = 2
	case hopr.F3 != "":
		// fmt.Println("we're on the Flop")
		ho = 3
	case hopr.H2 != "":
		// fmt.Println("we're on Pre-Flop")
		ho = 4
	default:
		fmt.Println("!!! You should not be getting here, so figure out how we did!")
		ho = 5
	}

	return math.Round(ho*100) / 100
}