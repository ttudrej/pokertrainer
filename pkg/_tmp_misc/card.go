// Package gameobjects comment string
package gameobjects

import (
	"log"

	"github.com/ttudrej/pokertrainer/pkg/debugging"
)

var (
	// Trace *log.Logger
	Info *log.Logger
	// Warning *log.Logger
	// Error   *log.Logger
)

// CardRank comment
type CardRank string

// CardSuit comment
type CardSuit string

// CardString comment
type CardString string

// card struct is meant for representing a physical card in a deck of cards,
// and all it's properties, that will be used in a hand of a card game.
type Card struct {
	Rank          CardRank
	Suit          CardSuit
	Community     bool // seen by all players
	DealtToPlayer bool // seen by whoever it was dealt to, only
	SeenByHero    bool
	SeenByVillain bool
	Sequence      int // assigned by our convention, just so we have another way to reference cards
}

type CardRankList [13]CardRank
type CardRankListFull [14]CardRank
type CardSuitList [4]CardSuit

type CardRankMap map[CardRank]int

const (
	X CardSuit = "Suit X"
	S CardSuit = "s"
	C CardSuit = "c"
	H CardSuit = "h"
	D CardSuit = "d"

	RX CardRank = "Rank X"
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

var RankList = CardRankList{RA, RK, RQ, RJ, RT, R9, R8, R7, R6, R5, R4, R3, R2}
var RankListFull = CardRankListFull{RA, RK, RQ, RJ, RT, R9, R8, R7, R6, R5, R4, R3, R2, RA}

var SuitList = CardSuitList{S, C, H, D}

// Used to indicate the absence of a card, instead of using an undefined or null pointer
// By convention, we number cards in the deck from 1 to 52
// So, there is no card 0, hence our noCard has the "sequnce" nuber assigned to 0.
var NoCardPtr = &Card{Rank: RX, Suit: X, Community: false, DealtToPlayer: false, SeenByHero: false, SeenByVillain: false, Sequence: 0}

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

// createCardRankMap provedes a mapping of card rank to an index.
// Useful in carda and hand ranking comparision.
func CreateCardRankMap() (crm CardRankMap) {

	Info.Println(debugging.ThisFunc())
	// Info.Println("### Starting createCardRankMap ###")
	crm = make(CardRankMap)

	crm[RA] = 14
	crm[RK] = 13
	crm[RQ] = 12
	crm[RJ] = 11
	crm[RT] = 10
	crm[R9] = 9
	crm[R8] = 8
	crm[R7] = 7
	crm[R6] = 6
	crm[R5] = 5
	crm[R4] = 4
	crm[R3] = 3
	crm[R2] = 2

	return crm
}
