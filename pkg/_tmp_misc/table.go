package gameobjects

import (
	"container/ring"
	"strconv"

	"github.com/ttudrej/pokertrainer/pkg/debugging"
)

type seat struct {
	number        int      `0` // The number of this seat, as they appear around the table, 1-9.
	assignedToPID playerID `0`
	pPtr          *player  `nil`
	occupied      bool     // to indicate if player's chips are present
	sittingIn     bool     // in, means player is ready and intends to play the next hand, out, player is taking a break.
	stackSize     int
	c1Ptr         *Card
	c2Ptr         *Card
	betAmount     int  // for keeping track of the amount of a bet/call/raise.
	folded        bool // flag, telling the dealer if the player folded out already, and therefore should be skipped on the next betting round.
	allIn         bool // flag, if all in, player has no actions to perform when it's his turn again.
	amButton      bool `false` // Let's us know if we're currently have the dealer button
	amSB          bool `false`
	amBB          bool `false`
}

// var emptySeatPtr = &seat{number: 0, assignedToPID: 0, pPtr: noPlayerPtr, occupied: false, sittingIn: false}
// emptySeatPtr is reserved to mean that it has no player associted with it and no stack.
// It is NOT empyt becuse player occupying it is sitting out and left his stack behind.
// It is ready to be filled by another/new player.
var emptySeatPtr = &seat{
	number:        0,
	assignedToPID: 0,
	pPtr:          noPlayerPtr,
	occupied:      false,
	sittingIn:     false,
	stackSize:     0,
	c1Ptr:         NoCardPtr,
	c2Ptr:         NoCardPtr,
	betAmount:     0,
	folded:        false,
	allIn:         false,
	amButton:      false,
}

type seatPIDPrinter interface {
	printSeatPID(*seat) error
}

type seatPIDGetter interface {
	getSeatPID(*seat) (playerID, error)
}

// table represents a physical poker/game table
// It has all the aspects that will allow the game to progress, even if players are absent,
// for example, to pitch cards to an empty seats, and be able to determine the outcome of the hand,
// regardless of player action or incacion.
// It's the seat that "owns" the cards pitched to it, not the player.
// The player get's to "see" the cards, and to make betting/folding decisions.
// The player get's to choose the betting amounts, up to the stack size.
// The dealer with a table, even in absence of players, needs to be able to progress the game.
type table struct {
	numberSeats int // 2-10, usually
	buttonPos   int
	tableID     int // table serial number, so we can keep track of multiple tables, if needed.
	gameType    gameType
	deckPtr     *cardDeck
	// deckMapPtr             *cardDeckMap
	// deckOrderedListPtr     *listOfPtrsToCards
	// deckOrderedListFullPtr *listFullOfPtrsToCards

	// for holding the poiters to the flop cards
	// 0,1,2 - flop cards 1,2,3
	// 4,5 - turn and river
	communityCardsList          [6]*Card
	seatList                    [10]*seat
	seatNumbersToBeDealtIn      []int // I think it may be easier to think about seats in terms of their actual number, vs links/refs to them.
	paused                      bool  // flag set by the poker_rom_manager, indicating wheter the dealer can start the next hand.
	newTable                    bool  // if true, we'll need to "draw for the button"
	sbAmount                    int
	bbAmount                    int
	seatPtrSB                   *seat // We'll need to know which seat is SB/BB in different situations.
	seatPtrBB                   *seat
	remainingActiveSeatsRingPtr *ring.Ring // For keeping track of the active plyers still in the hand.
	tableStatePtr               *tableStateForTemplateAccess
	hand                        hand // for keeping track of the betting round sequence
	potSizePreviousRounds       int  // What was collected on all prevoius betting rounds
	potSizeThisRoundOnly        int  // For keeping track of bets/raises/re-raises on this round
	potTotal                    int  // Not physical, better, virtual count. Tells you what's in the pot
	// 									from previos rounds, and, what each player invested alredy on this round.
	// 									Represents a number, not the pile of chips.
	timeToPushPot bool
}

