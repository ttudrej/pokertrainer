package main

import (
	"container/ring"
	"errors"

	"github.com/ttudrej/pokertrainer/debugging"
	"github.com/ttudrej/pokertrainer/tableitems"
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

// #####################################################################
func advanceButton(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())
	if tPtr.buttonPos == 0 {
		err = errors.New("Button at seat 0, have the dealer start the have first?")
		// err := Info.Errorf("user %q (id %d) not found", name, id)
	}
	if tPtr.buttonPos == 9 {
		tPtr.buttonPos = 1
	} else {
		tPtr.buttonPos++
	}

	return err
}

// #####################################################################
// pitchOneCard takes one card of the top of the shuffled deck, and
// "gives" it to the next seat(not player) in sequence, that is
// supposed to get a card on this round.
//
// During normal play, ther is never going to be to pitch just one car, always two,
// so it we should package the "pitch card" function to always pitch two... future work.

func pitchOneCardToAll(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())

	if tPtr.deckPtr.topCardIndex_shuffledLoptc >= 52 {
		err = errors.New("Error, We ran out of cards to deal")
	}

	for _, seatNum := range tPtr.seatNumbersToBeDealtIn {

		Info.Println("Dealing a card to seatNum: ", seatNum)
		Info.Println("tPtr.deckPtr.topCardIndex_shuffledLoptc: ", tPtr.deckPtr.topCardIndex_shuffledLoptc)

		if tPtr.seatList[seatNum-1].c2Ptr.Rank != tableitems.RX {
			err = errors.New("Error, trying to give player too many cards, already has 2")
		}

		if tPtr.seatList[seatNum-1].c1Ptr.Rank == tableitems.RX {
			tPtr.seatList[seatNum-1].c1Ptr = tPtr.deckPtr.shuffledLoptcPtr[tPtr.deckPtr.topCardIndex_shuffledLoptc]
			Info.Println("dealt c1: ", tPtr.seatList[seatNum-1].c1Ptr.Rank, tPtr.seatList[seatNum-1].c1Ptr.Suit)
		} else {
			tPtr.seatList[seatNum-1].c2Ptr = tPtr.deckPtr.shuffledLoptcPtr[tPtr.deckPtr.topCardIndex_shuffledLoptc]
			Info.Println("dealt c2: ", tPtr.seatList[seatNum-1].c2Ptr.Rank, tPtr.seatList[seatNum-1].c2Ptr.Suit)
		}

		// Advance the top card pointer
		tPtr.deckPtr.topCardIndex_shuffledLoptc++

		Info.Println("")
	}
	return err
}

// #####################################################################
// Initialize each seat and table, to make things ready for the next hand.
// Other functions rely on seeing rX where a card is supposed to be, to indicat
// that there is no card there yet.
func collectAllCards(tPtr *table) (err error) {
	Info.Println("in collectAllCards")

	for _, seatPtr := range tPtr.seatList {
		seatPtr.c1Ptr = tableitems.NoCardPtr
		seatPtr.c2Ptr = tableitems.NoCardPtr
	}

	for _, commCardPtr := range tPtr.communityCardsList {
		Info.Printf("%v", commCardPtr)
		commCardPtr = tableitems.NoCardPtr
	}
	return nil
}

// #####################################################################
func pitchCards() (err error) {
	Info.Println("in pitchCards")
	return err
}

// #####################################################################
func getNextSeatNum9Handed(sn int) (int, error) {
	Info.Println("in getNextSeatNum9Handed")
	nextSeatNum := 0

	if sn == 9 {
		nextSeatNum = 1 // list index 0
	} else {
		nextSeatNum = sn + 1
	}
	return nextSeatNum, nil
}

// #####################################################################
func getNextSeatNum10Handed(sn int) (int, error) {
	Info.Println("in getNextSeatNum10Hands")
	nextSeatNum := 0

	if sn == 10 {
		nextSeatNum = 1 // list index 0
	} else {
		nextSeatNum = sn + 1
	}
	return nextSeatNum, nil
}

