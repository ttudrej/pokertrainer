package main

import (
	"errors"
	"fmt"

	"github.com/ttudrej/pokertrainer/deck"

	"github.com/ttudrej/pokertrainer/hand_ranking"
)

type rankCtrMap map[deck.CardRank]int
type suitCtrMap map[deck.CardSuit]int

// rankCouter keeps track of counts of specific rank in a card list, and track of poker hands.
type rankCounter struct {
	max         int
	uniqeRankCt int           // count of all unique ranks in the list
	top4x1      deck.CardRank // rank of the top quads

	top3x1 deck.CardRank // rank of the top trips / three of a kind
	top3x2 deck.CardRank // rank of the 2nd trips / three of a kind

	top2x1 deck.CardRank // rank of the top pair
	top2x2 deck.CardRank // rank of the second pair ...
	top2x3 deck.CardRank

	top1x1 deck.CardRank
	top1x2 deck.CardRank
	top1x3 deck.CardRank
	top1x4 deck.Ca
	top1x5 deck.CardRank

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

type cardList [7]*card // Needs at most the max numbe of commuity cards + max num of the hole cards, so around 7 for NLH

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

type playerIDType int

type tableSeat struct {
	seatNumber int          // 1..10, depending on game
	playerID   playerIDType // use "0" for UN-occupied / empty, 1-infinity otherwise.
	statckSize int          // >= 0
	maxSeats   int          // 2..10. 2 - heads up, 6 - 6max, 9,10 - 9/10 handed.
}

type gameName string
type gameType string

// game size, expressed in cents ($0.01). A 1c/2c is "NL 2", a 50c/$100 is "NL 100", a $1/$2 is "NL 200", ...
type gameSize struct {
	name   gameName
	blind1 int // small blind in $0.01
	blind2 int // big blind
}

type pokerTable struct {
	numberSeats    int // 2-10, usually
	buttonPos      int
	tableNumber    int // table serial number, so we can keep track of multiple tables, if needed.
	kind           gameType
	size           gameSize
	seat1playerID  playerIDType
	seat2playerID  playerIDType
	seat3playerID  playerIDType
	seat4playerID  playerIDType
	seat5playerID  playerIDType
	seat6playerID  playerIDType
	seat7playerID  playerIDType
	seat8playerID  playerIDType
	seat9playerID  playerIDType
	seat10playerID playerIDType
}

type player struct {
	name                string
	id                  int
	behaviorDescription string

	// Starting ranges; sr9h = starting range 9 handed
	// The positions are relative to the button, going couter clockwise.
	// When players step away from the tble, the first relatie position to
	// be removed will be the U0, then U1, etc.
	// We choose to assign the position names this way on "non full" tables,
	// since it's the number of players behing you, that is considered most critical.

	sr9hU0open0limpers twoCardComboList

	sr9hU1open0limpers   twoCardComboList // EP
	sr9hU1open1limpersEp twoCardComboList

	sr9hU2open0limpers     twoCardComboList // EP/MP
	sr9hU2open1limpersEp   twoCardComboList
	sr9hU2open2limpersEpEp twoCardComboList

	sr9hU3open0limpers     twoCardComboList // MP
	sr9hU3open1limpersMp   twoCardComboList // The one Mp limper will be U2, one befor
	sr9hU3open2limpersEpEp twoCardComboList
	sr9hU3open2limpersEpMp twoCardComboList

	sr9hHJopen0limpers     twoCardComboList // MP/LP, up to 4 players in front
	sr9hHJopen1limpersEp   twoCardComboList
	sr9hHJopen1limpersMp   twoCardComboList
	sr9hHJopen2limpersEpEp twoCardComboList
	sr9hHJopen2limpersEpMp twoCardComboList
	sr9hHJopen2limpersMpMp twoCardComboList

	sr9hCOopen0limpers     twoCardComboList // LP
	sr9hCOopen1limpersEp   twoCardComboList
	sr9hCOopen1limpersMp   twoCardComboList
	sr9hCOopen2limpersEpEp twoCardComboList
	sr9hCOopen2limpersEpMp twoCardComboList
	sr9hCOopen2limpersEpLp twoCardComboList
	sr9hCOopen2limpersMpMp twoCardComboList
	sr9hCOopen2limpersMpLp twoCardComboList

	sr9hBTopen0limpers     twoCardComboList //
	sr9hBTopen1limpersEp   twoCardComboList
	sr9hBTopen1limpersMp   twoCardComboList
	sr9hBTopen1limpersLp   twoCardComboList
	sr9hBTopen2limpersEpEp twoCardComboList
	sr9hBTopen2limpersEpMp twoCardComboList
	sr9hBTopen2limpersEpLp twoCardComboList
	sr9hBTopen2limpersMpMp twoCardComboList
	sr9hBTopen2limpersMpLp twoCardComboList
	sr9hBTopen2limpersLpLP twoCardComboList

	sr9hSBopen0limpers     twoCardComboList
	sr9hSBopen1limpersEp   twoCardComboList // 1 limper in EP
	sr9hSBopen1limpersMp   twoCardComboList // 1 limper in MP ...
	sr9hSBopen1limpersLp   twoCardComboList
	sr9hSBopen2limpersEpEp twoCardComboList
	sr9hSBopen2limpersEpMp twoCardComboList
	sr9hSBopen2limpersEpLp twoCardComboList
	sr9hSBopen2limpersMpMp twoCardComboList
	sr9hSBopen2limpersMpLp twoCardComboList
	sr9hSBopen2limpersLpLp twoCardComboList

	sr9hBBopen0limpers     twoCardComboList
	sr9hBBopen1limpersEp   twoCardComboList // 1 limper in EP
	sr9hBBopen1limpersMp   twoCardComboList // 1 limper in MP ...
	sr9hBBopen1limpersLp   twoCardComboList
	sr9hBBopen1limpersSB   twoCardComboList
	sr9hBBopen2limpersEpEp twoCardComboList
	sr9hBBopen2limpersEpMp twoCardComboList
	sr9hBBopen2limpersEpLp twoCardComboList
	sr9hBBopen2limpersEpSB twoCardComboList
	sr9hBBopen2limpersMpMp twoCardComboList
	sr9hBBopen2limpersMpLp twoCardComboList
	sr9hBBopen2limpersLpLp twoCardComboList
	sr9hBBopen2limpersLpSB twoCardComboList
}

/*
type subRangeCombosPair struct { // sub in subRange refers to a subsection of the reange.
	numNotSeen int // 0-6

	cs twoCardCombo // club - spade
	ch twoCardCombo // club - heart
	cd twoCardCombo // club - diamond

	sh twoCardCombo // spade - heart
	sd twoCardCombo // spade - diamond

	hd twoCardCombo // heart - diamond
}

type subRangeCombosSuited struct {
	numNotSeen int // 0-4

	cc twoCardCombo // clubs
	ss twoCardCombo // spades
	hh twoCardCombo // hearts
	dd twoCardCombo // diamonds
}

type subRangeCombosUnSuited struct {
	numNotSeen int // 0-12

	sc twoCardCombo
	sh twoCardCombo
	sd twoCardCombo

	cs twoCardCombo
	ch twoCardCombo
	cd twoCardCombo

	hs twoCardCombo
	hc twoCardCombo
	hd twoCardCombo

	ds twoCardCombo
	dc twoCardCombo
	dh twoCardCombo
}
*/

var crm deck.CardRankMap

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

//
const (
	nl2   gameName = "NL 1c/2c"
	nl4   gameName = "NL 2c/4c"
	nl200 gameName = "NL $1/$2"
	nl300 gameName = "NL $1/$3"
	nl500 gameName = "NL $2/$5"

	nlh gameType = "NLH"
)

/*
var sfInfo FiveCardHandKindRanking
var x4Info FiveCardHandKindRanking
var fhInfo FiveCardHandKindRanking
var flInfo FiveCardHandKindRanking
var stInfo FiveCardHandKindRanking
var x3Info FiveCardHandKindRanking
var x22Info FiveCardHandKindRanking
var x2Info FiveCardHandKindRanking
var hcInfo FiveCardHandKindRanking
*/

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
// createClPtr - doc line
func createClPtr() (clPtr *cardList) {
	var cl cardList
	clPtr = &cl

	return clPtr
}

// firstFuncToTestWith doc line
func firstFuncToTestWith(s string) string {
	return s + s
}

// #####################################################################
// createRange doc line; creates a representation of the range.
func createRange() (tcclPtr *types.twoCardComboList, err error) {

	fmt.Println("### Starting createRange ###")

	var c1Ptr, c2Ptr *card
	tcclIndex := 0

	var tccl twoCardComboList
	tcclPtr = &tccl

	numberofcardsindeck := 52
	maxindex := numberofcardsindeck - 1

	for i, cPtr := range orderedList {
		c1Ptr = orderedList[i]
		for j := i + 1; j <= maxindex; j++ {
			c2Ptr = orderedList[j]
			tcclPtr[tcclIndex].c1Ptr = c1Ptr
			tcclPtr[tcclIndex].c2Ptr = c2Ptr
			tcclPtr[tcclIndex].inRange = false
			tcclPtr[tcclIndex].removed = false

			// Remove combos, in case some cards are already removed at this point
			if c1Ptr.community == true || c2Ptr.community == true {
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
			fmt.Printf("index: %v	hc: %v  lc: %v    removed: %v\n", index, tcc.c1, tcc.c2, tcc.removed)
		}
	*/
	// fmt.Printf("list: \n%v\n\n\n", tcclPtr)
	return tcclPtr, err
}

// #####################################################################
// allSeeCardInDeck marks the card in the deck as seen, ie, a community card seen by all players
func allSeeCardInDeck(cr deck.CardRank, cs deck.CardSuit) {
	fmt.Println("### Starting allSseeCardInDeck ###")
	cPtr := cdm[cdmKey{cr, cs}]
	cPtr.community = true
}

/*
// #####################################################################
// heroSeesCardInDeck marks the card in the deck as seen / epxposed.
func heroSeesCardInDeck(cr deck.CardRank, cs cardSuit) {
	fmt.Println("### Starting heroSeesCardInDeck ###")
	cPtr := cdm[cdmKey{cr, cs}]
	cPtr.seenByHero = true
}

// #####################################################################
// villainSeesCardInDeck marks the card in the deck as seen / epxposed.
func villainSeesCardInDeck(cr deck.CardRank, cs cardSuit) {
	fmt.Println("### Starting villainSeesCardInDeck ###")
	cPtr := cdm[cdmKey{cr, cs}]
	cPtr.seenByVillain = true
}

*/

// #####################################################################
//

// #####################################################################
// removeCardFromRange doc line; remove combos from tccl containing card.
func removeCardFromRange(cr deck.CardRank, cs deck.CardSuit, rangePtr *twoCardComboList) (numRemoved int, err error) {
	fmt.Println("### Starting removeCardFromRange ###")

	numRemoved = 0

	for i, tcc := range rangePtr {
		if (tcc.removed == false) && ((tcc.c1Ptr.rank == cr && tcc.c1Ptr.suit == cs) || (tcc.c2Ptr.rank == cr && tcc.c2Ptr.suit == cs)) {
			rangePtr[i].removed = true
			numRemoved++
		}
	}
	return numRemoved, err
}

// #####################################################################
// deal the Flop
// func (f flop) deal() (f1, f2, f3 card) {
func dealFlop() (f flop) {
	fmt.Println("### Starting dealFlop ###")
	done := false

	for i := 51; i >= 0; i-- {

		if done == false && orderedList[i].community == false && orderedList[i].dealtToPlayer == false {
			f.f1 = *orderedList[i]
			orderedList[i].community = true
			f.f2 = *orderedList[i-1]
			orderedList[i-1].community = true
			f.f3 = *orderedList[i-2]
			orderedList[i-2].community = true
			fmt.Printf("dealing the flop: %v %v %v\n", f.f1, f.f2, f.f3)

			done = true
		}
	}
	return f
}

// #####################################################################
// deal one card, for turns and rivers
func dealOne() (c card) {
	fmt.Println("### Starting dealOne ###")
	done := false

	for i := 51; i >= 0; i-- {

		if done == false && orderedList[i].community == false && orderedList[i].dealtToPlayer == false {
			c = *orderedList[i]
			orderedList[i].community = true

			fmt.Printf("dealing one card: %v\n", c)

			done = true
		}
	}
	return c
}

// #####################################################################
// getCardListLength doc line
func getCardListLength(clPtr *cardList) (l int) {
	fmt.Println("### Starting getCardListLen ###")

	// Figure out how many cards in list are defined / have been assigned
	l = 0

	for _, cPtr := range clPtr {
		if cPtr != nil {
			l++
		}
	}
	return l
}

// #####################################################################
// findBestHandInCardList examines the hand, and looks for hand rank, top to bottom, quits as soon as it finds one.
// By convention, index 0 and 1 are the hole cards.
// All functions used herein, use countRanksInCardList and countSuitsInCardList to as support functions
//
// In this release / for now, we'll ONLY consider card lists up to 7 cards long, 2 hole cards and 5 community cards
// In each case, we're looking for the best/top hand of it's type.
func findBestHandInCardList(clPtr *cardList) (handName string) {
	fmt.Println("### Starting findBestahqndInCardList ###")

	var handNamePtr *string
	handNamePtr = &handName

	// Figure out how many cards in list
	cardListLength := getCardListLength(clPtr)

	fmt.Println()
	fmt.Println("#### CARD LIST LEN: ", cardListLength)
	fmt.Println()

	rankCtrPtr := countRanksInCardList(clPtr)
	suitCtrPtr := countSuitsInCardList(clPtr)

	fmt.Println()
	fmt.Println("randCtrPtr: ", rankCtrPtr)
	fmt.Println("suitCtrPtr: ", suitCtrPtr)
	fmt.Println("dumping the CardList: ", clPtr[0], clPtr[1], clPtr[2], clPtr[3], clPtr[4], clPtr[5], clPtr[6])
	fmt.Println()

	switch {
	case findSF(clPtr, cardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
	case find4x(clPtr, cardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
	case findFH(clPtr, cardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
	case findFl(clPtr, cardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
	case findSt(clPtr, cardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
	case find3x(clPtr, cardListLength, rankCtrPtr, handNamePtr) == nil:
	case find2x2(clPtr, cardListLength, rankCtrPtr, handNamePtr) == nil:
	case find2x1(clPtr, cardListLength, rankCtrPtr, handNamePtr) == nil:
	case findHc(clPtr, cardListLength, rankCtrPtr, handNamePtr) == nil:
	default:
		*handNamePtr = "default case"
	}
	return handName
}

// #####################################################################
// countRanksInCardList - doc line
func countRanksInCardList(clPtr *cardList) (rcPtr *rankCounter) {
	fmt.Println("### Starting countRanksInCardList ###")

	fmt.Println("clPtr0: ", clPtr[0])

	var myRcm rankCtrMap
	myRcm = make(rankCtrMap)
	var rc rankCounter
	rcPtr = &rc
	rcPtr.rcm = myRcm

	rc.max = 0
	rc.uniqeRankCt = 0
	rc.top1x1 = rX // give it initial "unset" value
	rc.top1x2 = rX
	rc.top1x3 = rX
	rc.top1x4 = rX
	rc.top1x5 = rX
	rc.top2x1 = rX
	rc.top2x2 = rX
	rc.top2x3 = rX
	rc.top3x1 = rX
	rc.top3x2 = rX
	rc.top4x1 = rX

	for _, cPtr := range clPtr {

		// work only on the defined cards in the list
		if cPtr != nil {
			fmt.Printf("%v\n", *cPtr)
			rcPtr.rcm[cPtr.rank]++

			// fmt.Printf("c.rank: %v; count: %v\n", cPtr.rank, rcPtr.rcm[cPtr.rank])

			if rcPtr.rcm[cPtr.rank] > rcPtr.max {
				rcPtr.max = rcPtr.rcm[cPtr.rank]
			}
		}
		// fmt.Println("max: ", rcPtr.max)
	}

	// Count up how many different unique ranks there are in the list
	for _, rank := range rankList {
		if rcPtr.rcm[rank] > 0 {
			rcPtr.uniqeRankCt++
		}
	}

	fmt.Println()
	fmt.Println("Counts:")
	fmt.Printf("rcPtr %v: \n", rcPtr)

	return rcPtr
}

// #####################################################################
// countSuitsInCardList counts the cards in the 4 souits
func countSuitsInCardList(clPtr *cardList) (scPtr *suitCounter) {
	fmt.Println("### Starting countSuitsInCardList ###")

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

			// fmt.Printf("cPtr.suit: %v; count: %v\n", cPtr.suit, scPtr.scm[cPtr.suit])

			if scPtr.scm[cPtr.suit] > scPtr.max {
				scPtr.max = scPtr.scm[cPtr.suit]
			}
		}
		// fmt.Println("max: ", scPtr.max)
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
// findSF Find Straight Flush - doc line
// Looks for 5+ cards of same suit, then orders them, then looks for a straight
func findSF(clPtr *cardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {
	fmt.Println("### Starting findSF ###")

	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)
	fmt.Println("scPtr.max: ", scPtr.max)

	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
	// First check, the basics
	if cll < 5 || rcPtr.max > 4 || scPtr.max < 5 {
		// fmt.Println("Failed prelim SF checks, exiting findSF")
		err = errors.New("Failed prelim SF checks, exiting findSF")
		return err
	}

	flushClPtr, topSuit, err01 := find5OrMoreOfSameSuitInCardList(clPtr, cll, scPtr)

	if err01 != nil {
		fmt.Println(err01)
		err = errors.New("Trouble in find5OrMoreOfSameSuitInCardList")
		return err
	}

	flushClOrderedPtr, err02 := orderCardsOfSameSuit2(flushClPtr, topSuit)

	if err02 != nil {
		fmt.Println(err02)
		err = errors.New("Trouble in orderCardsOfSameSuit2")
		return err
	}

	rcFlPtr := countRanksInCardList(flushClOrderedPtr)
	scFlPtr := countSuitsInCardList(flushClOrderedPtr)

	err03 := findSt(flushClOrderedPtr, getCardListLength(flushClOrderedPtr), rcFlPtr, scFlPtr, handNamePtr)

	if err03 != nil {
		fmt.Println("err03: ", err03)
		err = errors.New("Did not find SF")
	} else {
		fmt.Println("Found SF XXXXXXXX")
		*handNamePtr = string(flushClOrderedPtr[0].rank) + " high Straight Flush"
	}
	return err
}

// #####################################################################
// find4x looks for quads - doc line
// func find4x(lPtr *cardList, rcPtr *rankCounter, handNamePtr *string) (found bool) {
func find4x(clPtr *cardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {
	fmt.Println("### Starting find4x ###")

	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)
	fmt.Println("scPtr.max: ", scPtr.max)

	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
	// First check, the basics
	if cll < 4 || rcPtr.max < 4 || rcPtr.uniqeRankCt > 4 {
		err = errors.New("Failed prelim 4x checks, exiting find4x")
		return err
	}

L01:
	for _, rank := range rankList {
		if rcPtr.rcm[rank] == 4 {
			rcPtr.top4x1 = rank
			break L01
		}
	}

	switch {
	case rcPtr.max == 4:
		*handNamePtr = "Four of a kind, " + string(rcPtr.top4x1) + "s"
	default:
		*handNamePtr = "did not find 4x"
		err = errors.New("err: Did not find 4x")
	}

	return err
}

// #####################################################################
// find3x looks for 3 of a kind - doc line
func find3x(clPtr *cardList, cll int, rcPtr *rankCounter, handNamePtr *string) (err error) {
	fmt.Println("### Starting find3x ###")

	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)

	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
	// First check, the basics
	if cll < 3 || rcPtr.max < 3 || rcPtr.uniqeRankCt > 5 {
		err = errors.New("Failed prelim 3x checks, exiting find3x")
		return err
	}

L03:
	for _, rank := range rankList {
		if rcPtr.rcm[rank] == 3 {
			rcPtr.top3x1 = rank
			break L03
		}
	}

	switch {
	case rcPtr.max == 3:
		*handNamePtr = "Three of a kind, " + string(rcPtr.top3x1) + "s"
	default:
		*handNamePtr = "did not find 3x"
		err = errors.New("err: Did not find 3x")
	}

	return err
}

// #####################################################################
// findFH looks for a Full House - doc line
func findFH(clPtr *cardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {
	fmt.Println("### Starting fineFH ###")

	fmt.Println("### Looking for a FH #######")
	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)
	fmt.Println("scPtr.max: ", scPtr.max)

	// Set default behavior
	*handNamePtr = "did not find a FH"
	err = errors.New("Did not find a FH")

	// First check, the basics
	if cll < 3 || rcPtr.max >= 6 {
		err = errors.New("Failed prelim FH checks, exiting findFH")
		fmt.Println("prelim check err: ", err)
		return err
	}

	if rcPtr.max == 3 {
		// Opting to NOT use (>= 3), so that we'll catch issues with 4x detection, if there is a 4x, and we got here anyway.
		// The idea is that 4x checker should have found 4x already, and we should NOT be in this section of code, if
		// that's the case.
		//
		// Possible configs of the 7 cards at this point can be:
		// 3x, 3x, 1x     *
		// 3x, 2x, 2x     *
		// 3x, 2x, 1x, 1x *
		// 3x, 1x, 1x, 1x, 1x
		// Got an FH, if we have one of the *-ed situations

		// Work through the possible configurations
		counter3x := 0
		counter2x := 0

		for _, rank := range rankList {

			switch {
			case rcPtr.rcm[rank] == 3: // found a 3x
				counter3x++

				switch {
				case counter3x == 1:
					rcPtr.top3x1 = rank
				case counter3x == 2: // 3x, 3x
					rcPtr.top3x2 = rank
					// found = true
				default:
				}

			case rcPtr.rcm[rank] == 2: // found a 2x
				counter2x++

				switch {
				case counter2x == 1:
					rcPtr.top2x1 = rank
					// found = true
				default:
				}

			default:
			} // End switch

		} // Since rankList is arranged from highest to lowest rank, the top3x1 and top3x2, ..., are properly set.

		switch {
		// case 3x, 3x, with 7 cards, there can't be a 2x here
		case rcPtr.top3x2 != rX:
			*handNamePtr = "FH, " + string(rcPtr.top3x1) + "s full of " + string(rcPtr.top3x2) + "s."
			// case 3x, 2x, 2x OR 3x, 2x, 1x, 1x
		case rcPtr.top2x1 != rX:
			*handNamePtr = "FH, " + string(rcPtr.top3x1) + "s full of " + string(rcPtr.top2x1) + "s."
		default:
			*handNamePtr = "did not find a FH"
			err = errors.New("Did not find a FH")
		}
	}
	return err
}

// #####################################################################
// findFl identifies the flush, and determines the the best 5 ranks that make it up. This allows us to compare flushes later.
func findFl(clPtr *cardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {
	fmt.Println("### Starting findFL ###")

	var mySuit deck.CardSuit

	var cr1, cr2, cr3, cr4, cr5 deck.CardRank
	cr1 = rX
	cr2 = rX
	cr3 = rX
	cr4 = rX
	cr5 = rX

	fmt.Println("### Looking for a Flush #######")
	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)
	fmt.Println("scPtr.max: ", scPtr.max)

	// First check, the basics
	if cll < 5 || scPtr.max < 5 {
		err = errors.New("Failed prelim Flush checks, exiting findFl")
		return err
	}

	switch {
	case scPtr.scm[s] >= 5:
		mySuit = s
	case scPtr.scm[c] >= 5:
		mySuit = c
	case scPtr.scm[h] >= 5:
		mySuit = h
	case scPtr.scm[d] >= 5:
		mySuit = d
	default:
		*handNamePtr = "did not find a Flush"
		err = errors.New("Did not find a Fl")
		return err
	}

	cr1, cr2, cr3, cr4, cr5 = orderCardsOfSameSuit(clPtr, mySuit)

	*handNamePtr = "Flush, " + string(cr1) + string(cr2) + string(cr3) + string(cr4) + string(cr5)
	return err
}

// orderCardsOfSameSuit takes 5 to 7 cards of the same suit, and returns a 5 card ordered list of ranks, high to low.
// The significance of "same suit" in this context is that there can be no 2 or more cards of the same rank.
//  Checks for a 5 high straight as well,
// to differentiate between an arbitrary A high flush and a 5 high Straight Flush, in which case it will return the A as the last card, not first.
func orderCardsOfSameSuit(clPtr *cardList, cs deck.CardSuit) (deck.CardRank, deck.CardRank, deck.CardRank, deck.CardRank, deck.CardRank) {
	fmt.Println("### Starting orderCardsOfSameSuit ###")

	var c card

	var myFlCard1, myFlCard2, myFlCard3, myFlCard4, myFlCard5, myFlCard6, myFlCard7 card
	myFlCard1 = card{rX, x, false, false, false, false, 0}
	myFlCard2 = card{rX, x, false, false, false, false, 0}
	myFlCard3 = card{rX, x, false, false, false, false, 0}
	myFlCard4 = card{rX, x, false, false, false, false, 0}
	myFlCard5 = card{rX, x, false, false, false, false, 0}
	myFlCard6 = card{rX, x, false, false, false, false, 0}
	myFlCard7 = card{rX, x, false, false, false, false, 0}

	for _, cPtr := range clPtr {

		c = *cPtr

		if c.suit == cs { // if we found a card of requested suit...

			if myFlCard1.rank == rX { // 1st card in the suit of the Fl
				myFlCard1 = *cPtr

			} else if myFlCard2.rank == rX { // 2nd card in the suit of the Fl
				if crm[c.rank] > crm[myFlCard1.rank] { // if new card is higher than the the first
					myFlCard2 = myFlCard1 // move the first to the second position
					myFlCard1 = c         // Set the first/top pos to c
				} else {
					myFlCard2 = c
				}

			} else if myFlCard3.rank == rX { // 3rd card in the suit of the Fl
				if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 2
					myFlCard3 = myFlCard2
					myFlCard2 = myFlCard1
					myFlCard1 = c
				} else if crm[c.rank] > crm[myFlCard2.rank] {
					myFlCard3 = myFlCard2
					myFlCard2 = c
				} else {
					myFlCard3 = c
				}

			} else if myFlCard4.rank == rX {
				if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 3
					myFlCard4 = myFlCard3
					myFlCard3 = myFlCard2
					myFlCard2 = myFlCard1
					myFlCard1 = c
				} else if crm[c.rank] > crm[myFlCard2.rank] {
					myFlCard4 = myFlCard3
					myFlCard3 = myFlCard2
					myFlCard2 = c
				} else if crm[c.rank] > crm[myFlCard3.rank] {
					myFlCard4 = myFlCard3
					myFlCard3 = c
				} else {
					myFlCard4 = c
				}

			} else if myFlCard5.rank == rX {
				if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 4
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = myFlCard2
					myFlCard2 = myFlCard1
					myFlCard1 = c
				} else if crm[c.rank] > crm[myFlCard2.rank] {
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = myFlCard2
					myFlCard2 = c
				} else if crm[c.rank] > crm[myFlCard3.rank] {
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = c
				} else if crm[c.rank] > crm[myFlCard4.rank] {
					myFlCard5 = myFlCard4
					myFlCard4 = c
				} else {
					myFlCard5 = c
				}

			} else if myFlCard6.rank == rX {
				if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 5
					myFlCard6 = myFlCard5
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = myFlCard2
					myFlCard2 = myFlCard1
					myFlCard1 = c
				} else if crm[c.rank] > crm[myFlCard2.rank] {
					myFlCard6 = myFlCard5
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = myFlCard2
					myFlCard2 = c
				} else if crm[c.rank] > crm[myFlCard3.rank] {
					myFlCard6 = myFlCard5
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = c
				} else if crm[c.rank] > crm[myFlCard4.rank] {
					myFlCard6 = myFlCard5
					myFlCard5 = myFlCard4
					myFlCard4 = c
				} else if crm[c.rank] > crm[myFlCard5.rank] {
					myFlCard6 = myFlCard5
					myFlCard5 = c
				} else {
					myFlCard6 = c
				}

			} else if myFlCard7.rank == rX {
				if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 6
					myFlCard7 = myFlCard6
					myFlCard6 = myFlCard5
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = myFlCard2
					myFlCard2 = myFlCard1
					myFlCard1 = c
				} else if crm[c.rank] > crm[myFlCard2.rank] {
					myFlCard7 = myFlCard6
					myFlCard6 = myFlCard5
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = myFlCard2
					myFlCard2 = c
				} else if crm[c.rank] > crm[myFlCard3.rank] {
					myFlCard7 = myFlCard6
					myFlCard6 = myFlCard5
					myFlCard5 = myFlCard4
					myFlCard4 = myFlCard3
					myFlCard3 = c
				} else if crm[c.rank] > crm[myFlCard4.rank] {
					myFlCard7 = myFlCard6
					myFlCard6 = myFlCard5
					myFlCard5 = myFlCard4
					myFlCard4 = c
				} else if crm[c.rank] > crm[myFlCard5.rank] {
					myFlCard7 = myFlCard6
					myFlCard6 = myFlCard5
					myFlCard5 = c
				} else if crm[c.rank] > crm[myFlCard6.rank] {
					myFlCard7 = myFlCard6
					myFlCard6 = c
				} else {
					myFlCard7 = c
				}
			}
		}
	}
	return myFlCard1.rank, myFlCard2.rank, myFlCard3.rank, myFlCard4.rank, myFlCard5.rank
}

// orderCardsOfSameSuit2 takes 5 to 7 cards of the same suit, and returns a 5 card ordered list of ranks, high to low.
// The significance of "same suit" in this context is that there can be no 2 cards of the same rank.
//
// Checks for a 5 high straight as well,
// to differentiate between an arbitrary A high flush and a 5 high Straight Flush, in which case it will return the A as the last card, not first.
func orderCardsOfSameSuit2(clPtr *cardList, cs deck.CardSuit) (resultingClPtr *cardList, err error) {
	fmt.Println("### Starting orderCardsOfSameSuit2 ###")

	fmt.Println("in orderCardsOfSameSuit2; clPtr: ", clPtr)

	resultingClPtr = createClPtr()

	var c card

	var myFlCard1, myFlCard2, myFlCard3, myFlCard4, myFlCard5, myFlCard6, myFlCard7 card
	myFlCard1 = card{rX, x, false, false, false, false, 0}
	myFlCard2 = card{rX, x, false, false, false, false, 0}
	myFlCard3 = card{rX, x, false, false, false, false, 0}
	myFlCard4 = card{rX, x, false, false, false, false, 0}
	myFlCard5 = card{rX, x, false, false, false, false, 0}
	myFlCard6 = card{rX, x, false, false, false, false, 0}
	myFlCard7 = card{rX, x, false, false, false, false, 0}

	myFlCard1Ptr := &myFlCard1
	myFlCard2Ptr := &myFlCard2
	myFlCard3Ptr := &myFlCard3
	myFlCard4Ptr := &myFlCard4
	myFlCard5Ptr := &myFlCard5
	myFlCard6Ptr := &myFlCard6
	myFlCard7Ptr := &myFlCard7

	for _, cPtr := range clPtr {

		if cPtr != nil {
			c = *cPtr
			fmt.Println("in orderCardsOfSameSuit2, card from range: ", c)

			if c.suit == cs { // if we found a card of requested suit...

				if myFlCard1.rank == rX { // 1st card in the suit of the Fl
					myFlCard1 = *cPtr

				} else if myFlCard2.rank == rX { // 2nd card in the suit of the Fl
					if crm[c.rank] > crm[myFlCard1.rank] { // if new card is higher than the the first
						myFlCard2 = myFlCard1 // move the first to the second position
						myFlCard1 = c         // Set the first/top pos to c
					} else {
						myFlCard2 = c
					}

				} else if myFlCard3.rank == rX { // 3rd card in the suit of the Fl
					if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 2
						myFlCard3 = myFlCard2
						myFlCard2 = myFlCard1
						myFlCard1 = c
					} else if crm[c.rank] > crm[myFlCard2.rank] {
						myFlCard3 = myFlCard2
						myFlCard2 = c
					} else {
						myFlCard3 = c
					}

				} else if myFlCard4.rank == rX {
					if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 3
						myFlCard4 = myFlCard3
						myFlCard3 = myFlCard2
						myFlCard2 = myFlCard1
						myFlCard1 = c
					} else if crm[c.rank] > crm[myFlCard2.rank] {
						myFlCard4 = myFlCard3
						myFlCard3 = myFlCard2
						myFlCard2 = c
					} else if crm[c.rank] > crm[myFlCard3.rank] {
						myFlCard4 = myFlCard3
						myFlCard3 = c
					} else {
						myFlCard4 = c
					}

				} else if myFlCard5.rank == rX {
					if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 4
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = myFlCard2
						myFlCard2 = myFlCard1
						myFlCard1 = c
					} else if crm[c.rank] > crm[myFlCard2.rank] {
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = myFlCard2
						myFlCard2 = c
					} else if crm[c.rank] > crm[myFlCard3.rank] {
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = c
					} else if crm[c.rank] > crm[myFlCard4.rank] {
						myFlCard5 = myFlCard4
						myFlCard4 = c
					} else {
						myFlCard5 = c
					}

				} else if myFlCard6.rank == rX {
					if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 5
						myFlCard6 = myFlCard5
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = myFlCard2
						myFlCard2 = myFlCard1
						myFlCard1 = c
					} else if crm[c.rank] > crm[myFlCard2.rank] {
						myFlCard6 = myFlCard5
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = myFlCard2
						myFlCard2 = c
					} else if crm[c.rank] > crm[myFlCard3.rank] {
						myFlCard6 = myFlCard5
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = c
					} else if crm[c.rank] > crm[myFlCard4.rank] {
						myFlCard6 = myFlCard5
						myFlCard5 = myFlCard4
						myFlCard4 = c
					} else if crm[c.rank] > crm[myFlCard5.rank] {
						myFlCard6 = myFlCard5
						myFlCard5 = c
					} else {
						myFlCard6 = c
					}

				} else if myFlCard7.rank == rX {
					if crm[c.rank] > crm[myFlCard1.rank] { // higher than the other 6
						myFlCard7 = myFlCard6
						myFlCard6 = myFlCard5
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = myFlCard2
						myFlCard2 = myFlCard1
						myFlCard1 = c
					} else if crm[c.rank] > crm[myFlCard2.rank] {
						myFlCard7 = myFlCard6
						myFlCard6 = myFlCard5
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = myFlCard2
						myFlCard2 = c
					} else if crm[c.rank] > crm[myFlCard3.rank] {
						myFlCard7 = myFlCard6
						myFlCard6 = myFlCard5
						myFlCard5 = myFlCard4
						myFlCard4 = myFlCard3
						myFlCard3 = c
					} else if crm[c.rank] > crm[myFlCard4.rank] {
						myFlCard7 = myFlCard6
						myFlCard6 = myFlCard5
						myFlCard5 = myFlCard4
						myFlCard4 = c
					} else if crm[c.rank] > crm[myFlCard5.rank] {
						myFlCard7 = myFlCard6
						myFlCard6 = myFlCard5
						myFlCard5 = c
					} else if crm[c.rank] > crm[myFlCard6.rank] {
						myFlCard7 = myFlCard6
						myFlCard6 = c
					} else {
						myFlCard7 = c
					}
				}
			}
		}
	}

	resultingClPtr[0] = myFlCard1Ptr
	resultingClPtr[1] = myFlCard2Ptr
	resultingClPtr[2] = myFlCard3Ptr
	resultingClPtr[3] = myFlCard4Ptr
	resultingClPtr[4] = myFlCard5Ptr

	switch {
	case myFlCard6.rank != rX:
		resultingClPtr[5] = myFlCard6Ptr
	case myFlCard7.rank != rX:
		resultingClPtr[6] = myFlCard7Ptr
	}

	fmt.Println("dumping the resultingClPtr: ", resultingClPtr[0], resultingClPtr[1], resultingClPtr[2], resultingClPtr[3], resultingClPtr[4], resultingClPtr[5], resultingClPtr[6])

	return resultingClPtr, err
}

