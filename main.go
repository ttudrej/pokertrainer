package main

import "fmt"

type cardRank string
type cardSuit string

/*
type rankCounter struct {
	rank    cardRank
	counter int
}
*/

type cardRankList [13]cardRank
type cardRankListFull [14]cardRank
type cardSuitList [4]cardSuit

type card struct {
	rank     cardRank
	suit     cardSuit
	seen     bool
	sequence int
}

var rankList = cardRankList{rA, rK, rQ, rJ, rT, r9, r8, r7, r6, r5, r4, r3, r2}
var rankListFull = cardRankListFull{rA, rK, rQ, rJ, rT, r9, r8, r7, r6, r5, r4, r3, r2, rA}
var suitList = cardSuitList{s, c, h, d}

type twoCardCombo struct {
	highCard card
	lowCard  card
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

type rankCtrMap map[cardRank]int
type suitCtrMap map[cardSuit]int

// rankCouter keeps track of counts of specific rank in a card list, and track of poker hands.
type rankCounter struct {
	max    int
	top4x1 cardRank // rank of the top quads

	top3x1 cardRank // rank of the top trips / three of a kind
	top3x2 cardRank // rank of the 2nd trips / three of a kind

	top2x1 cardRank // rank of the top pair
	top2x2 cardRank // rank of the second pair ...
	top2x3 cardRank

	top1x1 cardRank
	top1x2 cardRank
	top1x3 cardRank
	top1x4 cardRank
	top1x5 cardRank

	rcm rankCtrMap // hold how many of each rank are there in a card list
}

type suitCounter struct {
	max int
	scm suitCtrMap
}

// orderedListOfPtrsToCard
type orderedListOfPtrsToCards [52]*card

// orderedListOfPtrsToCard uses 56 not 52 slots, to accomodate for the Aces in 5-A straights
// Used for hand rank checks ONLY
type orderedListFullOfPtrsToCards [56]*card

type flop struct {
	f1 card
	f2 card
	f3 card
}

// type representing the community cards
type board struct {
	f flop
	t card
	r card
}

// The following are needed to represent sets of community cards, for use in
// determining what hands and draws we have
type cardList5 [5]*card
type cardList6 [6]*card
type cardList7 [7]*card

type cardList [7]*card // Needs at most the max numbe of commuity cards + max num of the hole cards, so around 7 for NLH

/*
// For keeping track of how many cards of a partiular rank we have in the list
type rankCounter struct {
	rAct int
	rKct int
	rQct int
	rJct int
	rTct int
	r9ct int
	r8ct int
	r7ct int
	r6ct int
	r5ct int
	r4ct int
	r3ct int
	r2ct int
	max  int // set to the larget value of individual counters
}

// For keeping track of how many cards of a partiular suit we have in the list
type suitCounter struct {
	sCt int
	cCt int
	hCt int
	dCt int
	max int // set to the larget value of individual counters
}
*/

/*
i: 0  card: &{A s false 1}
i: 1  card: &{A c false 2}
i: 2  card: &{A h false 3}
i: 3  card: &{A d false 4}
i: 4  card: &{K s false 5}
i: 5  card: &{K c false 6}
i: 6  card: &{K h false 7}
i: 7  card: &{K d false 8}
i: 8  card: &{Q s false 9}
i: 9  card: &{Q c false 10}
i: 10  card: &{Q h false 11}
i: 11  card: &{Q d false 12}
i: 12  card: &{J s false 13}
i: 13  card: &{J c false 14}
i: 14  card: &{J h false 15}
i: 15  card: &{J d false 16}
i: 16  card: &{T s false 17}
i: 17  card: &{T c false 18}
i: 18  card: &{T h false 19}
i: 19  card: &{T d false 20}
i: 20  card: &{9 s false 21}
i: 21  card: &{9 c false 22}
i: 22  card: &{9 h false 23}
i: 23  card: &{9 d false 24}
i: 24  card: &{8 s false 25}
i: 25  card: &{8 c false 26}
i: 26  card: &{8 h false 27}
i: 27  card: &{8 d false 28}
i: 28  card: &{7 s false 29}
i: 29  card: &{7 c false 30}
i: 30  card: &{7 h false 31}
i: 31  card: &{7 d false 32}
i: 32  card: &{6 s false 33}
i: 33  card: &{6 c false 34}
i: 34  card: &{6 h false 35}
i: 35  card: &{6 d false 36}
i: 36  card: &{5 s false 37}
i: 37  card: &{5 c false 38}
i: 38  card: &{5 h false 39}
i: 39  card: &{5 d false 40}
i: 40  card: &{4 s false 41}
i: 41  card: &{4 c false 42}
i: 42  card: &{4 h false 43}
i: 43  card: &{4 d false 44}
i: 44  card: &{3 s false 45}
i: 45  card: &{3 c false 46}
i: 46  card: &{3 h false 47}
i: 47  card: &{3 d false 48}
i: 48  card: &{2 s false 49}
i: 49  card: &{2 c false 50}
i: 50  card: &{2 h false 51}
i: 51  card: &{2 d false 52}
*/

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
var orderedListFull orderedListFullOfPtrsToCards

var testCardList5 cardList5

var testCardList cardList

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
// createDeck makes us a brand new deck, and also gives an ordered list of the cards in it.
func createDeck() (cdm cardDeckMap, ol orderedListOfPtrsToCards, olf orderedListFullOfPtrsToCards) {

	cdm = make(cardDeckMap)

	sequence := 1 // bottom card in the deck, card 52 is the top of the deck. Dealing starts from the top.

	for _, rank := range rankList {

		for _, suit := range suitList {
			var c = card{rank, suit, false, sequence}
			cPtr := &c

			cdm[cdmKey{rank, suit}] = cPtr
			ol[sequence-1] = cPtr
			olf[sequence-1] = cPtr

			// Also point at Aces that fit below the dueces. Needed for woking out Straigh relative ranking
			if rank == rA {
				olf[sequence-1+52] = cPtr
			}
			sequence++
		}
	}

	return cdm, ol, olf
}

// #####################################################################
// createRange doc line
func createRange() (tcclPtr *twoCardComboList) {

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
			tcclPtr[tcclIndex].highCard = *hcPtr
			tcclPtr[tcclIndex].lowCard = *lcPtr

			// Remove combos, in case some cards are already removed at this point
			if hcPtr.seen == true || lcPtr.seen == true {
				tcclPtr[tcclIndex].removed = true
			}
			tcclIndex++
		}
		fmt.Printf("i: %v  card: %v\n", i, cPtr)

	}
	/*
		for index, tcc := range tcclPtr {
			// for index := 0; index < 100; index++ {
			// fmt.Printf("index: %v	element: %v\n", index, tccl[index])
			fmt.Printf("index: %v	hc: %v  lc: %v    removed: %v\n", index, tcc.highCard, tcc.lowCard, tcc.removed)
		}
	*/
	// fmt.Printf("list: \n%v\n\n\n", tcclPtr)
	return tcclPtr
}

