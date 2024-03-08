// Package manage_table comment string
package manage_table

import (
	"log"
)

var (
	// Trace *log.Logger
	Info *log.Logger
	// Warning *log.Logger
	// Error   *log.Logger
)

// CardRank comment
type CardRank string

// CardRankMap comment
type CardRankMap map[CardRank]int

// CardSuit comment
type CardSuit string

// CardString comment
type CardString string

// card struct is meant for representing a physical card in a deck of cards,
// adn all it's properties, that will be used in a hand of a card game.
type Card struct {
	Rank          CardRank
	Suit          CardSuit
	Community     bool // seen by all players
	DealtToPlayer bool // seen by whoever it was dealt to, only
	SeenByHero    bool
	SeenByVillain bool
	Sequence      int // assigned by our convention, just so we have another way to reference cards
}

// 3 things are needed to use interfaces instead of functins directly:
//
// 1) One or more structs, sa, sb, sc, ... (type mystructname struct {})
// => type CardRankMap struct
//
// 2) One or more methods (functions with a rceiver argument),  and same functin name,
// 		where the reciver references one of the structs, sa, sb, sc, ....
// 		(func (<receiver-structref>) funcname(<inputs>) (<outputs>))
// => func (c Card) crate() (CardRankMap, error) {}
//
// 3) An Interface definition, which pools/ties the
//		functions (method with reciver argument stripped)
//		with the same name but different func signatures.
// => type cardRankMapCreator interface
//
// 4) Something that uses the interface.
// 	Note, the functin call via the interface takes the associated struct as the argument,
// 		not and not input args directly. The inputs are fed via the struct definition.
//

type CardRankList [13]CardRank
type CardRankListFull [14]CardRank
type CardSuitList [4]CardSuit

const (
	X CardSuit = "Suit X" // ?? How is this used
	S CardSuit = "s"      // Spades
	C CardSuit = "c"      // Clubs
	H CardSuit = "h"      // Hearts
	D CardSuit = "d"      // Diamonds

	RX CardRank = "Rank X" // ?? what is this needed for
	R2 CardRank = "2"
	R3 CardRank = "3"
	R4 CardRank = "4"
	R5 CardRank = "5"
	R6 CardRank = "6"
	R7 CardRank = "7"
	R8 CardRank = "8"
	R9 CardRank = "9"
	RT CardRank = "T"
	RJ CardRank = "J"
	RQ CardRank = "Q"
	RK CardRank = "K"
	RA CardRank = "A"

	NoCard CardString = ""

	S2 CardString = "2s"
	S3 CardString = "3s"
	S4 CardString = "4s"
	S5 CardString = "5s"
	S6 CardString = "6s"
	S7 CardString = "7s"
	S8 CardString = "8s"
	S9 CardString = "9s"
	ST CardString = "Ts"
	SJ CardString = "Js"
	SQ CardString = "Qs"
	SK CardString = "Ks"
	SA CardString = "As"

	C2 CardString = "2c"
	C3 CardString = "3c"
	C4 CardString = "4c"
	C5 CardString = "5c"
	C6 CardString = "6c"
	C7 CardString = "7c"
	C8 CardString = "8c"
	C9 CardString = "9c"
	CT CardString = "Tc"
	CJ CardString = "Jc"
	CQ CardString = "Qc"
	CK CardString = "Kc"
	CA CardString = "Ac"

	H2 CardString = "2h"
	H3 CardString = "3h"
	H4 CardString = "4h"
	H5 CardString = "5h"
	H6 CardString = "6h"
	H7 CardString = "7h"
	H8 CardString = "8h"
	H9 CardString = "9h"
	HT CardString = "Th"
	HJ CardString = "Jh"
	HQ CardString = "Qh"
	HK CardString = "Kh"
	HA CardString = "Ah"

	D2 CardString = "2d"
	D3 CardString = "3d"
	D4 CardString = "4d"
	D5 CardString = "5d"
	D6 CardString = "6d"
	D7 CardString = "7d"
	D8 CardString = "8d"
	D9 CardString = "9d"
	DT CardString = "Td"
	DJ CardString = "Jd"
	DQ CardString = "Qd"
	DK CardString = "Kd"
	DA CardString = "Ad"
)

// var RankList = CardRankList{RA, RK, RQ, RJ, RT, R9, R8, R7, R6, R5, R4, R3, R2}
// var RankListFull = CardRankListFull{RA, RK, RQ, RJ, RT, R9, R8, R7, R6, R5, R4, R3, R2, RA}

// var SuitList = CardSuitList{S, C, H, D}

// // NoCardPtr:
// // Used to indicate the absence of a card, instead of using an undefined or null pointer
// // By convention, we number cards in the deck from 1 to 52
// // So, there is no card 0, hence our noCard has the "sequnce" nuber assigned to 0.
// var NoCardPtr = &Card{Rank: RX, Suit: X, Community: false, DealtToPlayer: false, SeenByHero: false, SeenByVillain: false, Sequence: 0}

var ( // "factored" into a block
	RankList     = CardRankList{RA, RK, RQ, RJ, RT, R9, R8, R7, R6, R5, R4, R3, R2}
	RankListFull = CardRankListFull{RA, RK, RQ, RJ, RT, R9, R8, R7, R6, R5, R4, R3, R2, RA}
	SuitList     = CardSuitList{S, C, H, D}
	NoCardPtr    = &Card{Rank: RX, Suit: X, Community: false, DealtToPlayer: false, SeenByHero: false, SeenByVillain: false, Sequence: 0}
)

/*
#########################################################################
#########################################################################
#########################################################################

######## ##     ## ##    ##  ######   ######
##       ##     ## ###   ## ##    ## ##    ##
##       ##     ## ####  ## ##       ##
######   ##     ## ## ## ## ##        ######
##       ##     ## ##  #### ##             ##
##       ##     ## ##   ### ##    ## ##    ##
##        #######  ##    ##  ######   ######

#########################################################################
#########################################################################
#########################################################################
*/

type CardRankMapCreator interface {
	// Create() (CardRankMap, error)
	Create() (CardRankMap, error)
}

type CardRankMapStruct struct {
	RankMap CardRankMap
}

func (c CardRankMapStruct) Create() (CardRankMap, error) {
	// fmt.Println("c.RankMap: ", c.RankMap)

	// Info.Println(debugging.ThisFunc()) // ! throws a panic, currenlty
	// 	panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x10027a460]

	c.RankMap = make(CardRankMap)
	c.RankMap[RA] = 14
	c.RankMap[RK] = 13
	c.RankMap[RQ] = 12
	c.RankMap[RJ] = 11
	c.RankMap[RT] = 10
	c.RankMap[R9] = 9
	c.RankMap[R8] = 8
	c.RankMap[R7] = 7
	c.RankMap[R6] = 6
	c.RankMap[R5] = 5
	c.RankMap[R4] = 4
	c.RankMap[R3] = 3
	c.RankMap[R2] = 2
	// fmt.Println("c.RankMap after assignment: ", c.RankMap)

	return c.RankMap, nil
}
