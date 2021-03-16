package main

// cardRank comment
type cardRank string

// cardSuit comment
type cardSuit string

// cardString comment
type cardString string

// card struct is meant for representing a physical card in a deck of cards,
// adn all it's properties, that will be used in a hand of a card game.
type card struct {
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

	s2 cardString = "2s"
	s3 cardString = "3s"
	s4 cardString = "4s"
	s5 cardString = "5s"
	s6 cardString = "6s"
	s7 cardString = "7s"
	s8 cardString = "8s"
	s9 cardString = "9s"
	sT cardString = "Ts"
	sJ cardString = "Js"
	sQ cardString = "Qs"
	sK cardString = "Ks"
	sA cardString = "As"

	c2 cardString = "2c"
	c3 cardString = "3c"
	c4 cardString = "4c"
	c5 cardString = "5c"
	c6 cardString = "6c"
	c7 cardString = "7c"
	c8 cardString = "8c"
	c9 cardString = "9c"
	cT cardString = "Tc"
	cJ cardString = "Jc"
	cQ cardString = "Qc"
	cK cardString = "Kc"
	cA cardString = "Ac"

	h2 cardString = "2h"
	h3 cardString = "3h"
	h4 cardString = "4h"
	h5 cardString = "5h"
	h6 cardString = "6h"
	h7 cardString = "7h"
	h8 cardString = "8h"
	h9 cardString = "9h"
	hT cardString = "Th"
	hJ cardString = "Jh"
	hQ cardString = "Qh"
	hK cardString = "Kh"
	hA cardString = "Ah"

	d2 cardString = "2d"
	d3 cardString = "3d"
	d4 cardString = "4d"
	d5 cardString = "5d"
	d6 cardString = "6d"
	d7 cardString = "7d"
	d8 cardString = "8d"
	d9 cardString = "9d"
	dT cardString = "Td"
	dJ cardString = "Jd"
	dQ cardString = "Qd"
	dK cardString = "Kd"
	dA cardString = "Ad"
)

var rankList = cardRankList{rA, rK, rQ, rJ, rT, r9, r8, r7, r6, r5, r4, r3, r2}
var rankListFull = cardRankListFull{rA, rK, rQ, rJ, rT, r9, r8, r7, r6, r5, r4, r3, r2, rA}

var suitList = cardSuitList{s, c, h, d}

// Used to indicate the absence of a card, instead of using an undefined or null pointer
// By convention, we number cards in the deck from 1 to 52
// So, there is no card 0, hence our noCard has the "sequnce" nuber assibned to 0.
var noCardPtr = &card{rank: rX, suit: x, community: false, dealtToPlayer: false, seenByHero: false, seenByVillain: false, sequence: 0}

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

	Info.Println(thisFunc())
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
