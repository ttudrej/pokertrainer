package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

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

var prPtr *pokerRoom

var (
	// Trace *log.Logger
	Info *log.Logger
	// Warning *log.Logger
	// Error   *log.Logger
)

func Init(
	// traceHandle io.Writer,
	// infoHandle io.Writer,
	// warningHandle io.Writer,
	// errorHandle io.Writer) {

	infoHandle io.Writer) {

	/*
		Trace = log.New(traceHandle,
			"TRACE: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Warning = log.New(warningHandle,
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Error = log.New(errorHandle,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	*/

	Info = log.New(infoHandle,
		"INFO  : ",
		log.Lshortfile)

}

func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("TRACE : %s : %d - %s\n", file, line, f.Name())
}

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

// main func doc line
func main() {
	// Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Init(os.Stdout)

	/*
		shuffledListOfCardsPtr, err01 := shuffleDeck(deckOrderedListPtr)
		if err01 != nil {
			Info.Println("got us an error:", err01)
		}
	*/

	// _ = displayCardsInList(deckOrderedListPtr)
	// Info.Println()
	// _ = displayCardsInList(shuffledListOfCardsPtr)

	prPtr, _ = createPokerRoom()
	// _ = showPokerRoomTables(prPtr)

	p1Ptr, _ := createPlayer("plyr 1", 1)
	p2Ptr, _ := createPlayer("plyr 2", 2)
	p3Ptr, _ := createPlayer("plyr 3", 3)
	p4Ptr, _ := createPlayer("plyr 4", 4)
	p5Ptr, _ := createPlayer("Hero", 5)
	p6Ptr, _ := createPlayer("plry 6", 6)
	p7Ptr, _ := createPlayer("plyr 7", 7)
	p8Ptr, _ := createPlayer("plry 8", 8)
	p9Ptr, _ := createPlayer("plyr 9", 9)

	var pPtrListPtr *playerPtrList
	pPtrListPtr, _ = createPlayerPtrListPtr(p1Ptr, p2Ptr, p3Ptr, p4Ptr, p5Ptr, p6Ptr, p7Ptr, p8Ptr, p9Ptr)

	_ = setUpMockTable9Seat(prPtr, 0, pPtrListPtr)

	_ = prepareHandAnalysisTools()

	// ####################
	// set up table 2

	// _ = seatPlayer(p1Ptr, prPtr.tableList[1], 1, 300)
	// _ = seatPlayer(p2Ptr, prPtr.tableList[1], 2, 300)
	// _ = seatPlayer(p3Ptr, prPtr.tableList[1], 9, 300)
	// _ = seatPlayer(p4Ptr, prPtr.tableList[1], 8, 300)
	// _ = seatPlayer(p5Ptr, prPtr.tableList[1], 7, 300)
	// _ = seatPlayer(p6Ptr, prPtr.tableList[1], 6, 300)
	// _ = seatPlayer(p7Ptr, prPtr.tableList[1], 5, 300)
	// _ = seatPlayer(p8Ptr, prPtr.tableList[1], 4, 300)
	// _ = seatPlayer(p9Ptr, prPtr.tableList[1], 3, 300)

	// _ = announceIntentionToPlay(prPtr.tableList[1], 1)
	// _ = announceIntentionToPlay(prPtr.tableList[1], 2)
	// _ = announceIntentionToPlay(prPtr.tableList[1], 3)
	// _ = announceIntentionToPlay(prPtr.tableList[1], 4)
	// _ = announceIntentionToPlay(prPtr.tableList[1], 5)
	// _ = announceIntentionToPlay(prPtr.tableList[1], 6)
	// _ = announceIntentionToPlay(prPtr.tableList[1], 7)
	// _ = announceIntentionToPlay(prPtr.tableList[1], 8)
	// _ = announceIntentionToPlay(prPtr.tableList[1], 9)

	// ####################

	// Info.Println(" seat nums to be delt in: ", prPtr.tableList[0].seatNumbersToBeDealtIn)

	// ####################
	// _ = executeHand(prPtr.tableList[0])
	// _ = executeHand(prPtr.tableList[1])

	// ####################
	// _ = showPokerRoomTables(prPtr)

	startWebServer()

	Info.Println(drawCompleteRangeHTML())

}