// orderCardsOfMixedSuit takes 5 to 7 cards, irrespective of suit, and orders them by rank.
// Supports find2x2, find2x, and findHc, only.
// Also, fills in clPtr values, when possible, for the 2x... vars.
func orderCardsOfMixedSuit(clPtr *cardList) (resultingClPtr *cardList, err error) {
	fmt.Println("### Starting orderCardsOfMixedSuit ###")

	fmt.Println("in orderCardsOfMixedSuit; clPtr: ", clPtr)

	resultingClPtr = createClPtr()

	var c card

	var myCard1, myCard2, myCard3, myCard4, myCard5, myCard6, myCard7 card
	myCard1 = card{rX, x, false, false, false, false, 0}
	myCard2 = card{rX, x, false, false, false, false, 0}
	myCard3 = card{rX, x, false, false, false, false, 0}
	myCard4 = card{rX, x, false, false, false, false, 0}
	myCard5 = card{rX, x, false, false, false, false, 0}
	myCard6 = card{rX, x, false, false, false, false, 0}
	myCard7 = card{rX, x, false, false, false, false, 0}

	myCard1Ptr := &myCard1
	myCard2Ptr := &myCard2
	myCard3Ptr := &myCard3
	myCard4Ptr := &myCard4
	myCard5Ptr := &myCard5
	myCard6Ptr := &myCard6
	myCard7Ptr := &myCard7

	for _, cPtr := range clPtr {

		if cPtr != nil {
			c = *cPtr
			fmt.Println("in orderCardsOfMixedSuit, card from range: ", c)

			if myCard1.rank == rX { // 1st card in the suit of the Fl
				myCard1 = *cPtr

			} else if myCard2.rank == rX { // 2nd card in the suit of the Fl
				if crm[c.rank] > crm[myCard1.rank] { // if new card is higher than the the first
					myCard2 = myCard1 // move the first to the second position
					myCard1 = c       // Set the first/top pos to c
				} else {
					myCard2 = c
				}

			} else if myCard3.rank == rX { // 3rd card in the suit of the Fl
				if crm[c.rank] > crm[myCard1.rank] { // higher than the other 2
					myCard3 = myCard2
					myCard2 = myCard1
					myCard1 = c
				} else if crm[c.rank] > crm[myCard2.rank] {
					myCard3 = myCard2
					myCard2 = c
				} else {
					myCard3 = c
				}

			} else if myCard4.rank == rX {
				if crm[c.rank] > crm[myCard1.rank] { // higher than the other 3
					myCard4 = myCard3
					myCard3 = myCard2
					myCard2 = myCard1
					myCard1 = c
				} else if crm[c.rank] > crm[myCard2.rank] {
					myCard4 = myCard3
					myCard3 = myCard2
					myCard2 = c
				} else if crm[c.rank] > crm[myCard3.rank] {
					myCard4 = myCard3
					myCard3 = c
				} else {
					myCard4 = c
				}

			} else if myCard5.rank == rX {
				if crm[c.rank] > crm[myCard1.rank] { // higher than the other 4
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = myCard2
					myCard2 = myCard1
					myCard1 = c
				} else if crm[c.rank] > crm[myCard2.rank] {
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = myCard2
					myCard2 = c
				} else if crm[c.rank] > crm[myCard3.rank] {
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = c
				} else if crm[c.rank] > crm[myCard4.rank] {
					myCard5 = myCard4
					myCard4 = c
				} else {
					myCard5 = c
				}

			} else if myCard6.rank == rX {
				if crm[c.rank] > crm[myCard1.rank] { // higher than the other 5
					myCard6 = myCard5
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = myCard2
					myCard2 = myCard1
					myCard1 = c
				} else if crm[c.rank] > crm[myCard2.rank] {
					myCard6 = myCard5
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = myCard2
					myCard2 = c
				} else if crm[c.rank] > crm[myCard3.rank] {
					myCard6 = myCard5
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = c
				} else if crm[c.rank] > crm[myCard4.rank] {
					myCard6 = myCard5
					myCard5 = myCard4
					myCard4 = c
				} else if crm[c.rank] > crm[myCard5.rank] {
					myCard6 = myCard5
					myCard5 = c
				} else {
					myCard6 = c
				}

			} else if myCard7.rank == rX {
				if crm[c.rank] > crm[myCard1.rank] { // higher than the other 6
					myCard7 = myCard6
					myCard6 = myCard5
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = myCard2
					myCard2 = myCard1
					myCard1 = c
				} else if crm[c.rank] > crm[myCard2.rank] {
					myCard7 = myCard6
					myCard6 = myCard5
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = myCard2
					myCard2 = c
				} else if crm[c.rank] > crm[myCard3.rank] {
					myCard7 = myCard6
					myCard6 = myCard5
					myCard5 = myCard4
					myCard4 = myCard3
					myCard3 = c
				} else if crm[c.rank] > crm[myCard4.rank] {
					myCard7 = myCard6
					myCard6 = myCard5
					myCard5 = myCard4
					myCard4 = c
				} else if crm[c.rank] > crm[myCard5.rank] {
					myCard7 = myCard6
					myCard6 = myCard5
					myCard5 = c
				} else if crm[c.rank] > crm[myCard6.rank] {
					myCard7 = myCard6
					myCard6 = c
				} else {
					myCard7 = c
				}
			}
		}
	}

	resultingClPtr[0] = myCard1Ptr
	resultingClPtr[1] = myCard2Ptr
	resultingClPtr[2] = myCard3Ptr
	resultingClPtr[3] = myCard4Ptr
	resultingClPtr[4] = myCard5Ptr

	switch {
	case myCard6.rank != rX:
		resultingClPtr[5] = myCard6Ptr
	case myCard7.rank != rX:
		resultingClPtr[6] = myCard7Ptr
	}

	fmt.Println("dumping the resultingClPtr: ", resultingClPtr[0], resultingClPtr[1], resultingClPtr[2], resultingClPtr[3], resultingClPtr[4], resultingClPtr[5], resultingClPtr[6])

	return resultingClPtr, err
}