type tableStateForTemplateAccess struct {
	// The vars in this struct need to be exported, so must start with a Capital letter.
	// text/template or html/template can't grab them otherwise.
	TableID          int
	PotTotal         int
	ButtonMarkerList [11]string // Will hold a "D" in seat index correxponding to current button position.

	// Using 11 for max index, to accomodate 10 seat table.
	// We're not using index 0. Index will correspond to seat number.
	PidList       [11]playerID
	PUserNameList [11]userName
	StackList     [11]int
	BetAmount     [11]string // We use string, so have an option of using an empty string, "".
	SeatNumber    [11]int

	HoleCardsBackgroundColorClassList [11]string // for 2 color deck.
	// Used to name a CSS class for the BG color, face up, face down.
	// No Card is achieved by no BG color, so that the underlying color
	// of the tble or the player's seat comes through.

	C1rList                    [11]CardRank
	C1sList                    [11]CardSuit
	C1ColorClass               [11]string // Used to name a CSS class used with the card
	C1ColorClassSuitSymbol     [11]string // Used to name a CSS class used with the card
	C1sSymbolList              [11]string
	C1BackgroundColorClassList [11]string // for use in a 4 color deck

	C2rList                    [11]CardRank
	C2sList                    [11]CardSuit
	C2ColorClass               [11]string
	C2ColorClassSuitSymbol     [11]string // Used to name a CSS class used with the card
	C2sSymbolList              [11]string
	C2BackgroundColorClassList [11]string // for use in a 4 color deck

	ComCr                            [6]CardRank
	ComCs                            [6]CardSuit
	ComCColorClass                   [6]string
	ComCColorClassSuitSymbol         [11]string // Used to name a CSS class used with the card
	ComCSymbolList                   [6]string
	ComCardsBackgroundColorClassList [6]string
}

/*
type CircularList struct {
	list *ring.Ring
}
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

/*
The final incoming parameter in a function signature may have a type prefixed with "...".
A function with such a parameter is called variadic and may be invoked with zero or more
arguments for that parameter.

func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
*/

// #####################################################################
func (sPtr *seat) printSeatPID() (err error) {
	Info.Println(debugging.ThisFunc())
	Info.Println(sPtr.assignedToPID)
	return nil
}

// #####################################################################
func (sPtr *seat) getSeatPID() (playerID, error) {
	Info.Println(debugging.ThisFunc())
	return sPtr.assignedToPID, nil
}

// #####################################################################

// #####################################################################
// createTable constructs a new poker table (as in a kitchen table, not a tble of numbers), with a specific # of seats, 2-10.
// func createTable(numSeats int, tID tableId, gt gameType, gs gameSize) (tPtr *table, err error) {
// func createTable(gPtr *game) (tPtr *table, err error) {
func createTable(prPtr *pokerRoom, numSeats int, tID int, gt gameType) (tPtr *table, err error) {
	Info.Println(debugging.ThisFunc())

	var t table
	tPtr = &t

	Info.Println("Creating a table, table ID: ", tID)

	// To start the development, we'll assume that its' always a 9 seat table.
	numSeats = 9
	tPtr.numberSeats = numSeats

	tPtr.buttonPos = 0 // 0 indicates that button has not yet been assigned to any seat
	tPtr.tableID = tID
	tPtr.gameType = gt
	tPtr.paused = true
	tPtr.newTable = true
	tPtr.sbAmount = 1
	tPtr.bbAmount = 2
	tPtr.seatPtrSB = emptySeatPtr // easiest to id by "number == 0"
	tPtr.seatPtrBB = emptySeatPtr
	tPtr.potSizeThisRoundOnly = 0
	tPtr.potSizeThisRoundOnly = 0
	tPtr.potTotal = 0
	tPtr.timeToPushPot = false

	for i := range tPtr.communityCardsList {
		tPtr.communityCardsList[i] = NoCardPtr
	}

	for i := range tPtr.seatList {
		tPtr.seatList[i] = emptySeatPtr
	}

	// Point the table at the "stantdard" card deck
	tPtr.deckPtr, err = createDeck()

	// Info.Println(deckMap[cdmKey{r2, s}])
	// Info.Println(deckOrderedList[0])
	// Info.Println(deckOrderedListFull[0])

	tPtr.tableStatePtr, err = createTableStateStruct(tPtr)

	tPtr.hand.currentBettingRound = pflop

	return tPtr, err
}