// #####################################################################
// seeCardInDeck marks the card in the deck as seen / epxposed.
func seeCardInDeck(cr cardRank, cs cardSuit) {

	cPtr := cdm[cdmKey{cr, cs}]
	cPtr.seen = true
}

// #####################################################################
//

// #####################################################################
// removeCardFromRange doc line
func removeCardFromRange(cr cardRank, cs cardSuit, rangePtr *twoCardComboList) (numRemoved int) {
	numRemoved = 0

	for i, tcc := range rangePtr {
		if (tcc.removed == false) &&
			((tcc.highCard.rank == cr && tcc.highCard.suit == cs) ||
				(tcc.lowCard.rank == cr && tcc.lowCard.suit == cs)) {
			rangePtr[i].removed = true
			numRemoved++
		}
	}
	return numRemoved
}

// #####################################################################
// deal the Flop
// func (f flop) deal() (f1, f2, f3 card) {
func dealFlop() (f flop) {
	done := false

	for i := 51; i >= 0; i-- {

		if done == false && orderedList[i].seen == false {
			f.f1 = *orderedList[i]
			orderedList[i].seen = true
			f.f2 = *orderedList[i-1]
			orderedList[i-1].seen = true
			f.f3 = *orderedList[i-2]
			orderedList[i-2].seen = true
			fmt.Printf("dealing the flop: %v %v %v\n", f.f1, f.f2, f.f3)

			done = true
		}
	}
	return f
}

func dealOne() (c card) {
	done := false

	for i := 51; i >= 0; i-- {

		if done == false && orderedList[i].seen == false {
			c = *orderedList[i]
			orderedList[i].seen = true

			fmt.Printf("dealing one card: %v\n", c)

			done = true
		}
	}
	return c
}

