package main

import "github.com/ttudrej/pokertrainer/v2/debugging"

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
// updatePot updates the total amount in the Pot, based on what each players
// betAmount.
// We go around the table and add up each players bet amount.
// Since there can be bets/raiseas/re-raises, it will be easiest just to total up
// the entire table every time anyone puts money in, as opposed to keeping
// track of raises/re-reaises...
func updatePot(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())
	currentBetsTotal := 0

	for _, seatPtr := range tPtr.seatList {
		currentBetsTotal += seatPtr.betAmount
	}
	tPtr.potSizeThisRoundOnly = currentBetsTotal
	tPtr.potTotal = tPtr.potSizePreviousRounds + currentBetsTotal

	return nil
}

// #####################################################################
// scoopBetsIntoPot moves any chips from each seat, and puts it in the Pot.
// Meant for visual representation, rather than for keeping the track of the total amouint in the pot.
// Another func keeps trck of the total(s).
func scoopBetsIntoPot(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())
	for _, seatPtr := range tPtr.seatList {
		seatPtr.betAmount = 0
	}

	tPtr.potSizePreviousRounds = tPtr.potTotal

	return nil
}

// #####################################################################
func pushPotsToWinners(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())
	// we gonna start really simple, and for now, just push the POT to the BB, at end of every hand.
	tPtr.seatPtrBB.stackSize += tPtr.potTotal
	tPtr.potTotal = 0
	tPtr.potSizePreviousRounds = 0
	tPtr.potSizeThisRoundOnly = 0

	return nil
}

// #####################################################################
func pushPotToSeat(tPtr *table, sPtr *seat) (err error) {
	Info.Println(debugging.ThisFunc())

	sPtr.stackSize += tPtr.potTotal
	tPtr.potTotal = 0
	tPtr.potSizePreviousRounds = 0
	tPtr.potSizeThisRoundOnly = 0

	return nil
}
