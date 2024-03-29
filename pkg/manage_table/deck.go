package manage_table

import (
	"github.com/ttudrej/pokertrainer/pkg/debugging"
)

type CdmKey struct {
	Cr CardRank
	Cs CardSuit
}

type CardDeckMap map[CdmKey]*Card

// orderedListOfPtrsToCard
type listOfPtrsToCards [52]*Card

// orderedListOfPtrsToCard uses 56 not 52 slots, to accomodate for the Aces in 5-A straights
// Used for hand rank checks ONLY
type listFullOfPtrsToCards [56]*Card

// type CardDeck struct {
// 	CdmPtr                     *CardDeckMap
// 	OloptcPtr                  *listOfPtrsToCards     // Ordered List of *cards
// 	TopCardIndex_oloptc        int                    // 0-51, For keeping track of the top card in the deck, as they get dealt
// 	ShuffledLoptcPtr           *listOfPtrsToCards     // For storing the shuffled version. The shuffling job belongs to the dealer
// 	TopCardIndex_shuffledLoptc int                    // 0-51, For keeping track of the top card as they get dealt out.
// 	OloptcFPtr                 *listFullOfPtrsToCards // Ordered List, Full, of *cards
// }

type CardDeck struct {
	CdmPtr                     *CardDeckMap
	OloptcPtr                  *listOfPtrsToCards     // Ordered List of *cards
	TopCardIndex_oloptc        int                    // 0-51, For keeping track of the top card in the deck, as they get dealt
	ShuffledLoptcPtr           *listOfPtrsToCards     // For storing the shuffled version. The shuffling job belongs to the dealer
	TopCardIndex_shuffledLoptc int                    // 0-51, For keeping track of the top card as they get dealt out.
	OloptcFPtr                 *listFullOfPtrsToCards // Ordered List, Full, of *cards
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
/* CreateDeck makes us a brand new deck, and also gives an ordered list of the cards in it.
The new deck means that we now have a one more set of virtual cards, completely separte
from any other deck already in place. This way no mixing of cards between decks can occur.

Here is where we actually Create cards. It makes sense that we Create whole decks of cards,
and not cards individually.
*/
// func CreateDeck() (CdmPtr *CardDeckMap, olPtr *listOfPtrsToCards, olfPtr *listFullOfPtrsToCards, err error) {
// func CreateDeck() (cdPtr *CardDeck, err error) {
func CreateDeck() (cdPtr *CardDeck, err error) {
	// Info.Println(debugging.ThisFunc()) # This paniCs currently
	// Info.Println("### Starting CreateDeck ###")
	cdm := make(CardDeckMap)
	CdmPtr := &cdm

	// Since we're Createing a brand new, and ORDERED deck, our list will be "ordered"
	// ol = Ordered List
	var ol listOfPtrsToCards
	olPtr := &ol

	var shuffledL listOfPtrsToCards
	slPtr := &shuffledL

	// olf = Ordered List Full
	var olf listFullOfPtrsToCards
	olfPtr := &olf

	// collect all things realated to a deck of cards in one struct
	var cd CardDeck
	cdPtr = &cd

	cd.CdmPtr = CdmPtr
	cd.OloptcPtr = olPtr
	cd.ShuffledLoptcPtr = slPtr
	cd.OloptcFPtr = olfPtr
	cd.TopCardIndex_oloptc = 0
	cd.TopCardIndex_shuffledLoptc = 0

	sequence := 1 // bottom card in the deck, card 52 is the top of the deck.

	for _, rank := range RankList {
		for _, suit := range SuitList {
			var c = Card{rank, suit, false, false, false, false, sequence}
			cPtr := &c

			cdm[CdmKey{rank, suit}] = cPtr
			ol[sequence-1] = cPtr
			olf[sequence-1] = cPtr
			shuffledL[sequence-1] = cPtr
			// The dealer needs to perform a shuffle for this list to be actully shuffled.
			// It is ordered, when the deck is first Created.

			// Also point add Aces that fit below the duces.
			// Needed for working out Straigh relative ranking.
			if rank == RA {
				olf[sequence-1+52] = cPtr
			}
			sequence++
		}
	}
	// Return the card deck map, the Ordered List, and the Ordered List Full
	// return CdmPtr, olPtr, olfPtr, err
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