// #####################################################################
// findBestHandIn5 examines the hand, and looks for hand rank, top to bottom, quits as soon as it finds one.
// By convention, index 0 and 1 are the hole cards.
func findBestHandIn5(lPtr *cardList5) (handName string) {

	var handNamePtr *string
	handNamePtr = &handName

	switch {
	// check for the top SF
	// case findSF(lPtr, handNamePtr):
	// case find4xIn5(lPtr, handNamePtr):
	// case findFH(lPtr, handNamePtr):
	// case findFl(lPtr, handNamePtr):
	// case findSt(lPtr, handNamePtr):
	// case find3x(lPtr, handNamePtr):
	// case find2x2(lPtr, handNamePtr):

	// case find2x(lPtr, handNamePtr):
	default:
		*handNamePtr = "default case"
	}
	return handName
}

// #####################################################################
// findBestHand examines the hand, and looks for hand rank, top to bottom, quits as soon as it finds one.
// By convention, index 0 and 1 are the hole cards.
func findBestHand(lPtr *cardList, rcPtr *rankCounter, scPtr *suitCounter) (handName string) {

	var handNamePtr *string
	handNamePtr = &handName

	switch {
	// check for the top SF
	// case findSF(lPtr, handNamePtr):
	case find4x(lPtr, rcPtr, handNamePtr):
	// case findFH(lPtr, handNamePtr):
	// case findFl(lPtr, handNamePtr):
	// case findSt(lPtr, handNamePtr):
	// case find3x(lPtr, handNamePtr):
	// case find2x2(lPtr, handNamePtr):

	// case find2x(lPtr, handNamePtr):
	default:
		*handNamePtr = "default case"
	}
	return handName
}

// #####################################################################
// countRanksInCardList - doc line
func countRanksInCardList(clPtr *cardList) (rcPtr *rankCounter) {

	fmt.Println("clPtr0: ", clPtr[0])

	var myRcm rankCtrMap
	myRcm = make(rankCtrMap)
	var rc rankCounter
	rcPtr = &rc
	rcPtr.rcm = myRcm

	for _, cPtr := range clPtr {

		// work only on the defined cards in the list
		if cPtr != nil {
			fmt.Printf("%v\n", *cPtr)
			rcPtr.rcm[cPtr.rank]++

			fmt.Printf("c.rank: %v; count: %v\n", cPtr.rank, rcPtr.rcm[cPtr.rank])

			if rcPtr.rcm[cPtr.rank] > rcPtr.max {
				rcPtr.max = rcPtr.rcm[cPtr.rank]
			}
		}
		fmt.Println("max: ", rcPtr.max)
	}

	fmt.Println()
	fmt.Println("Counts:")
	fmt.Printf("rcPtr %v: \n", rcPtr)

	return rcPtr
}

// #####################################################################
// countSuitsInCardList - doc line
func countSuitsInCardList(clPtr *cardList) (scPtr *suitCounter) {

	fmt.Println("clPtr0: ", clPtr[0])

	var myScm suitCtrMap
	myScm = make(suitCtrMap)
	var sc suitCounter
	scPtr = &sc
	scPtr.scm = myScm

	for _, cPtr := range clPtr {

		// work only on the defined cards in the list
		if cPtr != nil {
			fmt.Printf("%v\n", *cPtr)
			scPtr.scm[cPtr.suit]++

			fmt.Printf("cPtr.suit: %v; count: %v\n", cPtr.suit, scPtr.scm[cPtr.suit])

			if scPtr.scm[cPtr.suit] > scPtr.max {
				scPtr.max = scPtr.scm[cPtr.suit]
			}
		}
		fmt.Println("max: ", scPtr.max)
	}

	fmt.Println("Counts:")
	fmt.Println("scPtr[s]: ", scPtr.scm[s])
	fmt.Println("scPtr[c]: ", scPtr.scm[c])
	fmt.Println("scPtr[h]: ", scPtr.scm[h])
	fmt.Println("scPtr[d]: ", scPtr.scm[d])
	fmt.Println("scPtr.max: ", scPtr.max)
	fmt.Println()

	return scPtr
}