// #####################################################################
// find5OrMoreOfSameSuitInCardList takes a card list and returns true, if there were
// 5 or more cards of given suit in it, and the list of cards of that suit.
// Otherwise returns false
func find5OrMoreOfSameSuitInCardList(clPtr *cardList, cll int, scPtr *suitCounter) (resultingClPtr *cardList, topSuit deck.CardSuit, err error) {
	fmt.Println("### Starting find5OrMoreOfSameSuitInCardList ###")
	// fmt.Println("in find5OrMoreOfSameSuitInCardList; clPtr[...]: ", clPtr[0], clPtr[1])

	resultingClPtr = createClPtr()

	// Run prelim checks
	if scPtr.max < 5 || cll < 5 {
		err = errors.New("scPtr.max < 5 OR cll < 5")
		return resultingClPtr, x, err
	}

	// Figure out which suit is most common in the list.
	topSuit = x
	for _, cs := range suitList {
		// fmt.Println("in find5OrMoreOfSameSuitInCardList; cs: ", cs)
		if scPtr.scm[cs] > 4 {
			topSuit = cs
		}

	}

	// Grab just the cards of the topSuit
	i := 0
	for _, cPtr := range clPtr {
		if cPtr.suit == topSuit {
			resultingClPtr[i] = cPtr
			i++
		}
	}

	return resultingClPtr, topSuit, err
}

