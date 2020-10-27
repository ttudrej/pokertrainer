/*
File: deck_shuffler_test.go, version history:
v0.1	yyyy/mm/dd	Tomasz Tudrej

*/

package main

import "testing"

// TestShuffleDeck_cardIndexMatching checks if we now have different card at specified index
// This test works most of the time. It will fail "wrongly", in case when the new index is
// rndomly picked to be the same as the old.
func TestShuffleDeck_cardIndexMatching(t *testing.T) {
	Info.Println(thisFunc())

	// Create a new deck of cards.
	// We will only need the ordered list of pointers for this test, so we discard the deck map and the "full" list.
	// _, olPtr, _, _ := createDeck()
	dPtr, _ := createDeck()

	// shuffledList, _ := shuffleDeck(olPtr)
	shuffledList, _ := shuffleDeck(dPtr.oloptcPtr)

	cases := []struct {
		inputProvided_index int
	}{
		{0},
		{51},
	}

	for _, c := range cases {
		// if shuffledList[c.inputProvided_index].sequence == olPtr[c.inputProvided_index].sequence {
		if shuffledList[c.inputProvided_index].sequence == dPtr.oloptcPtr[c.inputProvided_index].sequence {
			// Found the card with the same sequence # at the same index in the original and shuffled lists.
			t.Errorf("Card did not change index, initial index: %v", c.inputProvided_index)
		}
	}
}

// EOF