// #####################################################################
// findTopRanks "sorts" the same types of hands, ie. figurs out the top quads, top trips, ...
// For now, a version that deals with up to 7 cards in the list, for one player, should do.
func findTopRanks(rcPtr *rankCounter) {

	switch {
	case rcPtr.max == 4:

		for _, rank := range rankList {
			if rcPtr.rcm[rank] == 4 { // we have a 4x
				rcPtr.top4x1 = rank
			}
			switch {
			case rcPtr.top1x1 == "" && rcPtr.rcm[rank] == 3:
				rcPtr.top1x1 = rank
				rcPtr.top1x2 = rank
				rcPtr.top1x3 = rank
			case rcPtr.top1x1 != "" && rcPtr.top1x2 == "" && rcPtr.rcm[rank] == 3:
				rcPtr.top1x2 = rank
				rcPtr.top1x3 = rank
				rcPtr.top1x4 = rank
			case rcPtr.top1x2 != "" && rcPtr.top1x3 == "" && rcPtr.rcm[rank] == 3:
				rcPtr.top1x3 = rank
				rcPtr.top1x4 = rank
				rcPtr.top1x5 = rank
			case rcPtr.top1x1 == "" && rcPtr.rcm[rank] == 2:
				rcPtr.top1x1 = rank
				rcPtr.top1x2 = rank
			case rcPtr.top1x1 != "" && rcPtr.top1x2 == "" && rcPtr.rcm[rank] == 2:
				rcPtr.top1x2 = rank
				rcPtr.top1x3 = rank
			case rcPtr.top1x2 != "" && rcPtr.top1x3 == "" && rcPtr.rcm[rank] == 2:
				rcPtr.top1x3 = rank
				rcPtr.top1x4 = rank
			case rcPtr.top1x3 != "" && rcPtr.top1x4 == "" && rcPtr.rcm[rank] == 2:
				rcPtr.top1x4 = rank
				rcPtr.top1x5 = rank
			case rcPtr.top1x1 == "" && rcPtr.rcm[rank] == 1:
				rcPtr.top1x1 = rank
			case rcPtr.top1x1 != "" && rcPtr.top1x2 == "" && rcPtr.rcm[rank] == 1:
				rcPtr.top1x2 = rank
			case rcPtr.top1x2 != "" && rcPtr.top1x3 == "" && rcPtr.rcm[rank] == 1:
				rcPtr.top1x3 = rank
			case rcPtr.top1x3 != "" && rcPtr.top1x4 == "" && rcPtr.rcm[rank] == 1:
				rcPtr.top1x4 = rank
			case rcPtr.top1x4 != "" && rcPtr.top1x5 == "" && rcPtr.rcm[rank] == 1:
				rcPtr.top1x5 = rank
			default:

			}
		}
	}
}

// #####################################################################
// findSF Find Straight Flush - doc line
func findSF(lPtr *cardList5, handNamePtr *string) (found bool) {
	found = false

	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])

LSF:
	for i := 0; i <= 9; i++ {
		fmt.Printf("inFindSF rankListFull of i is: \n%v\n", rankListFull[i])
		fmt.Printf("lPtr 0 : %v; rankListFull of i : %v\n\n", lPtr[0].rank, rankListFull[i])

		if (lPtr[0].rank == rankListFull[i] || lPtr[1].rank == rankListFull[i] || lPtr[2].rank == rankListFull[i] || lPtr[3].rank == rankListFull[i] || lPtr[4].rank == rankListFull[i]) &&
			(lPtr[0].rank == rankListFull[i+1] || lPtr[1].rank == rankListFull[i+1] || lPtr[2].rank == rankListFull[i+1] || lPtr[3].rank == rankListFull[i+1] || lPtr[4].rank == rankListFull[i+1]) &&
			(lPtr[0].rank == rankListFull[i+2] || lPtr[1].rank == rankListFull[i+2] || lPtr[2].rank == rankListFull[i+2] || lPtr[3].rank == rankListFull[i+2] || lPtr[4].rank == rankListFull[i+2]) &&
			(lPtr[0].rank == rankListFull[i+3] || lPtr[1].rank == rankListFull[i+3] || lPtr[2].rank == rankListFull[i+3] || lPtr[3].rank == rankListFull[i+3] || lPtr[4].rank == rankListFull[i+3]) &&
			(lPtr[0].rank == rankListFull[i+4] || lPtr[1].rank == rankListFull[i+4] || lPtr[2].rank == rankListFull[i+4] || lPtr[3].rank == rankListFull[i+4] || lPtr[4].rank == rankListFull[i+4]) &&
			(lPtr[0].suit == lPtr[1].suit) &&
			(lPtr[0].suit == lPtr[2].suit) &&
			(lPtr[0].suit == lPtr[3].suit) &&
			(lPtr[0].suit == lPtr[4].suit) {

			fmt.Printf("inFindSF inside IF \n%v\n%v\n%v\n%v\n%v\n", rankListFull[i], rankListFull[i+1], rankListFull[i+2], rankListFull[i+3], rankListFull[i+4])

			found = true

			switch {
			case found == true:
				*handNamePtr = string(rankListFull[i]) + " high Straight Flush"
				break LSF
			// The default case is probably NOT needed here, since we take care of that in findBestHand the calling function
			default:
				*handNamePtr = "did not find SF"
			}
		}
		fmt.Println("###########")
	}

	return found
}

