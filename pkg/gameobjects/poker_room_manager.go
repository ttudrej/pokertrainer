package gameobjects

import (
	"github.com/ttudrej/pokertrainer/pkg/debugging"
	"github.com/ttudrej/pokertrainer/pkg/gameobjects"
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
func seatPlayer(pPtr *player, tPtr *table, seatNum int, buyinAmount int) (err error) {
	Info.Println(debugging.ThisFunc())
	// Info.Println("Seating player: ", pPtr.userName, " at a ", tPtr.gameType, " table number: ", tPtr.tableID, " seat: ", seatNum)

	i := seatNum - 1

	// Change out the global emtySeat for an acutal seat
	// var sPtr = &seat{pPtr.id, true, 300, gameobjects.NoCardPtr, gameobjects.NoCardPtr}

	var s seat
	sPtr := &s

	// Initialize the state of a New seat
	sPtr.number = seatNum
	sPtr.assignedToPID = pPtr.id
	sPtr.pPtr = pPtr
	sPtr.occupied = true
	sPtr.sittingIn = false // The player need to intentionally request participationin the game, after being seated.
	sPtr.stackSize = 300
	sPtr.c1Ptr = gameobjects.NoCardPtr
	sPtr.c2Ptr = gameobjects.NoCardPtr
	sPtr.betAmount = 0
	sPtr.allIn = false
	sPtr.folded = false

	tPtr.seatList[i] = sPtr

	return err
}

// #####################################################################
func setUpMockTable9Seat(prPtr *pokerRoom, tID int, pplPtr *playerPtrList) error {

	// _ = seatPlayer(p1Ptr, prPtr.tableList[tID], 1, 300)
	// _ = seatPlayer(p2Ptr, prPtr.tableList[tID], 2, 300)
	// _ = seatPlayer(p3Ptr, prPtr.tableList[tID], 3, 300)
	// _ = seatPlayer(p4Ptr, prPtr.tableList[tID], 4, 300)
	// _ = seatPlayer(p5Ptr, prPtr.tableList[tID], 5, 300)
	// _ = seatPlayer(p6Ptr, prPtr.tableList[tID], 6, 300)
	// _ = seatPlayer(p7Ptr, prPtr.tableList[tID], 7, 300)
	// _ = seatPlayer(p8Ptr, prPtr.tableList[tID], 8, 300)
	// _ = seatPlayer(p9Ptr, prPtr.tableList[tID], 9, 300)

	_ = seatPlayer(pplPtr[0], prPtr.tableList[tID], 1, 300)
	_ = seatPlayer(pplPtr[1], prPtr.tableList[tID], 2, 300)
	_ = seatPlayer(pplPtr[2], prPtr.tableList[tID], 3, 300)
	_ = seatPlayer(pplPtr[3], prPtr.tableList[tID], 4, 300)
	_ = seatPlayer(pplPtr[4], prPtr.tableList[tID], 5, 300)
	_ = seatPlayer(pplPtr[5], prPtr.tableList[tID], 6, 300)
	_ = seatPlayer(pplPtr[6], prPtr.tableList[tID], 7, 300)
	_ = seatPlayer(pplPtr[7], prPtr.tableList[tID], 8, 300)
	_ = seatPlayer(pplPtr[8], prPtr.tableList[tID], 9, 300)

	_ = announceIntentionToPlay(prPtr.tableList[tID], 1)
	_ = announceIntentionToPlay(prPtr.tableList[tID], 2)
	_ = announceIntentionToPlay(prPtr.tableList[tID], 3)
	_ = announceIntentionToPlay(prPtr.tableList[tID], 4)
	_ = announceIntentionToPlay(prPtr.tableList[tID], 5)
	_ = announceIntentionToPlay(prPtr.tableList[tID], 6)
	_ = announceIntentionToPlay(prPtr.tableList[tID], 7)
	_ = announceIntentionToPlay(prPtr.tableList[tID], 8)
	_ = announceIntentionToPlay(prPtr.tableList[tID], 9)

	return nil

}