// #####################################################################
// findSt looks for straights - doc line
func findSt(clPtr *cardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {
	fmt.Println("### Looking for a St ###")

	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)
	fmt.Println("scPtr.max: ", scPtr.max)

	// default state
	*handNamePtr = "did not find a St"
	err = errors.New("Did not find a St")
	// fmt.Println("did not find a St, err: ", err)

	// First check, the basics
	// Since St is getting to ge rather far in the hand classification, some of the prelim checks are becoming "absolete", ie. certain configurations
	// should never bee seen this far dowin in the hand rank check list.
	if cll < 5 || rcPtr.uniqeRankCt < 5 || scPtr.max > 4 {
		err = errors.New("Failed prelim St checks, exiting findSt")
		fmt.Println("prelim check err: ", err)
		return err
	}

L02:
	for i := 0; i <= 9; i++ {
		// fmt.Println("i: ", i)
		// fmt.Printf("in findSt rankListFull of i is: %v\n\n", rankListFull[i])
		// fmt.Printf("clPtr 0 : %v; 1: %v; 2: %v; 3: %v; 4: %v; 5: %v; 6: %v\n", clPtr[0].rank, clPtr[1], clPtr[2], clPtr[3], clPtr[4], clPtr[5], clPtr[6])
		// fmt.Printf("rankLFi : %v;+1: %v;+2: %v;+3: %v;+4: %v\n\n", rankListFull[i], rankListFull[i+1], rankListFull[i+2], rankListFull[i+3], rankListFull[i+4])

		if (clPtr[0].rank == rankListFull[i] || clPtr[1].rank == rankListFull[i] || clPtr[2].rank == rankListFull[i] || clPtr[3].rank == rankListFull[i] || clPtr[4].rank == rankListFull[i] || clPtr[5].rank == rankListFull[i] || clPtr[6].rank == rankListFull[i]) &&
			(clPtr[0].rank == rankListFull[i+1] || clPtr[1].rank == rankListFull[i+1] || clPtr[2].rank == rankListFull[i+1] || clPtr[3].rank == rankListFull[i+1] || clPtr[4].rank == rankListFull[i+1] || clPtr[5].rank == rankListFull[i+1] || clPtr[6].rank == rankListFull[i+1]) &&
			(clPtr[0].rank == rankListFull[i+2] || clPtr[1].rank == rankListFull[i+2] || clPtr[2].rank == rankListFull[i+2] || clPtr[3].rank == rankListFull[i+2] || clPtr[4].rank == rankListFull[i+2] || clPtr[5].rank == rankListFull[i+2] || clPtr[6].rank == rankListFull[i+2]) &&
			(clPtr[0].rank == rankListFull[i+3] || clPtr[1].rank == rankListFull[i+3] || clPtr[2].rank == rankListFull[i+3] || clPtr[3].rank == rankListFull[i+3] || clPtr[4].rank == rankListFull[i+3] || clPtr[5].rank == rankListFull[i+3] || clPtr[6].rank == rankListFull[i+3]) &&
			(clPtr[0].rank == rankListFull[i+4] || clPtr[1].rank == rankListFull[i+4] || clPtr[2].rank == rankListFull[i+4] || clPtr[3].rank == rankListFull[i+4] || clPtr[4].rank == rankListFull[i+4] || clPtr[5].rank == rankListFull[i+4] || clPtr[6].rank == rankListFull[i+4]) {

			fmt.Printf("in findSt inside IF \n%v\n%v\n%v\n%v\n%v\n", rankListFull[i], rankListFull[i+1], rankListFull[i+2], rankListFull[i+3], rankListFull[i+4])

			*handNamePtr = string(rankListFull[i]) + " high Straight"
			err = nil // since it could have been set to "not nil"
			break L02
		}
	}
	return err
}