// #####################################################################
// prepTableForNextHand "cleans" the table after the last hand, "gathers" cards, ie.
// initializes each players card display area, and the community card area.
func prepTableForNextHand(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())
	Info.Println("Preparing table for next hand, table ID: ", tPtr.tableID)

	for i := range tPtr.communityCardsList {
		tPtr.communityCardsList[i] = NoCardPtr
	}

	for i := range tPtr.seatList {
		tPtr.seatList[i].c1Ptr = NoCardPtr
		tPtr.seatList[i].c2Ptr = NoCardPtr
		tPtr.seatList[i].amSB = false
		tPtr.seatList[i].amBB = false
	}

	if tPtr.potTotal > 0 || tPtr.potSizePreviousRounds > 0 || tPtr.potSizeThisRoundOnly > 0 {
		Info.Println()
		Info.Println("POT amouts are NOT 0, and should be at thsi point !!")
		Info.Println()
	}

	return err
}

// #####################################################################

// #####################################################################
func displayTable(tPtr *table) (err error) {
	Info.Println(debugging.ThisFunc())
	// Trace.Println("I have something standard to say")
	Info.Println("Your log message here")
	// Warning.Println("There is something you need to know about")
	// Error.Println("Something has failed")

	Info.Println("####################################")
	Info.Println("Displaying table #: ", tPtr.tableID)
	Info.Println()

	// _, file_str, line_int, ok := runtime.Caller(0)
	// Info.Printf("%v, %v, %v, %v: %s", runtime.Caller(0), "MESSAGE")

	Info.Println("seats at table: ", tPtr.numberSeats)
	Info.Println("button pos: ", tPtr.buttonPos)

	Info.Println("table ID: ", tPtr.tableID)
	Info.Println("game type: ", tPtr.gameType)
	Info.Println("current betting round: ", tPtr.hand.currentBettingRound)
	Info.Println("POT total          : ", tPtr.potTotal)
	Info.Println("POT prev rounds    : ", tPtr.potSizePreviousRounds)
	Info.Println("POT this round only: ", tPtr.potSizeThisRoundOnly)

	Info.Println("time to push POT to winners: ", tPtr.timeToPushPot)
	Info.Println("SB/BB seat num: ", tPtr.seatPtrSB.number, tPtr.seatPtrBB.number)

	Info.Println()

	for i := range tPtr.communityCardsList {
		Info.Printf("flop card %v: %v\n", i+1, *tPtr.communityCardsList[i])
	}

	Info.Println()

	for j := range tPtr.seatList {
		// Info.Printf("Seat %v, PID: %v, occupied: %v, sitting in: %v, stack size: %v, c1: %v ptr: %p c2: %v ptr: %p\n", j+1, tPtr.seatList[j].assignedToPID, tPtr.seatList[j].occupied, tPtr.seatList[j].sittingIn, tPtr.seatList[j].stackSize, *tPtr.seatList[j].c1Ptr, tPtr.seatList[j].c1Ptr, *tPtr.seatList[j].c2Ptr, tPtr.seatList[j].c2Ptr)
		Info.Printf("Seat: %v, Seat self: %v, PID: %v, occ: %v, in?: %v, stack: %v, c1: %v, c2: %v\n", j+1, tPtr.seatList[j].number, tPtr.seatList[j].assignedToPID, tPtr.seatList[j].occupied, tPtr.seatList[j].sittingIn, tPtr.seatList[j].stackSize, *tPtr.seatList[j].c1Ptr, *tPtr.seatList[j].c2Ptr)
		// Info.Println(tPtr.seatList[j])
	}

	Info.Println()

	Info.Printf("table state struct for use in templates:\n%v\n\n", tPtr.tableStatePtr)

	Info.Println("----------")

	Info.Println()
	Info.Println("END of table #: ", tPtr.tableID)
	Info.Println("####################################")
	Info.Println()

	return err
}

