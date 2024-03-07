package main

import (
	"container/ring"
	"fmt"
	"log"
	"strconv"

	"github.com/ttudrej/pokertrainer/pkg/debugging"
)

var (
	// Trace *log.Logger
	Info *log.Logger
	// Warning *log.Logger
	// Error   *log.Logger
)

type playerID int

type userName string

type player struct {
	userName            userName
	id                  playerID
	behaviorDescription string

	// Starting ranges; sr9h = starting range 9 handed
	// The positions are relative to the button, going couter clockwise.
	// When players step away from the tble, the first relatie position to
	// be removed will be the U0, then U1, etc.
	// We choose to assign the position names this way on "non full" tables,
	// since it's the number of players behing you, that is considered most critical.

	/*
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
	*/
}

var noPlayerPtr = &player{"noPlayer", 0, "noBehavior defined"}

type playerPtrList [9]*player
type actionID string

const (
	// Using caps for first letter, NOT for ability to export
	// but for better visibility/contrast in positionStatus structs
	FLD actionID = "fold"
	out actionID = "out, folded already" // Indicates that this seat does not act, they folded earlier.
	CLL actionID = "call"
	BT0 actionID = "0 bet / check - Post Flop"
	BT1 actionID = "1 bet / pf BB"
	BT2 actionID = "2 bet / 1st raise / raise / pf open"
	BT3 actionID = "3 bet / 2nd raise / re-raise"
	BT4 actionID = "4 bet / 3rd raise / re-re-raise"
	BT5 actionID = "5 bet"
	BT6 actionID = "6 bet"
	CHK actionID = "check"  // post
	NOP actionID = "no op"  // chk / we bet / call, used in positionStatus vars
	WIN actionID = "winner" // to indicate we're the only one left in the hand

)

// playerInstanceState is meant to capture the "state" of a specific player, in a specific seat,
// at a specific table, at a specific moment in the hand.
// This will be some of the parameters that will go into making a decision by this player, and
// by ohters, about this player.
// We use the word "instance", since we'll allow for the same player to play at multiple tables at the same time.
type playerInstanceState struct {
	pPtr                         *player
	sPtr                         *seat
	tPtr                         *table
	spr                          float32     // Stack to Pot Ratio
	possibleActionIDactionIDList [3]actionID // fold/call/raise
	relPosPF                     relativePositionPF
	relPos                       relativePosition
	actionSeqDetail_PF           positionStatusPF
	actionSeqDetail_F            positionStatus
	actionSeqDetail_T            positionStatus
	actionSeqDetail_R            positionStatus
}