// #####################################################################
// find2x2 looks for two pair hands - doc line
func find2x2(clPtr *cardList, cll int, rcPtr *rankCounter, handNamePtr *string) (err error) {
	fmt.Println("### Looking for a 2x2 ###")

	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)

	orderedClPtr, err := orderCardsOfMixedSuit(clPtr)
	err04 := findPairs(orderedClPtr, cll, rcPtr) // Assigns the 2x var values in the rcPtr

	if err04 == nil && rcPtr.top2x2 != rX {
		*handNamePtr = "Two pair, " + string(rcPtr.top2x1) + "s and " + string(rcPtr.top2x2) + "s, " + string(rcPtr.top1x1) + " kicker"
	} else {
		err = errors.New("did not find a 2 pair hand")
		fmt.Println("did not find 2x2; err: ", err)
	}

	return err
}

// #####################################################################
// findPairs examines the card list and if it finds pairs, assigns appropriate ranks to rcCouter.topXXXX vars.
func findPairs(clPtr *cardList, cll int, rcPtr *rankCounter) (err error) {
	fmt.Println("### in findPairs ###")
	fmt.Println("### in findPairs ###")
	fmt.Println("### in findPairs ###")

	fmt.Println("cll-2: ", cll-2)
	fmt.Println("dump rcPtr: ", rcPtr)

	for i := 0; i < cll-2; i++ {
		fmt.Println("i: ", i)
		fmt.Println("clPtr[i]: ", clPtr[i], "clPtr[i+1]: ", clPtr[i+1], "rcPtr.top2x1: ", rcPtr.top2x1)

		if clPtr[i].rank == clPtr[i+1].rank {
			if rcPtr.top2x1 == rX {
				fmt.Println("in findPairs, inside IF, found a pair")
				rcPtr.top2x1 = clPtr[i].rank
			} else if rcPtr.top2x2 == rX {
				fmt.Println("in findPairs, inside IF, found 2nd pair")
				rcPtr.top2x2 = clPtr[i].rank
			}
			i = i + 1 // if we encountered the first card in a pair, then skip the second
		} else {
			switch {
			case rcPtr.top1x1 == rX:
				rcPtr.top1x1 = clPtr[i].rank
			case rcPtr.top1x2 == rX:
				rcPtr.top1x2 = clPtr[i].rank
			case rcPtr.top1x3 == rX:
				rcPtr.top1x3 = clPtr[i].rank
			case rcPtr.top1x4 == rX:
				rcPtr.top1x4 = clPtr[i].rank
			case rcPtr.top1x5 == rX:
				rcPtr.top1x5 = clPtr[i].rank
			default:
				fmt.Println("reached default case in findPairs; i: ", i)
			}
		}
	}

	if rcPtr.top2x1 == rX {
		err = errors.New("did not find any pairs")
		fmt.Println("did not find any pairs; err: ", err)
	}

	fmt.Println("rcPtr.top2x1: ", rcPtr.top2x1, "rcPtr.top2x2: ", rcPtr.top2x2)

	return err
}