// #####################################################################
// find4x looks for quads - doc line
func find4xIn5(lPtr *cardList, rcPtr *rankCounter, handNamePtr *string) (found bool) {
	found = false

	if rcPtr.max == 4 {
		found = true
	}

	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
	// For the first 2 card, count up
L4x:
	for i := 0; i <= 1; i++ {
		rank1 := lPtr[i].rank
		count := 1
		firstIndex := i + 1
		fmt.Printf("################ i: %v\n", i)
		for _, cardNPtr := range lPtr[firstIndex:] {
			if cardNPtr.rank == rank1 {
				count++
			}
			fmt.Println("looking for 4x")
			fmt.Printf("rank1: %v \n count: %v \n", rank1, count)
			fmt.Println(cardNPtr.rank)
		}
		if count == 4 {
			found = true
		}

		switch {
		case found == true:
			*handNamePtr = "Quad " + string(rank1) + "s"
			break L4x
		// The default case is probably NOT needed here, since we take care of that in findBestHand the calling function
		default:
			*handNamePtr = "did not find 4x"
		}

	}

	return found
}

// #####################################################################
// find4x looks for quads - doc line
func find4x(lPtr *cardList, rcPtr *rankCounter, handNamePtr *string) (found bool) {
	found = false

	if rcPtr.max == 4 {
		found = true
	}
	return found
}

/*
// #####################################################################
// findFH looks for a Full House / Boat - doc line
func findFH(lPtr *cardList5, handNamePtr *string) (found bool) {
	found = false

	var rcA, rcK, rcQ, rcJ, rcT, rc9, rc8, rc7, rc6, rc5, rc4, rc3, rc2 rankCounter
	var rcList = [13]rankCounter{rcA, rcK, rcQ, rcJ, rcT, rc9, rc8, rc7, rc6, rc5, rc4, rc3, rc2}

	for _, rc := range rcList {
		fmt.Printf("rc.counter: %v\n", rc.counter)
	}

	####################


	// for 2nd through 5th card
	for i := 1; i <= 4; i++ {
		if lPtr[i].rank == rc1.rank {
			rc1.counter++
		} else {
			rank2 = lPtr[i].rank // We found a different rank
			countRank2++
		}
	}

	switch {
	case countRank1 == 3 && countRank2 == 2:
		found = true
		*handNamePtr = "Full House, " + string(rank1) + "s over " + string(rank2) + "s"
	case countRank1 == 2 && countRank2 == 3:
		found = true
		*handNamePtr = "Full House, " + string(rank2) + "s over " + string(rank1) + "s"
	default:
		*handNamePtr = "did not find a FH"
	}
	return found
}
*/

/*
// #####################################################################
// findFlush - doc line
func findFl(lPtr *cardList5, handNamePtr *string) (found bool) {
	found = false

	countS := 0
	countC := 0
	countH := 0
	countD := 0

	var mySuit cardSuit

	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])

	for i := 0; i <= 4; i++ {
		switch {
		case lPtr[i].suit == s:
			countS++
		case lPtr[i].suit == c:
			countC++
		case lPtr[i].suit == h:
			countH++
		case lPtr[i].suit == d:
			countD++
		}
	}

	if countS == 5 || countC == 5 || countH == 5 || countD == 5 {
		found = true

		switch {
		case countS == 5:
			mySuit = s
		case countC == 5:
			mySuit = c
		case countH == 5:
			mySuit = h
		case countD == 5:
			mySuit = d
		}
		*handNamePtr = "Flush of " + string(mySuit)
	}
	return found
}
*/

