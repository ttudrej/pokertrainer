package gameobjects

import (
	"github.com/ttudrej/pokertrainer/pkg/debugging"
	"github.com/ttudrej/pokertrainer/pkg/gameobjects"
)

type cdmKey struct {
	cr gameobjects.CardRank
	cs gameobjects.CardSuit
}

type cardDeckMap map[cdmKey]*gameobjects.Card

// orderedListOfPtrsToCard
type listOfPtrsToCards [52]*gameobjects.Card

// orderedListOfPtrsToCard uses 56 not 52 slots, to accomodate for the Aces in 5-A straights
// Used for hand rank checks ONLY
type listFullOfPtrsToCards [56]*gameobjects.Card

type cardDeck struct {
	cdmPtr                     *cardDeckMap
	oloptcPtr                  *listOfPtrsToCards     // Ordered List of *cards
	topCardIndex_oloptc        int                    // 0-51, For keeping track of the top card in the deck, as they get dealt
	shuffledLoptcPtr           *listOfPtrsToCards     // For storing the shuffled version. The shuffling job belongs to the dealer
	topCardIndex_shuffledLoptc int                    // 0-51, For keeping track of the top card as they get dealt out.
	oloptcFPtr                 *listFullOfPtrsToCards // Ordered List, Full, of *cards
}

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

// #########################################################################################
/* createDeck makes us a brand new deck, and also gives an ordered list of the cards in it.
The new deck means that we now have a one more set of virtual cards, completely separte
from any other deck already in place. This way no mixing of cards between decks can occur.

Here is where we actually create cards. It makes sense that we create whole decks of cards,
and not cards individually.
*/
// func createDeck() (cdmPtr *cardDeckMap, olPtr *listOfPtrsToCards, olfPtr *listFullOfPtrsToCards, err error) {
func createDeck() (cdPtr *cardDeck, err error) {
	Info.Println(debugging.ThisFunc())
	// Info.Println("### Starting createDeck ###")
	cdm := make(cardDeckMap)
	cdmPtr := &cdm

	// Since we're createing a brand new, and ORDERED deck, our list will be "ordered"
	// ol = Ordered List
	var ol listOfPtrsToCards
	olPtr := &ol

	var shuffledL listOfPtrsToCards
	slPtr := &shuffledL

	// olf = Ordered List Full
	var olf listFullOfPtrsToCards
	olfPtr := &olf

	// collect all things realated to a deck of cards in one struct
	var cd cardDeck
	cdPtr = &cd

	cd.cdmPtr = cdmPtr
	cd.oloptcPtr = olPtr
	cd.shuffledLoptcPtr = slPtr
	cd.oloptcFPtr = olfPtr
	cd.topCardIndex_oloptc = 0
	cd.topCardIndex_shuffledLoptc = 0

	sequence := 1 // bottom card in the deck, card 52 is the top of the deck.

	for _, rank := range gameobjects.RankList {
		for _, suit := range gameobjects.SuitList {
			var c = gameobjects.Card{rank, suit, false, false, false, false, sequence}
			cPtr := &c

			cdm[cdmKey{rank, suit}] = cPtr
			ol[sequence-1] = cPtr
			olf[sequence-1] = cPtr
			shuffledL[sequence-1] = cPtr
			// The dealer needs to perform a shuffle for this list to be actully shuffled.
			// It is ordered, when the deck is first created.

			// Also point add Aces that fit below the duces.
			// Needed for working out Straigh relative ranking.
			if rank == gameobjects.RA {
				olf[sequence-1+52] = cPtr
			}
			sequence++
		}
	}
	// Return the card deck map, the Ordered List, and the Ordered List Full
	// return cdmPtr, olPtr, olfPtr, err
	return cdPtr, err
}

// #########################################################################################
// print the cards in order they appear in the list of cards (currtnt order of the deck)
func displayCardsInList(clPtr *listOfPtrsToCards) (err error) {
	Info.Println(debugging.ThisFunc())
	cl := *clPtr

	for index := range cl {
		Info.Printf("index: %v, card: %v %v\n", index, cl[index].Rank, cl[index].Suit)
	}

	return err
}
