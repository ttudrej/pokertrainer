package main

import (
	"fmt"

	"github.com/ttudrej/pokertrainer/pkg/manage_table"
)

// ###########################################################

func main() {
	// var m manage_table.CardRankMapCreator

	// m = manage_table.CardRankMapStruct{}
	// describeM(m)
	// crm, err := m.Create()

	crm, err := manage_table.CardRankMapCreator.Create(manage_table.CardRankMapStruct{})

	fmt.Printf("crm: %v, err: %v\n", crm, err)

}

// ###########################################################

// func describeM(m manage_table.CardRankMapCreator) {
// 	fmt.Printf("(%v, %T)\n", m, m)
// }