// #####################################################################
// getListOfSeatsToPitchCardsTo produeces a list of seat numbers, in order, to which cards
// should be pitched, and stores that list in tablePtr.seatNumbersToBeDealtIn.
// In case of a full table, it will be 1-9, list[0] == 9, and list[8] == 9, if the button is in
// seat 1.
// For a full table with button in seat 2, it will be list[0] == 2, list[8] == 1, ...
// #####################################################################
func makeListOfSeatsToPitchCardsTo(tPtr *table) (err error) {
	Info.Println("in makeListOfSeatsToPitchCardsTo")

	if tPtr.buttonPos == 0 || tPtr.buttonPos > 9 {
		err = errors.New("Button not yet properly assigned")
		Info.Println("Error!: ", err)
	}

	seatNumWithButton := tPtr.buttonPos

	// Need to know which seat is going to be the first after the button, and
	// supposed to have cards pitched to it
	var firstSeatToExamine int

	// !! We'll need to dig up the rules for moving the button, in all possible situatinos that can arrise.
	// !! For now, we'll do simplified version of this, just to get going.

	// The assumption, for now, is that the dealer will always move the button in such a way, that there
	// is always a player which should be dealt to, right behind it.

	// Figure out which seat to look at first
	if seatNumWithButton == 9 {
		firstSeatToExamine = 1 // list index 0
	} else {
		firstSeatToExamine = seatNumWithButton + 1
	}

	// First element to be appended should be at index 0, so we want to end up with elements
	// indexed 0-8, or 0-(<8), in case all seats are not occupied, or should not be pitched
	// to, for any reason.
	seatNum := 0

	// Make sure we're starting with a 0 length slice.
	tPtr.seatNumbersToBeDealtIn = tPtr.seatNumbersToBeDealtIn[:0]

	Info.Println("in dealer_NLH, initialized tPtr.seatNumbersToBeDealtIn ", tPtr.seatNumbersToBeDealtIn)

	for i := firstSeatToExamine; i <= firstSeatToExamine+8; i++ {
		Info.Println("first seat to examine: ", firstSeatToExamine)
		Info.Println("i: ", i)

		if i == 9 {
			seatNum = 9
		} else {
			seatNum = i % 9
		}
		// this produces lists like the followng:
		// 1 2 3 4 5 6 7 8 9
		// 2 3 4 5 6 7 8 9 1
		// ...
		// 7 8 9 1 2 3 4 5 6
		Info.Println("seatNum: ", seatNum)

		// XXXXXXXXXXX
		// index out of range:

		Info.Println("i-1: ", i-1)

		Info.Println("tPtr.seatList[seatNum-1].occupied", tPtr.seatList[seatNum-1].occupied)
		Info.Println("tPtr.seatList[seatNum-1].sittingIn", tPtr.seatList[seatNum-1].sittingIn)

		if tPtr.seatList[seatNum-1].occupied == true && tPtr.seatList[seatNum-1].sittingIn == true {
			tPtr.seatNumbersToBeDealtIn = append(tPtr.seatNumbersToBeDealtIn, seatNum)
			Info.Println("in dealer_NLH, ASSEMBLED tPtr.seatNumbersToBeDealtIn ", tPtr.seatNumbersToBeDealtIn)
		} else {
			Info.Println("Continueing...")
			continue
		}
	}

	Info.Println("FINAL; in dealer_NLH, ASSEMBLED tPtr.seatNumbersToBeDealtIn ", tPtr.seatNumbersToBeDealtIn)

	return err
}

// #####################################################################
// Like makeListOfSeatsToPitchCardsTo, but in ring(circular list) form.
// The role of this "ring" will be to keep track of ONLY the seats that still
// have an active role in the hand.
// Every time a player folds, that seat will be removed from the ring.
// The hand ends when there is only one element(Ring) left in the ring(circular list).
// Meant to allow the dealer to keep track of the state of the current hand.
func makeRingOfActiveSeatPtrs(tPtr *table) error {
	Info.Println("in makeRingOfActiveSeatPtrs")

	// A "Ring" (upper case) is an element of a circular list, or "ring" (lower case).
	// => Ring : element of a ring (circular list)
	// => ring : a circular list, of 0 or more Rings
	//
	// Rings do not have a beginning or end;
	// A pointer to any ring element serves as reference to the entire ring.
	// Empty rings are represented as nil Ring pointers.
	// The zero value for a Ring is a one-element ring with a nil Value.

	Info.Println("tPtr.seatNumbersToBeDealtIn.Len(): ", len(tPtr.seatNumbersToBeDealtIn))

	// Create a new ring of size numSeats
	rPtr := ring.New(len(tPtr.seatNumbersToBeDealtIn))

	tPtr.remainingActiveSeatsRingPtr = rPtr

	// Get the length of the ring
	numRingElements := rPtr.Len()
	Info.Println("numRingElements: ", numRingElements)

	// Initialize the ring with some values, in this case, pointers to seats which are "active"
	for _, seatNum := range tPtr.seatNumbersToBeDealtIn {

		rPtr.Value = tPtr.seatList[seatNum-1]
		sPtr, ok := rPtr.Value.(*seat)
		if ok {
			// Info.Printf("rPtr.Val type: %T, value: %v \n", rPtr.Value, rPtr.Value)
			Info.Printf("sPtr number: %v; assigned to player id: %v\n", sPtr.number, sPtr.assignedToPID)
		}
		// Move ring pointer to next element in the ring.
		rPtr = rPtr.Next()
	}
	return nil
}

