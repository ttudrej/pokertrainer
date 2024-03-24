package main

import (
	"fmt"

	ha "github.com/ttudrej/pokertrainer/pkg/hand_analysis"
	mt "github.com/ttudrej/pokertrainer/pkg/manage_table"
)

// ###########################################################

func main() {

	// Display ranking of hand types
	ffl := ha.FchkrFL
	fmt.Println()
	fmt.Println("Display ranking of hand types")
	fmt.Println(ffl.SfInfo, ffl.X4Info, ffl.FhInfo, ffl.FlInfo, ffl.StInfo, ffl.X3Info, ffl.X22Info, ffl.X2Info, ffl.HcInfo)

	// Display card rank key:value assignemtn (map)
	crm, err := mt.CardRankMapCreator.Create(mt.CardRankMapStruct{})
	fmt.Println()
	fmt.Println("Display card rank key:value assignemtn (map)")
	fmt.Printf("crm: %v, err: %v\n", crm, err)

	// Display index assignemnts for all cards in a full deck.
	var cd mt.CardDeck
	cdPtr := &cd
	// Create a global card deck, since there will be only one, for any one hand played.
	cdPtr, _ = mt.CreateDeck()
	fmt.Println(cdPtr.TopCardIndex_oloptc, cdPtr.TopCardIndex_shuffledLoptc)
	// fmt.Println((*cdPtr.CdmPtr)[mt.CdmKey{Cr: mt.R2, Cs: mt.C}])
	fmt.Println()
	fmt.Println("Display index assignemnts for all cards in a full deck.")
	for _, cr := range mt.RankList {
		for _, cs := range mt.SuitList {
			fmt.Println((*cdPtr.CdmPtr)[mt.CdmKey{Cr: cr, Cs: cs}])
		}
	}

	// Display card ranks
	fmt.Println()
	fmt.Println("Display card ranks")
	for i, rank := range mt.RankList {
		fmt.Print(i+1, " r:", rank, ", ")
	}
	fmt.Println()

	// Display "full" list of card ranks
	fmt.Println()
	fmt.Println("Display full list of card ranks")
	for i, rank := range mt.RankListFull {
		fmt.Print(i+1, " r:", rank, ", ")
	}
	fmt.Println()

	// Display card suits
	fmt.Println()
	fmt.Println("Display card suits")
	for i, suit := range mt.SuitList {
		fmt.Print(i+1, " suit:", suit, ", ")
	}
	fmt.Println()

}

// ###########################################################

// func describeM(m mt.CardRankMapCreator) {
// 	fmt.Printf("(%v, %T)\n", m, m)
// }
