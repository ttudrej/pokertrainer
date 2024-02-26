// package main

// import (
// 	"fmt"
// )

// type CardRank string
// type CardRankMap map[CardRank]int

// const (
// 	RK CardRank = "K"
// 	RA CardRank = "A"
// )

// // ###########################################################

// type CardRankMapCreator interface {
// 	// Create() (CardRankMap, error)
// 	Create() (CardRankMap, error)
// }

// type CardRankMapStruct struct {
// 	RankMap CardRankMap
// }

// func (c CardRankMapStruct) Create() (CardRankMap, error) {
// 	fmt.Println("c.RankMap: ", c.RankMap)

// 	c.RankMap = make(CardRankMap)
// 	c.RankMap[RA] = 14
// 	c.RankMap[RK] = 13
// 	fmt.Println("c.RankMap after assignment: ", c.RankMap)

// 	return c.RankMap, nil
// }

// // ###########################################################

// func main() {
// 	var m CardRankMapCreator

// 	m = CardRankMapStruct{}
// 	describeM(m)
// 	crm, err := m.Create()

// 	fmt.Printf("crm: %v, err: %v\n", crm, err)

// }

// // ###########################################################

// func describeM(m CardRankMapCreator) {
// 	fmt.Printf("(%v, %T)\n", m, m)
// }

// ###########################################################
// ###########################################################
// ###########################################################
// ###########################################################
// ###########################################################

// "github.com/ttudrej/pokertrainer/pkg/tableitems"

package main

import (
	"fmt"

	"github.com/ttudrej/pokertrainer/pkg/tableitems"
)

// type CardRank string
type CardRankMap map[tableitems.CardRank]int

const (
	RK tableitems.CardRank = "K"
	RA tableitems.CardRank = "A"
)

// ###########################################################

type CardRankMapCreator interface {
	// Create() (CardRankMap, error)
	Create() (CardRankMap, error)
}

type CardRankMapStruct struct {
	RankMap CardRankMap
}

func (c CardRankMapStruct) Create() (CardRankMap, error) {
	fmt.Println("c.RankMap: ", c.RankMap)

	c.RankMap = make(CardRankMap)
	c.RankMap[RA] = 14
	c.RankMap[RK] = 13
	fmt.Println("c.RankMap after assignment: ", c.RankMap)

	return c.RankMap, nil
}

// ###########################################################

func main() {
	// var m CardRankMapCreator
	var m CardRankMapCreator

	m = CardRankMapStruct{}
	describeM(m)
	crm, err := m.Create()

	fmt.Printf("crm: %v, err: %v\n", crm, err)

}

// ###########################################################

func describeM(m CardRankMapCreator) {
	fmt.Printf("(%v, %T)\n", m, m)
}
