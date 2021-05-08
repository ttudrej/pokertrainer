package card

import (
	"log"
	"pokertrainer/debugging"
)

var (
	// Trace *log.Logger
	Info *log.Logger
	// Warning *log.Logger
	// Error   *log.Logger
)

// cardRank comment
type cardRank string

// cardSuit comment
type cardSuit string

// CardString comment
type CardString string

// card struct is meant for representing a physical card in a deck of cards,
// adn all it's properties, that will be used in a hand of a card game.
type Card struct {
	rank          cardRank
	suit          cardSuit
	community     bool // seen by all players
	dealtToPlayer bool // seen by whoever it was dealt to, only
	seenByHero    bool
	seenByVillain bool
	sequence      int // assigned by our convention, just so we have another way to reference cards
}

type cardRankList [13]cardRank
type cardRankListFull [14]cardRank
type cardSuitList [4]cardSuit

type cardRankMap map[cardRank]int

const (
	x cardSuit = "Suit X"
	s cardSuit = "s"
	c cardSuit = "c"
	h cardSuit = "h"
	d cardSuit = "d"

	rX cardRank = "Rank X"
	r2 cardRank = "2"
	r3 cardRank = "3"
	r4 cardRank = "4"
	r5 cardRank = "5"
	r6 cardRank = "6"
	r7 cardRank = "7"
	r8 cardRank = "8"
	r9 cardRank = "9"
	rT cardRank = "T"
	rJ cardRank = "J"
	rQ cardRank = "Q"
	rK cardRank = "K"
	rA cardRank = "A"

	s2 CardString = "2s"
	s3 CardString = "3s"
	s4 CardString = "4s"
	s5 CardString = "5s"
	s6 CardString = "6s"
	s7 CardString = "7s"
	s8 CardString = "8s"
	s9 CardString = "9s"
	sT CardString = "Ts"
	sJ CardString = "Js"
	sQ CardString = "Qs"
	sK CardString = "Ks"
	sA CardString = "As"

	c2 CardString = "2c"
	c3 CardString = "3c"
	c4 CardString = "4c"
	c5 CardString = "5c"
	c6 CardString = "6c"
	c7 CardString = "7c"
	c8 CardString = "8c"
	c9 CardString = "9c"
	cT CardString = "Tc"
	cJ CardString = "Jc"
	cQ CardString = "Qc"
	cK CardString = "Kc"
	cA CardString = "Ac"

	h2 CardString = "2h"
	h3 CardString = "3h"
	h4 CardString = "4h"
	h5 CardString = "5h"
	h6 CardString = "6h"
	h7 CardString = "7h"
	h8 CardString = "8h"
	h9 CardString = "9h"
	hT CardString = "Th"
	hJ CardString = "Jh"
	hQ CardString = "Qh"
	hK CardString = "Kh"
	hA CardString = "Ah"

	d2 CardString = "2d"
	d3 CardString = "3d"
	d4 CardString = "4d"
	d5 CardString = "5d"
	d6 CardString = "6d"
	d7 CardString = "7d"
	d8 CardString = "8d"
	d9 CardString = "9d"
	dT CardString = "Td"
	dJ CardString = "Jd"
	dQ CardString = "Qd"
	dK CardString = "Kd"
	dA CardString = "Ad"
)

var rankList = cardRankList{rA, rK, rQ, rJ, rT, r9, r8, r7, r6, r5, r4, r3, r2}
var rankListFull = cardRankListFull{rA, rK, rQ, rJ, rT, r9, r8, r7, r6, r5, r4, r3, r2, rA}

var suitList = cardSuitList{s, c, h, d}

// Used to indicate the absence of a card, instead of using an undefined or null pointer
// By convention, we number cards in the deck from 1 to 52
// So, there is no card 0, hence our noCard has the "sequnce" nuber assibned to 0.
var noCardPtr = &Card{rank: rX, suit: x, community: false, dealtToPlayer: false, seenByHero: false, seenByVillain: false, sequence: 0}

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
func createCardRankMap() (crm cardRankMap) {

	Info.Println(debugging.ThisFunc())
	// Info.Println("### Starting createCardRankMap ###")
	crm = make(cardRankMap)

	crm[rA] = 14
	crm[rK] = 13
	crm[rQ] = 12
	crm[rJ] = 11
	crm[rT] = 10
	crm[r9] = 9
	crm[r8] = 8
	crm[r7] = 7
	crm[r6] = 6
	crm[r5] = 5
	crm[r4] = 4
	crm[r3] = 3
	crm[r2] = 2

	return crm
}