// #####################################################################
// drawForButton determines which active seat gets' the button at a new
// game.
// Returns the seat number for the button, and an error.
//
// Starting with a dummy seat picker.
func drawForButton(tPtr *table) (int, error) {
	Info.Println("in drawForButton")
	return tPtr.seatNumbersToBeDealtIn[0], nil
}

// #####################################################################
// Shuffles and distributes hole cards to all participating
func startHand(tPtr *table) (err error) {

	Info.Println(debugging.ThisFunc())
	Info.Println("#############################################")
	Info.Println("#############################################")
	Info.Println("#############################################")
	Info.Println("Starting a hand at table: ", tPtr.tableID)
	Info.Println()

	_ = shuffleDeck(tPtr.deckPtr)
	// _ = showShuffledDeck(tPtr)

	// Reset the top card pointer to top of freshly shuffled deck.
	tPtr.deckPtr.topCardIndex_shuffledLoptc = 0

	// Draw for button, if brand new table, OR, advance the button, otherwise.
	if tPtr.newTable == true {
		// We're gonna use the simplest method to start with, just place button in first seat on list
		// that's allowed to play
		// !!! Adopt for tables with players missing, ie. < 9.
		// tPtr.buttonPos = tPtr.seatNumbersToBeDealtIn[0] // index is 0-8.
		tPtr.buttonPos = 1
		tPtr.newTable = false
		// tPtr.buttonPos, _ = drawForButton()  XXXX what we want to get to.
	} else {
		// Advance the button on an existing table
		advanceButton(tPtr)
	}

	_ = updateBlindPtrs(tPtr)

	// Get a list of pointers to seats, to which we're gonna pitch cards
	_ = makeListOfSeatsToPitchCardsTo(tPtr)
	_ = makeRingOfActiveSeatPtrs(tPtr)

	// give out 2 cards to all allowed to play the next hand.
	_ = pitchOneCardToAll(tPtr)
	_ = pitchOneCardToAll(tPtr)

	// _  pitchCardsToSeats()
	// _ = advanceButton(tPtr)

	return err
}

// #####################################################################
// dealCommunityCard
func dealCommunityCard(tPtr *table, cardIndex int) (err error) {
	Info.Println("in dealCommunityCard")

	tPtr.communityCardsList[cardIndex] = tPtr.deckPtr.shuffledLoptcPtr[tPtr.deckPtr.topCardIndex_shuffledLoptc]
	tPtr.deckPtr.topCardIndex_shuffledLoptc++

	return err
}

// #####################################################################
func dealFlop(tPtr *table) (err error) {
	Info.Println("in dealFlop")

	_ = dealCommunityCard(tPtr, 0)
	_ = dealCommunityCard(tPtr, 1)
	_ = dealCommunityCard(tPtr, 2)

	return err
}

// #####################################################################
func dealTurn(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())

	_ = dealCommunityCard(tPtr, 3)
	return err
}

// #####################################################################
func dealRiver(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())

	_ = dealCommunityCard(tPtr, 4)
	return err
}