// #################################################################
// createTableStateStruct creates and initializes a tableStateForTemplateAccess struct.
func createTableStateStruct(tPtr *table) (*tableStateForTemplateAccess, error) {
	Info.Println(debugging.ThisFunc())
	var tss tableStateForTemplateAccess
	tssPtr := &tss

	tss.PotTotal = tPtr.potTotal
	tss.TableID = tPtr.tableID

	for seatNum := 1; seatNum <= tPtr.numberSeats; seatNum++ {
		tss.PidList[seatNum] = noPlayerPtr.id
		tss.StackList[seatNum] = 0
		tss.BetAmount[seatNum] = "not set"
		tss.ButtonMarkerList[seatNum] = ""
		tss.HoleCardsBackgroundColorClassList[seatNum] = ""
		tss.PUserNameList[seatNum] = "not set"

		tss.C1rList[seatNum] = NoCardPtr.Rank
		tss.C1sList[seatNum] = NoCardPtr.Suit
		tss.C1ColorClass[seatNum] = "not set"
		tss.C1sSymbolList[seatNum] = "S"
		tss.C1BackgroundColorClassList[seatNum] = ""
		tss.C1ColorClassSuitSymbol[seatNum] = ""

		tss.C2rList[seatNum] = NoCardPtr.Rank
		tss.C2sList[seatNum] = NoCardPtr.Suit
		tss.C2ColorClass[seatNum] = "not set"
		tss.C2sSymbolList[seatNum] = "S"
		tss.C2BackgroundColorClassList[seatNum] = ""
		tss.C2ColorClassSuitSymbol[seatNum] = ""
	}

	for comCardNum := 1; comCardNum <= 5; comCardNum++ {
		tss.ComCr[comCardNum] = NoCardPtr.Rank
		tss.ComCs[comCardNum] = NoCardPtr.Suit
		tss.ComCColorClass[comCardNum] = "not set"
		tss.ComCSymbolList[comCardNum] = "S"
		tss.ComCardsBackgroundColorClassList[comCardNum] = ""
		tss.ComCColorClassSuitSymbol[comCardNum] = ""
	}

	return tssPtr, nil
}