/*
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

// ##################################################
func createPlayer(un userName, pid playerID) (pPtr *player, err error) {
	Info.Println(debugging.ThisFunc())
	var p player
	pPtr = &p

	p.userName = un
	p.id = pid

	return pPtr, err
}

// ##################################################
// Indicate to the table that we want to participate in the game, ie. gimme some cards to play with.
func announceIntentionToPlay(tPtr *table, seatNum int) (err error) {
	Info.Println(debugging.ThisFunc())
	tPtr.seatList[seatNum-1].sittingIn = true
	// Info.Println("player id wants to play: ", tPtr.seatList[seatNum-1].pPtr.userName)
	return err
}

// ##################################################
// Indicate to the table that we want to participate in the game, ie. gimme some cards to play with.
func announceIntentionToSitOut(tPtr *table, seatNum int) (err error) {
	Info.Println(debugging.ThisFunc())
	tPtr.seatList[seatNum-1].sittingIn = false
	Info.Println("player is sitting out: ", tPtr.seatList[seatNum-1].pPtr.userName)
	return err
}

// ##################################################
func postBet(tIDStr, pIDStr, fractionStr string) (err error) {
	// the values passed to us come in as string, so we need to deal with that first.
	Info.Println(debugging.ThisFunc())

	tID, _ := getIntFromNumValStr(tIDStr)
	tPtr, _ := getTablePtrFromTIDStr(tIDStr)
	// pID, _ := getPlayerIDFromPIDString(pIDStr)
	pID, _ := getIntFromNumValStr(pIDStr)
	fraction, _ := getFloatFromStringNumVal(fractionStr)

	Info.Println()
	Info.Println("postBet; tID / pID / fractin: ", tIDStr, pIDStr, fractionStr)
	Info.Println("postBet; tID / pID / fractin: ", tID, pID, fraction)
	Info.Println()

	// Get a table pointer based on the table ID
	// tPtr := prPtr.tableList[tID-1]
	// Get player pointer based on table ID and player ID
	pPtr := tPtr.seatList[pID-1]

	frOfPot := int(float64(tPtr.potTotal) * fraction)

	pPtr.betAmount = frOfPot

	return nil
}

// ##################################################
func createPlayerPtrListPtr(p1Ptr, p2Ptr, p3Ptr, p4Ptr, p5Ptr, p6Ptr, p7Ptr, p8Ptr, p9Ptr *player) (pplPtr *playerPtrList, err error) {

	var ppl playerPtrList
	pplPtr = &ppl

	pplPtr[0] = p1Ptr
	pplPtr[1] = p2Ptr
	pplPtr[2] = p3Ptr
	pplPtr[3] = p4Ptr
	pplPtr[4] = p5Ptr
	pplPtr[5] = p6Ptr
	pplPtr[6] = p7Ptr
	pplPtr[7] = p8Ptr
	pplPtr[8] = p9Ptr

	return pplPtr, err
}

// ##################################################
// performAction is where the player get's to apply their game strategy, in response
// to the current state and progress of the hand.
func performAction(sPtr *seat, rPtr *ring.Ring, tPtr *table) error {

	Info.Println("Wroking on seat num / PID:", sPtr.number, sPtr.assignedToPID)

	switch tPtr.hand.currentBettingRound {
	case pflop:
		// Decide which acion is best to take
		chosenAction, amount, _ := determineBestAction(tPtr, pflop)

		switch chosenAction {
		case FLD:
			Info.Println("PID folded", sPtr.assignedToPID)
			Info.Println("bet amount was: ", amount)

			// XXX  ??? What needs to happen when we fold ??? XXX
			sPtr.folded = true

			// Go back one element, so tht we can exec the Unlink on "this seat"
			rPtr = rPtr.Prev()
			// Remove the NEXT element
			rPtr.Unlink(1)

		case CLL:
			fmt.Println("we called")
			fmt.Println("bet amount was: ", amount)
		case CHK:
			fmt.Println("we checked")
			fmt.Println("bet amount was: ", amount)
		case BT1:
			fmt.Println("we bet")
			fmt.Println("bet amount was: ", amount)
		case BT2:
		case BT3:
		case BT4:
		case BT5:
		case BT6:
		default:
			Info.Println("ERROR: Reached default action")
		}
	case flop:
		Info.Println()
	case turn:
		Info.Println()
	case river:
		Info.Println()
	default:
		Info.Println("in swithch tPtr.hand.currentBettingRound : default case")

	}

	return nil
}

// ##################################################
func fold(tPtr *table, sPtr *seat) error {

	return nil
}

/*
// ##################################################
func getPlayerIDFromPIDString(pIDStr string) (int, error) {

	pID, _ := strconv.ParseInt(pIDStr, 10, 8) // Produces int64

	Info.Println("pIDStr: ", pIDStr)
	Info.Println("pID   : ", pID)

	return int(pID), nil
}
*/

// ##################################################
func getIntFromNumValStr(s string) (int, error) {
	i64, _ := strconv.ParseInt(s, 10, 8) // Produces int64

	Info.Println("iStr: ", s)
	Info.Println("i64   : ", i64)

	return int(i64), nil
}

// ##################################################
func getFloatFromStringNumVal(numValStr string) (float64, error) {

	f, _ := strconv.ParseFloat(numValStr, 64)

	return f, nil
}

// ##################################################