/*
// #####################################################################
// findSt looks for straights - doc line
func findSt(lPtr *cardList5, handNamePtr *string) (found bool) {
	found = false

	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
LSt:
	for i := 0; i <= 10; i++ {
		fmt.Printf("inFindSF rankListFull of i is: \n%v\n", rankListFull[i])
		fmt.Printf("lPtr 0 : %v; rankListFull of i : %v\n\n", lPtr[0].rank, rankListFull[i])

		if (lPtr[0].rank == rankListFull[i] || lPtr[1].rank == rankListFull[i] || lPtr[2].rank == rankListFull[i] || lPtr[3].rank == rankListFull[i] || lPtr[4].rank == rankListFull[i]) &&
			(lPtr[0].rank == rankListFull[i+1] || lPtr[1].rank == rankListFull[i+1] || lPtr[2].rank == rankListFull[i+1] || lPtr[3].rank == rankListFull[i+1] || lPtr[4].rank == rankListFull[i+1]) &&
			(lPtr[0].rank == rankListFull[i+2] || lPtr[1].rank == rankListFull[i+2] || lPtr[2].rank == rankListFull[i+2] || lPtr[3].rank == rankListFull[i+2] || lPtr[4].rank == rankListFull[i+2]) &&
			(lPtr[0].rank == rankListFull[i+3] || lPtr[1].rank == rankListFull[i+3] || lPtr[2].rank == rankListFull[i+3] || lPtr[3].rank == rankListFull[i+3] || lPtr[4].rank == rankListFull[i+3]) &&
			(lPtr[0].rank == rankListFull[i+4] || lPtr[1].rank == rankListFull[i+4] || lPtr[2].rank == rankListFull[i+4] || lPtr[3].rank == rankListFull[i+4] || lPtr[4].rank == rankListFull[i+4]) {

			fmt.Printf("inFindSt inside IF \n%v\n%v\n%v\n%v\n%v\n", rankListFull[i], rankListFull[i+1], rankListFull[i+2], rankListFull[i+3], rankListFull[i+4])

			found = true

			switch {
			case found == true:
				*handNamePtr = string(rankListFull[i]) + " high Straight"
				break LSt
			// The default case is probably NOT needed here, since we take care of that in findBestHand the calling function,
			// but may be useful in debugging at some point
			default:
				*handNamePtr = "did not find St"
			}
		}
		fmt.Println("###########")
	}

	return found
}
*/

/*
// #####################################################################
// find3x looks for sets and trips / 3 of a kind - doc line
func find3x(lPtr *cardList5, handNamePtr *string) (found bool) {
	found = false

L3x:
	for i := 0; i <= 2; i++ {
		rank1 := lPtr[i].rank
		count := 1
		firstIndex := i + 1
		fmt.Printf("################ i: %v\n", i)
		for _, cardNPtr := range lPtr[firstIndex:] {
			if cardNPtr.rank == rank1 {
				count++
			}
			fmt.Println("looking for 3x")
			fmt.Printf("rank1: %v \n count: %v \n", rank1, count)
			fmt.Println(cardNPtr.rank)
		}
		if count == 3 {
			found = true
		}

		switch {
		case found == true:
			*handNamePtr = "Three of a kind, " + string(rank1) + "s"
			break L3x
		// The default case is probably NOT needed here, since we take care of that in findBestHand the calling function
		default:
			*handNamePtr = "did not find 3x"
		}

	}

	return found
}
*/