// #################################################################
// updateTableStateStruct grabs a host of player and table specific var
// that are often stored in multi-level structs, and copies them to a
// single level struct, for ease of access with vi text/html/template
// package in the corresponding html files.
// De-referencing inside html files gets hairy quickly.
func updateTableStateStruct(tPtr *table) error {
	Info.Println(debugging.ThisFunc())

	tPtr.tableStatePtr.PotTotal = tPtr.potTotal

	for seatNum := 1; seatNum <= tPtr.numberSeats; seatNum++ {
		// Info.Println("*** seat num: ", seatNum)
		tPtr.tableStatePtr.PidList[seatNum] = tPtr.seatList[seatNum-1].assignedToPID
		tPtr.tableStatePtr.StackList[seatNum] = tPtr.seatList[seatNum-1].stackSize
		tPtr.tableStatePtr.PUserNameList[seatNum] = tPtr.seatList[seatNum-1].pPtr.userName

		// We want to convert an int into a string.
		if tPtr.seatList[seatNum-1].betAmount == 0 {
			tPtr.tableStatePtr.BetAmount[seatNum] = ""
		} else {
			// tPtr.tableStatePtr.BetAmount[seatNum] = string(tPtr.seatList[seatNum-1].betAmount)
			tPtr.tableStatePtr.BetAmount[seatNum] = strconv.Itoa(tPtr.seatList[seatNum-1].betAmount)
		}

		// PLAYER HOLE CARDS ############################################

		// Grab the Rank and Suit for card 1 and 2 from the seats
		tPtr.tableStatePtr.C1rList[seatNum] = tPtr.seatList[seatNum-1].c1Ptr.Rank
		tPtr.tableStatePtr.C1sList[seatNum] = tPtr.seatList[seatNum-1].c1Ptr.Suit

		tPtr.tableStatePtr.C2rList[seatNum] = tPtr.seatList[seatNum-1].c2Ptr.Rank
		tPtr.tableStatePtr.C2sList[seatNum] = tPtr.seatList[seatNum-1].c2Ptr.Suit

		// Clear out the Rank marking for any players noCard, so that they show up "blank" on the HTML page
		if tPtr.tableStatePtr.C1rList[seatNum] == RX {
			tPtr.tableStatePtr.C1rList[seatNum] = ""
			// We only do it for c1, since if it's true for c1, it must be true for c2. Player either has both cards or none, never just one.
			tPtr.tableStatePtr.HoleCardsBackgroundColorClassList[seatNum] = "" // 2 COLOR deck
			tPtr.tableStatePtr.C1BackgroundColorClassList[seatNum] = ""        // 4 COLOR deck
			tPtr.tableStatePtr.C1ColorClassSuitSymbol[seatNum] = ""

		} else {
			tPtr.tableStatePtr.HoleCardsBackgroundColorClassList[seatNum] = "card_bg_face_up" // 2 COLOR

			// 4 COLOR
			switch tPtr.tableStatePtr.C1sList[seatNum] {
			case H:
				tPtr.tableStatePtr.C1sSymbolList[seatNum] = "&hearts;"
				tPtr.tableStatePtr.C1ColorClass[seatNum] = "suit_color_4cd_h"
				tPtr.tableStatePtr.C1BackgroundColorClassList[seatNum] = "card_bg_face_up_4cd_h"
				tPtr.tableStatePtr.C1ColorClassSuitSymbol[seatNum] = "suit_symbol_color_4cd_h"
			case D:
				tPtr.tableStatePtr.C1sSymbolList[seatNum] = "&diams;"
				tPtr.tableStatePtr.C1ColorClass[seatNum] = "suit_color_4cd_d"
				tPtr.tableStatePtr.C1BackgroundColorClassList[seatNum] = "card_bg_face_up_4cd_d"
				tPtr.tableStatePtr.C1ColorClassSuitSymbol[seatNum] = "suit_symbol_color_4cd_d"
			case S:
				tPtr.tableStatePtr.C1sSymbolList[seatNum] = "&spades;"
				tPtr.tableStatePtr.C1ColorClass[seatNum] = "suit_color_4cd_s"
				tPtr.tableStatePtr.C1BackgroundColorClassList[seatNum] = "card_bg_face_up_4cd_s"
				tPtr.tableStatePtr.C1ColorClassSuitSymbol[seatNum] = "suit_symbol_color_4cd_s"
			case C:
				tPtr.tableStatePtr.C1sSymbolList[seatNum] = "&clubs;"
				tPtr.tableStatePtr.C1ColorClass[seatNum] = "suit_color_4cd_c"
				tPtr.tableStatePtr.C1BackgroundColorClassList[seatNum] = "card_bg_face_up_4cd_c"
				tPtr.tableStatePtr.C1ColorClassSuitSymbol[seatNum] = "suit_symbol_color_4cd_c"
			default:
				Info.Println("WARNING : Could not set css class name for bg suit color")
				tPtr.tableStatePtr.C1BackgroundColorClassList[seatNum] = "could not set"
			}
		}

		if tPtr.tableStatePtr.C2rList[seatNum] == RX {
			tPtr.tableStatePtr.C2rList[seatNum] = ""
			tPtr.tableStatePtr.C2BackgroundColorClassList[seatNum] = "" // 4 COLOR deck
			tPtr.tableStatePtr.C2ColorClassSuitSymbol[seatNum] = ""
		} else {

			switch tPtr.tableStatePtr.C2sList[seatNum] {
			case H:
				tPtr.tableStatePtr.C2sSymbolList[seatNum] = "&hearts;"
				tPtr.tableStatePtr.C2ColorClass[seatNum] = "suit_color_4cd_h"
				tPtr.tableStatePtr.C2BackgroundColorClassList[seatNum] = "card_bg_face_up_4cd_h"
				tPtr.tableStatePtr.C2ColorClassSuitSymbol[seatNum] = "suit_symbol_color_4cd_h"
			case D:
				tPtr.tableStatePtr.C2sSymbolList[seatNum] = "&diams;"
				tPtr.tableStatePtr.C2ColorClass[seatNum] = "suit_color_4cd_d"
				tPtr.tableStatePtr.C2BackgroundColorClassList[seatNum] = "card_bg_face_up_4cd_d"
				tPtr.tableStatePtr.C2ColorClassSuitSymbol[seatNum] = "suit_symbol_color_4cd_d"
			case S:
				tPtr.tableStatePtr.C2sSymbolList[seatNum] = "&spades;"
				tPtr.tableStatePtr.C2ColorClass[seatNum] = "suit_color_4cd_s"
				tPtr.tableStatePtr.C2BackgroundColorClassList[seatNum] = "card_bg_face_up_4cd_s"
				tPtr.tableStatePtr.C2ColorClassSuitSymbol[seatNum] = "suit_symbol_color_4cd_s"
			case C:
				tPtr.tableStatePtr.C2sSymbolList[seatNum] = "&clubs;"
				tPtr.tableStatePtr.C2ColorClass[seatNum] = "suit_color_4cd_c"
				tPtr.tableStatePtr.C2BackgroundColorClassList[seatNum] = "card_bg_face_up_4cd_c"
				tPtr.tableStatePtr.C2ColorClassSuitSymbol[seatNum] = "suit_symbol_color_4cd_c"
			default:
				Info.Println("WARNING : Could not set css class name for bg suit color")
				tPtr.tableStatePtr.C2BackgroundColorClassList[seatNum] = "could not set"
			}

		}

		// Clear the button assignment
		tPtr.tableStatePtr.ButtonMarkerList[seatNum] = ""
	}

	// COMMUNITY CARDS ############################################

	for comCardNum := 1; comCardNum <= 5; comCardNum++ {
		tPtr.tableStatePtr.ComCr[comCardNum] = tPtr.communityCardsList[comCardNum-1].Rank
		tPtr.tableStatePtr.ComCs[comCardNum] = tPtr.communityCardsList[comCardNum-1].Suit
		// tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = "card_bg_face_up"

		// Clear out the Rank marking for any community noCard, so that they show up "bank" on the HTML page
		if tPtr.tableStatePtr.ComCr[comCardNum] == RX {
			tPtr.tableStatePtr.ComCr[comCardNum] = ""
			tPtr.tableStatePtr.ComCSymbolList[comCardNum] = ""
			tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = ""
			tPtr.tableStatePtr.ComCColorClassSuitSymbol[comCardNum] = ""

		} else {
			tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = "card_bg_face_up"
		}

		// Assign appropriate CSS color class for the HTML page
		if tPtr.communityCardsList[comCardNum-1].Suit == S || tPtr.communityCardsList[comCardNum-1].Suit == C {
			tPtr.tableStatePtr.ComCColorClass[comCardNum] = "suit_color_black"
		} else {
			tPtr.tableStatePtr.ComCColorClass[comCardNum] = "suit_color_red"
		}

		// 4 COLOR Deck:
		switch tPtr.tableStatePtr.ComCs[comCardNum] {
		case S:
			tPtr.tableStatePtr.ComCs[comCardNum] = "&spades;"
			tPtr.tableStatePtr.ComCColorClass[comCardNum] = "suit_color_4cd_s"
			tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = "card_bg_face_up_4cd_s"
			tPtr.tableStatePtr.ComCColorClassSuitSymbol[comCardNum] = "suit_symbol_color_4cd_s"
		case C:
			tPtr.tableStatePtr.ComCs[comCardNum] = "&clubs;"
			tPtr.tableStatePtr.ComCColorClass[comCardNum] = "suit_color_4cd_c"
			tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = "card_bg_face_up_4cd_c"
			tPtr.tableStatePtr.ComCColorClassSuitSymbol[comCardNum] = "suit_symbol_color_4cd_c"
		case H:
			tPtr.tableStatePtr.ComCs[comCardNum] = "&hearts;"
			tPtr.tableStatePtr.ComCColorClass[comCardNum] = "suit_color_4cd_h"
			tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = "card_bg_face_up_4cd_h"
			tPtr.tableStatePtr.ComCColorClassSuitSymbol[comCardNum] = "suit_symbol_color_4cd_h"
		case D:
			tPtr.tableStatePtr.ComCs[comCardNum] = "&diams;"
			tPtr.tableStatePtr.ComCColorClass[comCardNum] = "suit_color_4cd_d"
			tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = "card_bg_face_up_4cd_d"
			tPtr.tableStatePtr.ComCColorClassSuitSymbol[comCardNum] = "suit_symbol_color_4cd_d"
		case NoCardPtr.Suit:
			// This card has not yet been dealt.
			tPtr.tableStatePtr.ComCs[comCardNum] = ""
		default:
			Info.Println("WARNING 03: Could not set symbol for the card suit")
			tPtr.tableStatePtr.ComCs[comCardNum] = "could not set"
		}
	}

	// BUTTON ############################################

	// Place button in the right seat on the HTML page
	tPtr.tableStatePtr.ButtonMarkerList[tPtr.buttonPos] = "D"

	return nil
}