// #####################################################################
// find2x1 looks for a single pair
// So, find2x2 ran the findPairs, which set up the rcPtr to the end. We can just use that info here
// and do not need to figure anything out.
func find2x1(clPtr *cardList, cll int, rcPtr *rankCounter, handNamePtr *string) (err error) {
	fmt.Println("### Looking for a 2x1 ###")

	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)

	// orderedClPtr, err := orderCardsOfMixedSuit(clPtr)
	// err05 := findPairs(orderedClPtr, cll, rcPtr) // Assigns the 2x var values in the rcPtr

	if rcPtr.top2x1 != rX {
		*handNamePtr = "Pair, " + string(rcPtr.top2x1) + "s, " + string(rcPtr.top1x1) + string(rcPtr.top1x2) + string(rcPtr.top1x3) + " kicker."
	} else {
		err = errors.New("did not find any pairs")
		fmt.Println("did not find 2x1; err: ", err)
	}

	return err
}

// #####################################################################
// findHc looks for a high card hand.
// Well, at this stage any other hand should have been caught by the pervious findXXX functions.
// So, find2x2 ran the findPairs, which set up the rcPtr to the end. We can just use that info here
// and do not need to figure anything out.
func findHc(clPtr *cardList, cll int, rcPtr *rankCounter, handNamePtr *string) (err error) {
	fmt.Println("### Looking for a High Card hand ###")

	fmt.Println("card list len, cll: ", cll)
	fmt.Println("rcPtrmax: ", rcPtr.max)

	if rcPtr.top2x1 == rX {
		*handNamePtr = "High card: " + string(rcPtr.top1x1) + ", " + string(rcPtr.top1x2) + string(rcPtr.top1x3) + string(rcPtr.top1x4) + string(rcPtr.top1x5) + " kicker."
	} else {
		err = errors.New("did not find any pairs")
		fmt.Println("did not find 2x1; err: ", err)
	}

	return err
}

