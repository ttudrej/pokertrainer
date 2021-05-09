package main

import "github.com/ttudrej/pokertrainer/v2/debugging"

const (
	maxAllowedTables int = 1
)

type pokerRoom struct {
	tableList [maxAllowedTables]*table
	// We will use one card deck for the entire poker room.
	// There is no need to have one per table. With many tables, this will save memory.
	// Each table will only need to keep track of the list of pointers to cards in this deck,
	// and it is those lists that will be shuffled, therefore making each tables card list
	// independant.
	// Having single "global deck", will also allow for some hand comparison operations
	// to be done in a more "minimal" and less compicated fashion, than would be otherwise
	// possible, if we were forced to use the decks on specific tables.
	cardDeckPtr *cardDeck
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

// #####################################################################
func createPokerRoom() (prPtr *pokerRoom, err error) {
	Info.Println(debugging.ThisFunc())
	var pr pokerRoom
	prPtr = &pr

	// Info.Println("Opening TT's \"Poker Smoker\" Training Room ... ", "\n")
	Info.Println("Opening TT's \"Poker Do-jo\" ... ")
	Info.Println()

	for i := 1; i <= maxAllowedTables; i++ {
		pr.tableList[i-1], _ = createTable(prPtr, 9, i, nlh)
	}

	Info.Println()
	return prPtr, err
}

// #####################################################################
func showPokerRoomTables(prPtr *pokerRoom) (err error) {
	Info.Println(debugging.ThisFunc())
	for i := range prPtr.tableList {
		err = displayTable(prPtr.tableList[i])
	}

	return err
}