// #################################################################
// clearTableStateStruct sets all values to blank/nil/default
func clearTableStateStruct(tPtr *table) error {
	Info.Println(debugging.ThisFunc())

	tPtr.tableStatePtr.PotTotal = 0

	for seatNum := 1; seatNum <= tPtr.numberSeats; seatNum++ {
		// Info.Println("*** seat num: ", seatNum)
		tPtr.tableStatePtr.PidList[seatNum] = noPlayerPtr.id
		tPtr.tableStatePtr.StackList[seatNum] = 0
		tPtr.tableStatePtr.PUserNameList[seatNum] = noPlayerPtr.userName
		tPtr.tableStatePtr.BetAmount[seatNum] = ""

		// PLAYER HOLE CARDS ############################################

		// Grab the Rank and Suit for card 1 and 2 from the seats
		tPtr.tableStatePtr.C1rList[seatNum] = NoCardPtr.Rank
		tPtr.tableStatePtr.C1sList[seatNum] = NoCardPtr.Suit
		tPtr.tableStatePtr.C1sSymbolList[seatNum] = ""

		tPtr.tableStatePtr.C2rList[seatNum] = NoCardPtr.Rank
		tPtr.tableStatePtr.C2sList[seatNum] = NoCardPtr.Suit
		tPtr.tableStatePtr.C2sSymbolList[seatNum] = ""

		tPtr.tableStatePtr.HoleCardsBackgroundColorClassList[seatNum] = "" // 2 COLOR deck

		tPtr.tableStatePtr.C1BackgroundColorClassList[seatNum] = "" // 4 COLOR deck
		tPtr.tableStatePtr.C1ColorClassSuitSymbol[seatNum] = ""

		tPtr.tableStatePtr.C2BackgroundColorClassList[seatNum] = "" // 4 COLOR deck
		tPtr.tableStatePtr.C2ColorClassSuitSymbol[seatNum] = ""

		// Clear the button assignment
		tPtr.tableStatePtr.ButtonMarkerList[seatNum] = ""
	}

	// COMMUNITY CARDS ############################################

	for comCardNum := 1; comCardNum <= 5; comCardNum++ {
		tPtr.tableStatePtr.ComCr[comCardNum] = NoCardPtr.Rank
		tPtr.tableStatePtr.ComCs[comCardNum] = NoCardPtr.Suit
		// tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = "card_bg_face_up"

		// Clear out the Rank marking for any community noCard, so that they show up "bank" on the HTML page
		tPtr.tableStatePtr.ComCr[comCardNum] = ""
		tPtr.tableStatePtr.ComCSymbolList[comCardNum] = ""
		tPtr.tableStatePtr.ComCardsBackgroundColorClassList[comCardNum] = ""
		tPtr.tableStatePtr.ComCColorClassSuitSymbol[comCardNum] = ""
	}

	// BUTTON ############################################

	// Place button in the right seat on the HTML page
	tPtr.tableStatePtr.ButtonMarkerList[tPtr.buttonPos] = "D"

	return nil
}