// #####################################################################
// create5CardHandRankings gives us hand type lists, that we can compare our hands against, and decide which is better.
//

func create5CardHandRankings() (err error) {

	var rankingSf fiveCardHandOrderedList // Will be the sme as St
	var ranking4x fiveCardHandOrderedList
	var rankingFh fiveCardHandOrderedList
	var rankingFl fiveCardHandOrderedList
	var rankingSt fiveCardHandOrderedList
	var ranking3x fiveCardHandOrderedList
	var ranking2x2 fiveCardHandOrderedList
	var ranking2x fiveCardHandOrderedList
	var rankingHc fiveCardHandOrderedList

	rankingSf.handType = sf
	ranking4x.handType = x4
	rankingFh.handType = fh
	rankingFl.handType = fl
	rankingSt.handType = st
	ranking3x.handType = x3
	ranking2x2.handType = x22
	ranking2x.handType = x2
	rankingHc.handType = hc

	// ##########################
	// Create the St/Sf ranking
	for i := 0; i <= 9; i++ {
		rankingSf.handList[i].c1r = rankListFull[i]
		rankingSf.handList[i].c2r = rankListFull[i+1]
		rankingSf.handList[i].c3r = rankListFull[i+2]
		rankingSf.handList[i].c4r = rankListFull[i+3]
		rankingSf.handList[i].c5r = rankListFull[i+4]
	}

	/*
		fmt.Println("Sf/St ranking, card list: \n")
		for i := 0; i <= 10; i++ {
			fmt.Println("i: ", i, " ", rankingSt.handList[i])
		}
		fmt.Println()
	*/

	// ##########################
	// Create the 4x ranking
	// We don't care if there are "gaps" in the list, due to the kicker being the same as the quads, during the
	// internal loop iteration. As long as the better hand has a lowe index, all is good with the world.
	//
	// With quads, the kicker does not matter, as two people can't have the same rank queads.
	// Having a "complete" rank list does not hurt us, so unless performance becomes an issue,
	// leaving the ranking as is, with kickers.

	for i := 0; i <= 12; i++ { // 13 ranks * 12 kickers each; the loop for the 4x
		for j := 0; j <= 12; j++ { // Assign all possible kickers
			ranking4x.handList[i*13+j].c1r = rankList[i]
			ranking4x.handList[i*13+j].c2r = rankList[i]
			ranking4x.handList[i*13+j].c3r = rankList[i]
			ranking4x.handList[i*13+j].c4r = rankList[i]
			if i == j {
				ranking4x.handList[i*13+j].c5r = rX
			} else {
				ranking4x.handList[i*13+j].c5r = rankList[j]
			}
		}
	}

	/*
		fmt.Println("4x ranking, card list: \n")
		for i := 0; i <= 10; i++ {
			fmt.Println("i: ", i, " ", ranking4x.handList[i])
		}
	*/

	// ##########################
	// Create the Fh ranking

	for i := 0; i <= 12; i++ { // 13 ranks for the 3x
		for j := 0; j <= 12; j++ { // All possible 2x ranks
			rankingFh.handList[i*13+j].c1r = rankList[i]
			rankingFh.handList[i*13+j].c2r = rankList[i]
			rankingFh.handList[i*13+j].c3r = rankList[i]

			if i == j {
				rankingFh.handList[i*13+j].c4r = rX
				rankingFh.handList[i*13+j].c5r = rX
			} else {
				rankingFh.handList[i*13+j].c4r = rankList[j]
				rankingFh.handList[i*13+j].c5r = rankList[j]
			}
		}
	}

	/*
		fmt.Println("Fh ranking, card list: \n")
		for i := 0; i <= 10; i++ {
			fmt.Println("i: ", i, " ", rankingFh.handList[i])
		}
		fmt.Println()

	*/

	// ##########################
	// Create the Fl ranking
	// Here we got to got to work the list out out for 10 major categories of flushes, based on the
	// high card, so, A-T, down to 5-A.
	// The Sfs are going to be included in the list, BUT, the card list we'll be comparing with them
	// are NOT expected to coatain any Sfs, so the ranking method should give correct results.

	index := 0

	for i := 0; i <= 8; i++ { // 1st
		for j := i + 1; j <= 9; j++ { // 2nd
			for k := j + 1; k <= 10; k++ { // 3rd
				for l := k + 1; l <= 11; l++ { // 4th
					for m := l + 1; m <= 12; m++ { // 5th/last

						rankingFl.handList[index].c1r = rankList[i]
						rankingFl.handList[index].c2r = rankList[j]
						rankingFl.handList[index].c3r = rankList[k]
						rankingFl.handList[index].c4r = rankList[l]
						rankingFl.handList[index].c5r = rankList[m]

						/*
							if index > 50 {
								for a := index - 50; a <= index; a++ {
									fmt.Println("*** rank: ", a, "      ", "l: ", l, " m: ", m, " index: ", index, "   ", rankingFl.handList[a])
								}
								fmt.Println("######")
							}
						*/
						index++
					}
				}
			}
		}
	}

	/*
		fmt.Println("Fl ranking, card list, TOP: \n")
		for i := 0; i <= 20; i++ {
			fmt.Println("i: ", i, " ", rankingFl.handList[i])
		}
		fmt.Println()

		fmt.Println("Fl ranking, card list, BOTTOM: \n")
		for i := index - 20; i <= index; i++ {
			fmt.Println("i: ", i, " ", rankingFl.handList[i])
		}
		fmt.Println()
	*/

	// ##########################
	// Create the St ranking

	rankingSt.handList = rankingSf.handList

	// ##########################
	// Create the 3x ranking
	// Here too, the assumption is that we're working on an ordered list of cards, by rank, A to 2.

	index = 0

	for i := 0; i <= 12; i++ { // 13 ranks * 12 kickers each; the loop for the 4x

		ranking3x.handList[index].c1r = rankList[i]
		ranking3x.handList[index].c2r = rankList[i]
		ranking3x.handList[index].c3r = rankList[i]

		ranking3x.handList[index].c4r = rX // mark the kicker slots as "not yet filled / unused"
		ranking3x.handList[index].c5r = rX

		for j := 0; j <= 12; j++ { // Assign the top kicker

			if rankList[j] != rankList[i] && ranking3x.handList[index].c4r == rX {
				ranking3x.handList[index].c4r = rankList[j]

				for k := j + 1; k <= 12; k++ { // Assign the second kicker, lower than the top

					if rankList[k] != rankList[i] && ranking3x.handList[index].c5r == rX {
						ranking3x.handList[index].c5r = rankList[k]

						index++

						ranking3x.handList[index].c1r = rankList[i]
						ranking3x.handList[index].c2r = rankList[i]
						ranking3x.handList[index].c3r = rankList[i]

						ranking3x.handList[index].c4r = rankList[j]
						ranking3x.handList[index].c5r = rX

					}
				}
			}
		}
	}

	/*
		fmt.Println("3x ranking, card list, TOP: \n")
		for i := 0; i <= 50; i++ {
			fmt.Println("i: ", i, " ", ranking3x.handList[i])
		}
		fmt.Println()

		fmt.Println("3x ranking, card list, BOTTOM: \n")
		for i := index - 30; i <= index; i++ {
			fmt.Println("i: ", i, " ", ranking3x.handList[i])
		}
		fmt.Println()
	*/

	// ##########################
	// Create the 2x2 ranking
	// Here too, the assumption is that we're working on an ordered list of cards, by rank, A to 2.
	// So, the 2nd pair will always be lower in rank than the first.

	index = 0

	for i := 0; i <= 11; i++ { // going over 12 ranks, since last iteration will be 3322x, i corresponding to r3

		ranking2x2.handList[index].c1r = rankList[i]
		ranking2x2.handList[index].c2r = rankList[i]

		ranking2x2.handList[index].c5r = rX

		for j := i + 1; j <= 12; j++ { // Assign the second pair

			ranking2x2.handList[index].c3r = rankList[j]
			ranking2x2.handList[index].c4r = rankList[j]

			for k := 0; k <= 12; k++ { // Assign the kicker

				if rankList[k] != rankList[i] && rankList[k] != rankList[j] && ranking2x2.handList[index].c5r == rX {
					ranking2x2.handList[index].c5r = rankList[k]

					index++

					ranking2x2.handList[index].c1r = rankList[i]
					ranking2x2.handList[index].c2r = rankList[i]

					ranking2x2.handList[index].c3r = rankList[j]
					ranking2x2.handList[index].c4r = rankList[j]

					ranking2x2.handList[index].c5r = rX
				}
			}
		}
	}
	/*
		fmt.Println("2x2 ranking, card list, TOP: \n")
		for i := 0; i <= 50; i++ {
			fmt.Println("i: ", i, " ", ranking2x2.handList[i])
		}
		fmt.Println()

		fmt.Println("2x2 ranking, card list, BOTTOM: \n")
		for i := index - 30; i <= index; i++ {
			fmt.Println("i: ", i, " ", ranking2x2.handList[i])
		}
		fmt.Println()
	*/

	// ##########################
	// Create the 2x ranking
	// Here too, the assumption is that we're working on an ordered list of cards, by rank, A to 2.

	index = 0

	for i := 0; i <= 12; i++ { //
		// fmt.Println("entering iiii loop ***")

		ranking2x.handList[index].c1r = rankList[i]
		ranking2x.handList[index].c2r = rankList[i]

		ranking2x.handList[index].c3r = rX
		ranking2x.handList[index].c4r = rX // mark the kicker slots as "not yet filled / unused"
		ranking2x.handList[index].c5r = rX

	LC3:
		for j := 0; j <= 10; j++ { // Assign the 1st/top kicker
			// fmt.Println("entering jjjj loop ***")

			if rankList[j] != rankList[i] && ranking2x.handList[index].c3r == rX {
				ranking2x.handList[index].c3r = rankList[j]

			LC4:
				for k := j + 1; k <= 11; k++ { // Assign the 2nd kicker, lower than the top
					// fmt.Println("entering kkkk loop ***")

					if rankList[k] != rankList[i] && ranking2x.handList[index].c4r == rX {
						ranking2x.handList[index].c4r = rankList[k]

						for l := k + 1; l <= 12; l++ { // Assign the 3rd kicker, lower than 2nd
							// fmt.Println("entering llll loop ***")

							// Check for the "i:  2694   {2 2 A 4 3} / i:  2695   {2 2 A unknown unknown}" scenario, and move us forward.
							if i == 12 && k == 10 && l == 12 {
								ranking2x.handList[index].c3r = rX
								ranking2x.handList[index].c4r = rX
								continue LC3
							}

							// Check for the 	"i:  2649   {2 2 A K 3} /  i:  2650   {2 2 A K unknown}" scenario, and move us forward.
							if i == 12 && l == 12 {
								ranking2x.handList[index].c4r = rX
								continue LC4
							}

							if (rankList[l] != rankList[i]) && (rankList[l] != rankList[j]) && (rankList[l] != rankList[k]) && (ranking2x.handList[index].c5r == rX) {
								ranking2x.handList[index].c5r = rankList[l]

								index++

								ranking2x.handList[index].c1r = rankList[i]
								ranking2x.handList[index].c2r = rankList[i]

								// If c4r is at r3, and therefore c5r is at r2, then c3r needs to get bumped at next itereation,
								// we need to commiunicate this to the 2nd outer loop.
								if rankList[k] == r3 || (rankList[k] == r4 && rankList[i] == r3) {
									ranking2x.handList[index].c3r = rX
								} else {
									ranking2x.handList[index].c3r = rankList[j]
								}

								// If c5r is at r2, then c4r needs to get bumped at next itereation,
								// we need to commiunicate this to the outer loop

								if rankList[l] == r2 {
									ranking2x.handList[index].c4r = rX
								} else {
									ranking2x.handList[index].c4r = rankList[k]
								}

								ranking2x.handList[index].c5r = rX

								/*
									fmt.Println("i, j, k, l: ", i, j, k, l)
									fmt.Println("index - 1 : ", index-1, " ", ranking2x.handList[index-1])
									fmt.Println("index:      ", index, " ", ranking2x.handList[index])
									fmt.Println("index + 1 : ", index+1, " ", ranking2x.handList[index+1])
									fmt.Println()
								*/
							}
						}
					}
				}
			}
		}
	}
	/*
		fmt.Println("2x ranking, card list, TOP: \n")
		for i := 0; i <= 20; i++ {
			fmt.Println("i: ", i, " ", ranking2x.handList[i])
		}
		fmt.Println()

		fmt.Println("2x ranking, card list, BOTTOM: \n")
		for i := 2000; i <= index; i++ {
			fmt.Println("i: ", i, " ", ranking2x.handList[i])
		}
		fmt.Println()
	*/

	// ##########################
	// Create the Hc ranking

	rankingHc.handList = rankingFl.handList

	return err
}