// #####################################################################
// updateBlindPtrs figures out which seats will take the SB and the BB in the next hand.
func updateBlindPtrs(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())

	// Doing the simplest method first. Assuming full table, always (wrong)
	// It will need to be refined !!!
	// Keep in mind we're calculate the seat index (0-8) from a real seat number (1-9)

	nextSBSeatIndex := -1
	nextBBSeatIndex := -1

	Info.Println()
	Info.Println("next sb index / next bb index / button Pos: ", nextSBSeatIndex, nextBBSeatIndex, tPtr.buttonPos)
	Info.Println()

	if tPtr.buttonPos >= 1 && tPtr.buttonPos <= 7 {
		nextSBSeatIndex = tPtr.buttonPos
		nextBBSeatIndex = tPtr.buttonPos + 1
	}

	if tPtr.buttonPos == 8 {
		nextSBSeatIndex = 8
		nextBBSeatIndex = 0
	}

	if tPtr.buttonPos == 9 {
		nextSBSeatIndex = 0
		nextBBSeatIndex = 1
	}

	Info.Println()
	Info.Println("next sb index / next bb index: ", nextSBSeatIndex, nextBBSeatIndex, tPtr.buttonPos)
	Info.Println()

	tPtr.seatPtrSB = tPtr.seatList[nextSBSeatIndex]
	tPtr.seatPtrBB = tPtr.seatList[nextBBSeatIndex]

	return nil
}

// #####################################################################
// Prompt at each player in turn, to perform their desired action, check/bet/raise/fold.
func postBlinds(tPtr *table) (err error) {
	Info.Println()

	// needs to be after "draw fro button" or later.
	// _ = updateBlindPtrs(tPtr)

	_ = postSB(tPtr)
	_ = postBB(tPtr)

	done := false
	// Flag to indicatte whether we need all but one folded, OR
	// whether all called.

	for !done {
		done = true
	}

	return err
}

// #####################################################################
func postSB(tPtr *table) error {
	Info.Println(debugging.ThisFunc())

	// Assuming that SB is at buttonPos + 1, in the seatNumbersToBeDealtIn list
	// Decrease seat stack by amount up to sbAmount
	// If stack pre blind post is <= to sbAmount, mark seat as "allIn"

	// we want the index of the SB, which is buttonPos -1 + 1, so just buttonPos.

	Info.Println("postSB: tPtr.buttonPos: ", tPtr.buttonPos)

	if tPtr.seatPtrSB.stackSize <= tPtr.sbAmount {
		tPtr.seatPtrSB.allIn = true
		tPtr.seatPtrSB.betAmount = tPtr.seatPtrSB.stackSize
		// tPtr.potSizeThisRoundOnly += tPtr.seatPtrSB.stackSize
		tPtr.seatPtrSB.stackSize = 0
	} else {
		tPtr.seatPtrSB.betAmount = tPtr.sbAmount
		tPtr.seatPtrSB.stackSize -= tPtr.sbAmount
		// tPtr.potSizeThisRoundOnly += tPtr.sbAmount
	}

	tPtr.seatPtrSB.amSB = true
	Info.Println("postSB, seat betAmt: ", tPtr.seatPtrSB.betAmount)

	return nil
}

// #####################################################################
func postBB(tPtr *table) error {
	Info.Println(debugging.ThisFunc())

	// Assuming that BB is at buttonPos + 2, in the seatNumbersToBeDealtIn list
	// Decrease seat stack by amount up to sbAmount
	// If stack pre blind post is <= to sbAmount, mark seat as "allIn"

	Info.Println("postBB: tPtr.buttonPos: ", tPtr.buttonPos)

	if tPtr.seatPtrBB.stackSize <= tPtr.bbAmount {
		tPtr.seatPtrBB.allIn = true
		tPtr.seatPtrBB.betAmount = tPtr.seatPtrBB.stackSize
		tPtr.seatPtrBB.stackSize = 0
		// tPtr.potSize += tPtr.seatPtrBB.stackSize
	} else {
		tPtr.seatPtrBB.betAmount = tPtr.bbAmount
		tPtr.seatPtrBB.stackSize -= tPtr.bbAmount
		// tPtr.potSize += tPtr.bbAmount
	}
	tPtr.seatPtrBB.amBB = true
	Info.Println("postBB, seat betAmt: ", tPtr.seatPtrBB.betAmount)

	return nil
}

