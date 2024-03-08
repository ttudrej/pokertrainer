// *build unit
/*
To run ONLY the tests in this test file, run:
>go test -tags=unit -v
*/

package math

import (
	"testing"

	"github.com/ttudrej/pokertrainer/gameobjects"
)

// ############################################################
type CalcPOTest struct {
	popr     PotOddsPostRequest
	expected float64
}

var CalcPOTests = []CalcPOTest{
	{popr: PotOddsPostRequest{3, 1}, expected: 25.0},
	{popr: PotOddsPostRequest{100, 100}, expected: 50.0},
	{popr: PotOddsPostRequest{100, 200}, expected: 66.67},
	{popr: PotOddsPostRequest{1, 1000}, expected: 99.90},
	{popr: PotOddsPostRequest{1000, 1}, expected: 0.10},
	{popr: PotOddsPostRequest{0, 1}, expected: 0.0},
	{popr: PotOddsPostRequest{1, 0}, expected: 0.0},
	{popr: PotOddsPostRequest{0, 0}, expected: 0.0},
}

func TestPotOddsCalculator(t *testing.T) {
	for _, test := range CalcPOTests {

		result := test.popr.PotOdds()

		// t.Log("Result - Expected: ", result, " - ", test.expected)
		if result != test.expected {
			// t.Fatal("Result not equal to Expected: ", result, " != ", test.expected)
			t.Error("Result not equal to Expected: ", result, " != ", test.expected)
		}
	}
}

// ############################################################
type CalcHOTest struct {
	hopr     HandOddsPostRequest
	expected float64
}

var CalcHOTests = []CalcHOTest{
	{hopr: HandOddsPostRequest{
		gameobjects.C2, gameobjects.S2,
		gameobjects.NoCard, gameobjects.NoCard, gameobjects.NoCard,
		gameobjects.NoCard,
		gameobjects.NoCard}, expected: 4.0},
	{hopr: HandOddsPostRequest{
		gameobjects.C2, gameobjects.S2,
		gameobjects.SA, gameobjects.CA, gameobjects.DA,
		gameobjects.NoCard,
		gameobjects.NoCard}, expected: 3.0},
	{hopr: HandOddsPostRequest{
		gameobjects.C2, gameobjects.S2,
		gameobjects.SA, gameobjects.CA, gameobjects.DA,
		gameobjects.SK,
		gameobjects.NoCard}, expected: 2.0},
	{hopr: HandOddsPostRequest{
		gameobjects.C2, gameobjects.S2,
		gameobjects.SA, gameobjects.CA, gameobjects.DA,
		gameobjects.SK,
		gameobjects.SQ}, expected: 1.0},
}

func TestHandOddsCalculator(t *testing.T) {
	for _, test := range CalcHOTests {

		result := test.hopr.HandOdds()

		// t.Log("Result - Expected: ", result, " - ", test.expected)
		if result != test.expected {
			// t.Fatal("Result not equal to Expected: ", result, " != ", test.expected)
			t.Error("Result not equal to Expected: ", result, " != ", test.expected)
		}
	}
}

// ############################################################

// ############################################################

// ############################################################