// #####################################################################
// createTable constructs a new poker table (as in a kitchen table, not a tble of numbers), with a specific # of seats, 2-10.
func createTable(numSeats, buttonPos, tableNumber int, gt gameType, gs gameSize) (tablePtr *pokerTable, err error) {

	var table pokerTable
	tablePtr = &table

	tablePtr.numberSeats = numSeats
	tablePtr.buttonPos = buttonPos
	tablePtr.tableNumber = tableNumber
	tablePtr.kind = gt
	tablePtr.size = gs

	return tablePtr, err
}

/*
XXXXXXXXXXXXXX
NEXT:

Assign some combos to ranges, for testing

Develop functions to assign pairs, suiteds, and unsuiteds, just like we use in the combo matrix, to the tccl/range.

Develop a "map" method for manipulating the inRange seeting for tcc. We will need this for reading in, and reading out,
specific player ranges.
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

	ffl, _ := hand_ranking.CreateFiveCardHandKindRankings()
	// fmt.Println(ffl.sfInfo, ffl.x4Info, fhInfo, flInfo, stInfo, x3Info, x22Info, x2Info, hcInfo)
	fmt.Println(ffl.sfInfo)

	// Create a map for rank value look up, since constant structs are not supported, so we could
	// not have a rank struct with indexes, as a constant.
	crm = createCardRankMap()

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

	heroRangePtr, _ := createRange()
	villainRangePtr, _ := createRange()

	// Do some card removal manipulations
	// Remember to remove card from ALL ranges, when "running allSeeCard"
	allSeeCardInDeck(r2, s) // Use for Community cards, Flop/Turn/River
	// Remove combos based on card
	numRemovedFromHerosRange, _ := removeCardFromRange(r2, s, heroRangePtr)
	numRemovedFromVillainsRange, _ := removeCardFromRange(r2, s, villainRangePtr)

	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
	// Allways remove the card from All the users ranges
	// If you do do this, there should be no need for "reconciling" between cards seen and combos removed in ranges.

	for index, tcc := range heroRangePtr {
		// for index := 0; index < 100; index++ {
		// fmt.Printf("index: %v	element: %v\n", index, tccl[index])
		fmt.Printf("index: %v	hc: %v  lc: %v    removed: %v\n", index, tcc.c1Ptr, tcc.c2Ptr, tcc.removed)
	}

	fmt.Println()
	fmt.Println("removed from heros    r: ", numRemovedFromHerosRange)
	fmt.Println("removed from villains r: ", numRemovedFromVillainsRange)

	var communityCards board
	communityCards.f = dealFlop()
	communityCards.t = dealOne()
	communityCards.r = dealOne()

	maketesthands.RunCheckForHands()

	create5CardHandRankings()

	var gs gameSize
	gs.name = nl2
	gs.blind1 = 1
	gs.blind2 = 2

	tPtr, _ := createTable(9, 1, 1, nlh, gs)
	fmt.Printf("tPtr : %v\n", tPtr)

}

// EOF