// #####################################################################
// executeHand just runs through a mock hand start to finish
// Test
func executeHand(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())

	_ = prepTableForNextHand(tPtr)

	// Remove the previously created ring, ie. recoup memory space.
	tPtr.remainingActiveSeatsRingPtr = nil

	// updateTableStateStruct(tPtr)

	// NEW HAND ############################################
	_ = startHand(tPtr)
	// _ = showPokerRoomTables(prPtr)
	updateTableStateStruct(tPtr)

	// Post SB/BB
	_ = postBlinds(tPtr)
	// _ = showPokerRoomTables(prPtr)
	updateTableStateStruct(tPtr)

	_ = dealFlop(tPtr)
	updateTableStateStruct(tPtr)

	_ = dealTurn(tPtr)
	updateTableStateStruct(tPtr)

	_ = dealRiver(tPtr)
	updateTableStateStruct(tPtr)

	_ = pushPotsToWinners(tPtr)
	updateTableStateStruct(tPtr)

	return err
}

// #####################################################################

// conductBettingRound makes each player take their action, when it's their
// turn, until all the action closes on the current betting round.
func conductBettingRound(tPtr *table, br bettingRound) error {
	Info.Println(debugging.ThisFunc())

	// betSize := 2

	rPtr, _ := findFirstSeatToAct(tPtr, br)

	// Take note of which seat went first
	sPtr, _ := rPtr.Value.(*seat)
	firstToActSeatNumber := sPtr.number
	Info.Println("first seat to act: ", firstToActSeatNumber)

L01:
	for {
		Info.Printf("Ring lngth: type: %T, value: %v ", rPtr.Len(), rPtr.Len())

		switch {

		case rPtr.Len() == 0:
			// The hand is over, all folded and/or the winner(s) of the hand have been decided and the
			// pots distributed.
			// We shoiuld never be at ring length 0.
			// Once there is only 1 player left, that is the winner, and we move onto the next hand, avoiding
			// the trouble of dealing with the special case ring of lenght 0.
			break L01

		case rPtr.Len() == 1:
			// There si only one player left, and therefore the winner by default.
			tPtr.timeToPushPot = true
			_ = pushPotToSeat(tPtr, sPtr)

			break L01

		case rPtr.Len() > 1:
			// If there is at leat one more seat left in the hand, then player get's to decide
			// what to do. Otherwise, he's the last person in the hand, and all others have folded,
			// therefore he wins, and takes the pot.

			Info.Println("Len > 1: ", rPtr.Len())
			Info.Println("RING BEORE ACTION")

			// Iterate through the remaining ring and print its contents
			rPtr.Do(func(p interface{}) {
				sPtr = p.(*seat)
				Info.Printf("Ring lngth: %v ", rPtr.Len())
				Info.Println("iterting BEFORE ACTION, seat pid: ", sPtr.assignedToPID)
			})

			_ = performAction(sPtr, rPtr, tPtr)

			Info.Println("RING AFTER ACTION")
			// Iterate through the remaining ring and print its contents
			rPtr.Do(func(p interface{}) {
				sPtr = p.(*seat)
				Info.Printf("Ring lngth: value: %v ", rPtr.Len())
				Info.Println("iterting AFTER ACTION, seat pid: ", sPtr.assignedToPID)
			})

			// Move pointer at the next seat
			rPtr = rPtr.Next()
			sPtr, _ = rPtr.Value.(*seat)

			Info.Printf("Ring lngth: %v ", rPtr.Len())

			// XXX

			// Here we first need to have each seat just Fold, until there is only 1 seat, the winner left.

		}
	}

	/*

	   func (r *Ring) Prev() *Ring
	   Prev returns the previous ring element. r must not be empty.

	   func (r *Ring) Unlink(n int) *Ring
	   Unlink removes n % r.Len() elements from the ring r, starting at r.Next().
	   If n % r.Len() == 0, r remains unchanged. The result is the removed subring.
	   r must not be empty.

	   L 10 10 10
	   n 1  2  10
	   % 1  2  0
	*/

	/*
		for {
			Info.Println("Looping...")

			if sPtr.sittingIn && sPtr.stackSize > 200 {
				if sPtr.amSB {
					// We are only need BB amound - SB amount, to call
					sPtr.betAmount = betSize
					sPtr.stackSize -= 1
				} else if sPtr.amBB {
					// We are only need BB amound - SB amount, to call
					sPtr.betAmount = betSize
					sPtr.stackSize -= betSize - tPtr.bbAmount
				} else {
					sPtr.betAmount = betSize
					sPtr.stackSize -= betSize
				}
			} else {
				sPtr.folded = true
			}
			updateTableStateStruct(tPtr)

			// time.Sleep(2 * 100 * time.Millisecond)

			// Move poiter at the next seat
			rPtr = rPtr.Next()
			sPtr, _ = rPtr.Value.(*seat)

			if sPtr.number == firstToActSeatNumber {
				// we went all the way around
				Info.Println("Breaking, we went all the way around. firstToActSeatNumber: ", firstToActSeatNumber)
				break
			}
		}
		// At this poin all bets have been called, and or/all but one player folded.

	*/

	return nil
}