/*
// #####################################################################
// find2x2 looks for 2 pair hands - doc line
func find2x2(lPtr *cardList5, handNamePtr *string) (found bool) {
	found = false

	var rcA, rcK, rcQ, rcJ, rcT, rc9, rc8, rc7, rc6, rc5, rc4, rc3, rc2 rankCounter
	var rcList = [13]rankCounter{rcA, rcK, rcQ, rcJ, rcT, rc9, rc8, rc7, rc6, rc5, rc4, rc3, rc2}

	for _, rc := range rcList {
		fmt.Printf("rc.counter: %v\n", rc.counter)
	}

	return true
}
*/
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

	// Create a global card deck, since there will be only one, for any one hand played.
	cdm, orderedList, orderedListFull = createDeck()
	// fmt.Println(cdm[cdmKey{r2, s}])

	// boardPtr := &board{}

	for i, cPtr := range cdm {
		fmt.Printf("index: %v card: %v\n", i, cPtr)
	}

	fmt.Println()
	fmt.Printf("map len: %v\n", len(cdm))

	for i, cPtr := range orderedList {
		fmt.Printf("index: %v card: %v\n", i, cPtr)
	}

	heroRangePtr := createRange()
	villainRangePtr := createRange()

	// Do some card removal manipulations
	// seeCardInDeck(r2, s)
	numRemovedFromHerosRange := removeCardFromRange(r2, s, heroRangePtr)
	numRemovedFromVillainsRange := removeCardFromRange(r2, s, villainRangePtr)
	/*
		for index, tcc := range heroRangePtr {
			// for index := 0; index < 100; index++ {
			// fmt.Printf("index: %v	element: %v\n", index, tccl[index])
			fmt.Printf("index: %v	hc: %v  lc: %v    removed: %v\n", index, tcc.highCard, tcc.lowCard, tcc.removed)
		}
	*/
	fmt.Println()
	fmt.Println("removed from heros    r: ", numRemovedFromHerosRange)
	fmt.Println("removed from villains r: ", numRemovedFromVillainsRange)

	var communityCards board
	communityCards.f = dealFlop()
	communityCards.t = dealOne()
	communityCards.r = dealOne()

	/* // SF setup
	j := 39 // Legal index range is 0-39; 39 + 4*4(16) = 55(56-1)
	testCardList5[0] = orderedList[j]
	testCardList5[1] = orderedList[j+4]
	testCardList5[2] = orderedList[j+8]
	testCardList5[3] = orderedList[j+12]
	testCardList5[4] = orderedList[j+16]
	*/

	/*
		// 3x
		testCardList5[0] = orderedList[12]
		testCardList5[1] = orderedList[51]
		testCardList5[2] = orderedList[50]
		testCardList5[3] = orderedList[49]
		testCardList5[4] = orderedList[1]
	*/
	/*
		// Straight
		testCardList5[0] = orderedList[0]
		testCardList5[1] = orderedList[42]
		testCardList5[2] = orderedList[47]
		testCardList5[3] = orderedList[50]
		testCardList5[4] = orderedList[39]
	*/

	/*
		// Two pair
		testCardList5[0] = orderedList[0]
		testCardList5[1] = orderedList[1]
		testCardList5[2] = orderedList[4]
		testCardList5[3] = orderedList[5]
		testCardList5[4] = orderedList[39]

		fmt.Printf("\n%v\n%v\n%v\n%v\n%v\n", testCardList5[0], testCardList5[1], testCardList5[2], testCardList5[3], testCardList5[4])

		bestHand5 := findBestHandIn5(&testCardList5)
		fmt.Printf("\n\n\n")
		fmt.Println("best hand 5: ", bestHand5)
		fmt.Printf("\n%v\n%v\n%v\n%v\n%v\n", testCardList5[0], testCardList5[1], testCardList5[2], testCardList5[3], testCardList5[4])

	*/

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// !!! As a convention, any cardList MUST be ordered by rank, A to 2.
	// Functions that process RELY on the order being correct
	// This will simpify some code down the line, as certain assumptions will
	// be valid, and certain manipulations no longer necessary.
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	testCardList[0] = orderedList[0]
	testCardList[1] = orderedList[1]
	testCardList[2] = orderedList[2]
	testCardList[3] = orderedList[3]
	testCardList[4] = orderedList[39]

	rankCtrPtr := countRanksInCardList(&testCardList)
	fmt.Println()
	fmt.Println(rankCtrPtr)

	suitCtrPtr := countSuitsInCardList(&testCardList)
	fmt.Println(suitCtrPtr)
	/*
		// Find the top 4x, 3x, 2x2, 2x, ...
		findTopRanks(rankCtrPtr)

		bestHand := findBestHand(&testCardList, rankCtrPtr, suitCtrPtr)
		fmt.Println("bestHand: ", bestHand)

		fmt.Printf("\n\n\n")
		fmt.Println("best hand ALL: ", bestHand)
		fmt.Printf("\n%v\n%v\n%v\n%v\n%v\n", testCardList[0], testCardList[1], testCardList[2], testCardList[3], testCardList[4])
	*/
}
