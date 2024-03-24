package main

import (
	"fmt"

	"github.com/ttudrej/pokertrainer/pkg/hand_analysis"
	"github.com/ttudrej/pokertrainer/pkg/manage_table"
)

// ###########################################################

func main() {

	ffl, _ := hand_analysis.CreateFiveCardHandKindRankings()
	fmt.Println(ffl.SfInfo, ffl.X4Info, ffl.FhInfo, ffl.FlInfo, ffl.StInfo, ffl.X3Info, ffl.X22Info, ffl.X2Info, ffl.HcInfo)
	// fmt.Println(ffl.SfInfo)

	// var m manage_table.CardRankMapCreator

	// m = manage_table.CardRankMapStruct{}
	// describeM(m)
	// crm, err := m.Create()

	crm, err := manage_table.CardRankMapCreator.Create(manage_table.CardRankMapStruct{})
	fmt.Printf("crm: %v, err: %v\n", crm, err)

	// var cdm manage_table.CardDeckMap
	var cd manage_table.CardDeck
	cdPtr := &cd

	// Create a global card deck, since there will be only one, for any one hand played.
	// cdm, orderedList, orderedListFull = manage_table.CreateDeck()
	cdPtr, _ = manage_table.CreateDeck()
	// fmt.Println(cdm[cdmKey{r2, s}])
	fmt.Println(cdPtr.TopCardIndex_oloptc, cdPtr.TopCardIndex_shuffledLoptc)

}

// ###########################################################

// func describeM(m manage_table.CardRankMapCreator) {
// 	fmt.Printf("(%v, %T)\n", m, m)
// }