// #####################################################################
// findFirstSeatToAct determines who the first player to do sometning is, on given betting round.
func findFirstSeatToAct(tPtr *table, br bettingRound) (rPtr *ring.Ring, err error) {
	Info.Println(debugging.ThisFunc())

	rPtr = tPtr.remainingActiveSeatsRingPtr

	switch br {
	// Only PF starts with the first active seat after the BB
	case pflop:
		Info.Println("In swithch, P-FLOP")

		for {
			sPtr, _ := rPtr.Value.(*seat)
			Info.Println("sPtr.amBB: ", sPtr.amBB)
			Info.Println("seat dump: ", sPtr)
			if sPtr.amBB {
				rPtr = rPtr.Next()
				break
				// return rPtr, nil
			} else {
				rPtr = rPtr.Next()
			}
		}
	// All but pf start with first active seat after the BU
	default:
		Info.Println("In swithch, DEFAULT, betting round provided: ", br)

		/*
			for {
				sPtr, _ := rPtr.Value.(*seat)
				if sPtr.amButton {
					rPtr = rPtr.Next()
					// return rPtr, nil
				} else {
					rPtr = rPtr.Next()
				}
			}
		*/
	}
	return rPtr, nil
}

// #####################################################################
func executeNextStep(tPtr *table) error {

	Info.Println("")
	Info.Println(debugging.ThisFunc())
	Info.Println("")

	switch tPtr.hand.currentBettingRound {

	// case handStart:
	case pflop:
		Info.Println("executeNextStep : pflop")
		// Info.Println("executeNextStep : handStart")
		_ = startHand(tPtr)
		_ = postBlinds(tPtr)
		_ = updatePot(tPtr)

		// updateTableStateStruct(tPtr)
		// _ = showPokerRoomTables(prPtr)

		// Here is were we go around the table and let the players perform their actions.
		// To start with, we'll just have every player around the table FOLD to the BB.
		// The BB will get a walk and scoops the POT.

		_ = conductBettingRound(tPtr, pflop)

		_ = updatePot(tPtr)

		tPtr.hand.currentBettingRound = flop
		updateTableStateStruct(tPtr)
		// _ = showPokerRoomTables(prPtr)

		// This happens afer all seats finish all betting actions
		// _ = scoopBetsIntoPot(tPtr)

	case flop:
		Info.Println("executeNextStep : flop")
		_ = dealFlop(tPtr)

		_ = updatePot(tPtr)
		_ = scoopBetsIntoPot(tPtr)

		tPtr.hand.currentBettingRound = turn
		updateTableStateStruct(tPtr)
		// _ = showPokerRoomTables(prPtr)

	case turn:
		Info.Println("executeNextStep : turn")

		_ = dealTurn(tPtr)
		_ = updatePot(tPtr)

		tPtr.hand.currentBettingRound = river
		updateTableStateStruct(tPtr)
		// _ = showPokerRoomTables(prPtr)

	case river:
		Info.Println("executeNextStep : river")
		_ = dealRiver(tPtr)
		_ = updatePot(tPtr)

		tPtr.hand.currentBettingRound = handEnd
		updateTableStateStruct(tPtr)
		// _ = showPokerRoomTables(prPtr)

	case handEnd:
		Info.Println("executeNextStep : handEnd")

		// Force pot push during dev
		tPtr.timeToPushPot = true
		if tPtr.timeToPushPot {
			_ = pushPotsToWinners(tPtr)
		}
		tPtr.timeToPushPot = false

		_ = prepTableForNextHand(tPtr)
		_ = clearTableStateStruct(tPtr)

		tPtr.hand.currentBettingRound = pflop
		updateTableStateStruct(tPtr)
		// _ = showPokerRoomTables(prPtr)

	default:
		Info.Println("executeNextStep : default")
		Info.Println("")
	}

	return nil
}

// #####################################################################

// #####################################################################

// #####################################################################
