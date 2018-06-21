package main

import "fmt"

type cardRank string
type cardSuit string

type cardRankList [13]cardRank
type cardSuitList [4]cardSuit

type card struct {
	rank     cardRank
	suit     cardSuit
	seen     bool
	sequence int
}

var rankList = cardRankList{rA, rK, rQ, rJ, rT, r9, r8, r7, r6, r5, r4, r3, r2}
var suitList = cardSuitList{s, c, h, d}

type twoCardCombo struct {
	highCard *card
	lowCard  *card
	// removed, an idicator if one of the cards has been seen, and therefore should
	// be removed from the combo matrx, ie. perform card removal.
	removed bool
}

type twoCardComboList [1326]twoCardCombo

//
const (
	s cardSuit = "s"
	c cardSuit = "c"
	h cardSuit = "h"
	d cardSuit = "d"

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
)

// type CardDeck [52]card

type cdmKey struct {
	cr cardRank
	cs cardSuit
}

type cardDeckMap map[cdmKey]*card

type orderedListOfPtrsToCards [52]*card

/*
#########################################################################
#########################################################################
#########################################################################

 ######   ##        #######  ########     ###    ##        ######
##    ##  ##       ##     ## ##     ##   ## ##   ##       ##    ##
##        ##       ##     ## ##     ##  ##   ##  ##       ##
##   #### ##       ##     ## ########  ##     ## ##        ######
##    ##  ##       ##     ## ##     ## ######### ##             ##
##    ##  ##       ##     ## ##     ## ##     ## ##       ##    ##
 ######   ########  #######  ########  ##     ## ########  ######

#########################################################################
#########################################################################
#########################################################################
*/
const version string = "v0.1"

var cdm cardDeckMap
var orderedList orderedListOfPtrsToCards
var tcclPtr *twoCardComboList

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

// #####################################################################
// makeCardDeckMap doc line
func makeCardDeckMap() (cdm cardDeckMap, ol orderedListOfPtrsToCards) {

	cdm = make(cardDeckMap)

	sequence := 1

	for _, rank := range rankList {

		for _, suit := range suitList {
			var c = card{rank, suit, false, sequence}
			cPtr := &c

			cdm[cdmKey{rank, suit}] = cPtr
			ol[sequence-1] = cPtr

			sequence++
		}
	}
	return cdm, ol
}

// #####################################################################
// createDeck doc line
func createTwoCardComboList() (tcclPtr *twoCardComboList) {

	var hcPtr, lcPtr *card
	tcclIndex := 0

	var tccl twoCardComboList
	tcclPtr = &tccl

	numberofcardsindeck := 52
	maxindex := numberofcardsindeck - 1

	for i, cPtr := range orderedList {
		hcPtr = orderedList[i]
		for j := i + 1; j <= maxindex; j++ {
			lcPtr = orderedList[j]
			tcclPtr[tcclIndex].highCard = hcPtr
			tcclPtr[tcclIndex].lowCard = lcPtr

			// Remove combos, in case some cards are already removed at this point
			if hcPtr.seen == true || lcPtr.seen == true {
				tcclPtr[tcclIndex].removed = true
			}
			tcclIndex++
		}
		fmt.Printf("2c list index: %v    Top card: %v\n", i, cPtr)

	}

	for index, tcc := range tcclPtr {
		// for index := 0; index < 100; index++ {
		// fmt.Printf("index: %v	element: %v\n", index, tccl[index])
		fmt.Printf("index: %v	hc: %v  lc: %v    removed: %v\n", index, tcc.highCard, tcc.lowCard, tcc.removed)
	}

	// fmt.Printf("list: \n%v\n\n\n", tcclPtr)
	return tcclPtr
}

// #####################################################################
// removeCard doc line
func removeCard(cr cardRank, cs cardSuit) {

	cPtr := cdm[cdmKey{cr, cs}]
	cPtr.seen = true

	// Adjust two card bombo list
	for i, tcc := range tcclPtr {
		if tcc.highCard.seen == true || tcc.lowCard.seen == true {
			tcclPtr[i].removed = true
		}
	}
}

/*
#########################################################################
#########################################################################
#########################################################################

##     ##    ###    #### ##    ##
###   ###   ## ##    ##  ###   ##
#### ####  ##   ##   ##  ####  ##
## ### ## ##     ##  ##  ## ## ##
##     ## #########  ##  ##  ####
##     ## ##     ##  ##  ##   ###
##     ## ##     ## #### ##    ##

#########################################################################
#########################################################################
#########################################################################
*/

func main() {

	cdm, orderedList = makeCardDeckMap()
	// fmt.Println(cdm[cdmKey{r2, s}])

	for i, cPtr := range cdm {
		fmt.Printf("index: %v card: %v\n", i, cPtr)
	}

	fmt.Println()
	fmt.Printf("map len: %v\n", len(cdm))

	for i, cPtr := range orderedList {
		fmt.Printf("index: %v card: %v\n", i, cPtr)
	}

	for i, cPtr := range orderedList {
		fmt.Printf("index: %v card: %v\n", i, cPtr)
	}

	tcclPtr = createTwoCardComboList()

	removeCard(r2, s)
	removeCard(rA, s)

	for index, tcc := range tcclPtr {
		// for index := 0; index < 100; index++ {
		// fmt.Printf("index: %v	element: %v\n", index, tccl[index])
		fmt.Printf("index: %v	hc: %v  lc: %v    removed: %v\n", index, tcc.highCard, tcc.lowCard, tcc.removed)
	}
}
