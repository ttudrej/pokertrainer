package main

import (
	"fmt"

	"github.com/ttudrej/pokertrainer/pkg/tableitems"
)

// ###########################################################

func main() {
	// var m tableitems.CardRankMapCreator

	// m = tableitems.CardRankMapStruct{}
	// describeM(m)
	// crm, err := m.Create()

	crm, err := tableitems.CardRankMapCreator.Create(tableitems.CardRankMapStruct{})

	fmt.Printf("crm: %v, err: %v\n", crm, err)

}

// ###########################################################

// func describeM(m tableitems.CardRankMapCreator) {
// 	fmt.Printf("(%v, %T)\n", m, m)
// }