// #####################################################################
// resetTable resets the table struct to an "initial" state, so that we can "start over"
// with a pres of a button on the page, and won't have to stop/start the application.
func resetTable(tPtr *table) error {
	Info.Println(debugging.ThisFunc())

	Info.Println("Resetting table: ", tPtr.tableID)

	tPtr.buttonPos = 0 // 0 indicates that button has not yet been assigned to any seat
	tPtr.paused = true
	tPtr.newTable = true
	tPtr.seatPtrSB = emptySeatPtr // easiest to id by "number == 0"
	tPtr.seatPtrBB = emptySeatPtr
	tPtr.potSizeThisRoundOnly = 0
	tPtr.potSizeThisRoundOnly = 0
	tPtr.potTotal = 0
	tPtr.timeToPushPot = false

	for i := range tPtr.communityCardsList {
		tPtr.communityCardsList[i] = NoCardPtr
	}

	for i := range tPtr.seatList {
		tPtr.seatList[i] = emptySeatPtr
	}

	_ = clearTableStateStruct(tPtr)

	tPtr.hand.currentBettingRound = pflop

	return nil
}

// ##################################################
func getTablePtrFromTIDStr(tIDStr string) (*table, error) {
	Info.Println(debugging.ThisFunc())
	// tID, _ := strconv.ParseInt(tIDStr, 10, 8)
	tID, _ := getIntFromNumValStr(tIDStr)

	tPtr := prPtr.tableList[tID-1]

	return tPtr, nil
}
