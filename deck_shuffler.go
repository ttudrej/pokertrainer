package main

import (
	"math/rand"
	"time"
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

// Take a specific deck, and return a ponter to a new and shuffled lists of pointers to cards, for that same deck.
// func shuffleDeck(loptcPtr *listOfPtrsToCards) (shuffledLoptcPtr *listOfPtrsToCards, err error) {
func shuffleDeck(dPtr *cardDeck) (err error) {
	Info.Println(ThisFunc())

	rPtr := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Seems we need accuracy in time greater than 1sec, so that the seed differs among consequent shuffles,
	// otherwise we end up with decks shuffled exaclty the same at two or more tables.

	permutatedList := rPtr.Perm(len(*dPtr.oloptcPtr))
	Info.Println("")
	Info.Println("len: ", len(*dPtr.oloptcPtr))

	for i, randomizedOldIndex := range permutatedList {
		dPtr.shuffledLoptcPtr[i] = dPtr.oloptcPtr[randomizedOldIndex]
		// Info.Printf("new: %v, old: %v\n", i, randomizedOldIndex)
	}

	return err
}

/*
This works because rand.Perm() will return a random permutation of
the numbers 0 through N (not including N),
so when we call it we don’t get a shuffled array,
but we do receive a random list of indexes that we could access.
*/

// #########################################################################
// showShuffledDeck shows cards in the shuffled order.
func showShuffledDeck(tPtr *table) error {

	Info.Println(ThisFunc())

	for i, cPtr := range tPtr.deckPtr.shuffledLoptcPtr {
		Info.Println("card: ", i, cPtr.rank, cPtr.suit)
	}

	return nil
}

// #########################################################################
// showShuffledDeck shows cards in the shuffled order.
func showdDeck(tPtr *table) error {

	Info.Println(ThisFunc())
	Info.Println("Un-Shuffled deck: ")

	for i, cPtr := range tPtr.deckPtr.oloptcPtr {
		Info.Println("card: ", i, cPtr.rank, cPtr.suit)
	}
	return nil
}

// #########################################################################
// #########################################################################
