package experimenting

import (
	"fmt"
)

// #####################################################################
// Functions for creating a full list of all possible hands.

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!! As a convention, any cardList MUST be ordered by rank, A to 2.
// Functions that process RELY on the order being correct
// This will simpify some code down the line, as certain assumptions will
// be valid, and certain manipulations no longer necessary.
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// #####################################################################
// createSFCL - doc line
func createSFCL() (clPtr *cardList) {
	clPtr = createClPtr()

	j := 39 // Legal index range is 0-39; 39 + 4*4(16) = 55(56-1), 5h sf
	j = 0   // 0-3: A high sf, 4-7: K high SF, ... 36-39: 5h SF
	clPtr[0] = orderedListFull[j]
	clPtr[1] = orderedListFull[j+4]
	clPtr[2] = orderedListFull[j+8]
	clPtr[3] = orderedListFull[j+12]
	clPtr[4] = orderedListFull[j+16]
	// need 2 more for 7 cards total
	clPtr[5] = orderedListFull[j+30]
	clPtr[6] = orderedListFull[j+40]

	return clPtr
}

// #####################################################################
// create4xCL - doc line
func create4xCL() (clPtr *cardList) {
	clPtr = createClPtr()

	clPtr[0] = orderedList[0]
	clPtr[1] = orderedList[1]
	clPtr[2] = orderedList[2]
	clPtr[3] = orderedList[3]
	clPtr[4] = orderedList[4]
	// need 2 more for 7 cards total
	clPtr[5] = orderedList[18]
	clPtr[6] = orderedList[45]

	return clPtr
}

// #####################################################################
// createFHCL - doc line
func createFHCL() (clPtr *cardList) {
	clPtr = createClPtr()

	// FH 3x, 3x, 1x
	clPtr[0] = orderedList[0]
	clPtr[1] = orderedList[1]
	clPtr[2] = orderedList[2]

	clPtr[3] = orderedList[4]
	clPtr[4] = orderedList[5]
	clPtr[5] = orderedList[6]

	clPtr[6] = orderedList[45]

	return clPtr
}

// #####################################################################
// createFlCL - doc line
func createFlCL() (clPtr *cardList) {
	clPtr = createClPtr()

	j := 0 //

	clPtr[0] = orderedList[j]
	clPtr[1] = orderedList[j+4]
	clPtr[2] = orderedList[j+8]
	clPtr[3] = orderedList[j+12]
	clPtr[4] = orderedList[j+24]

	clPtr[5] = orderedList[j+30]
	clPtr[6] = orderedList[j+45]

	return clPtr
}

// #####################################################################
// createStCL - doc line
func createStCL() (clPtr *cardList) {
	clPtr = createClPtr()

	clPtr[0] = orderedListFull[0]
	clPtr[1] = orderedListFull[42]
	clPtr[2] = orderedListFull[47]
	clPtr[3] = orderedListFull[50]
	clPtr[4] = orderedListFull[39]

	clPtr[5] = orderedListFull[5]
	clPtr[6] = orderedListFull[9]

	return clPtr
}

// #####################################################################

// create3xCL - doc line. Same def exists in hand_analysis.go
func create3xCL() (clPtr *cardList) {
	clPtr = createClPtr()

	// 3x
	// j := 39 // Legal index range is 0-39; 39 + 4*4(16) = 55(56-1), 5h sf
	j := 0 // 0-3: A high sf, 4-7: K high SF, ... 36-39: 5h SF

	clPtr[0] = orderedList[j]
	clPtr[1] = orderedList[j+1]
	clPtr[2] = orderedList[j+3]

	clPtr[3] = orderedList[20]
	clPtr[4] = orderedList[25]
	clPtr[5] = orderedList[30]
	clPtr[6] = orderedList[35]

	return clPtr
}

// #####################################################################
// create2x2CL - doc line
func create2x2CL() (clPtr *cardList) {
	clPtr = createClPtr()

	// Two pair
	clPtr[0] = orderedList[0]
	clPtr[1] = orderedList[1]
	clPtr[2] = orderedList[4]
	clPtr[3] = orderedList[5]
	clPtr[4] = orderedList[39]

	clPtr[5] = orderedList[30]
	clPtr[6] = orderedList[40]

	return clPtr
}

// #####################################################################
// create2x1CL - doc line
func create2x1CL() (clPtr *cardList) {
	clPtr = createClPtr()

	// Two pair
	clPtr[0] = orderedList[0]
	clPtr[1] = orderedList[1]
	clPtr[2] = orderedList[10]
	clPtr[3] = orderedList[15]
	clPtr[4] = orderedList[39]

	clPtr[5] = orderedList[30]
	clPtr[6] = orderedList[40]

	return clPtr
}

// #####################################################################
// createHcCL - doc line
func createHcCL() (clPtr *cardList) {
	clPtr = createClPtr()

	// Two pair
	clPtr[0] = orderedList[0]
	clPtr[1] = orderedList[5]
	clPtr[2] = orderedList[10]
	clPtr[3] = orderedList[15]
	clPtr[4] = orderedList[39]
	clPtr[5] = orderedList[30]
	clPtr[6] = orderedList[40]

	return clPtr
}

// #####################################################################
// runCheckForHands exercises the findBestHandInCardList func, to give us some confidence it's doing
// the right thing.

func runCheckForHands() (err error) {

	var bestHand string

	fmt.Println("###################################################################### SF check")
	bestHand = findBestHandInCardList(createSFCL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	fmt.Println("###################################################################### 4x check")
	bestHand = findBestHandInCardList(create4xCL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	fmt.Println("###################################################################### FH check")
	bestHand = findBestHandInCardList(createFHCL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	fmt.Println("###################################################################### Fl check")
	bestHand = findBestHandInCardList(createFlCL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	fmt.Println("###################################################################### St check")
	bestHand = findBestHandInCardList(createStCL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	fmt.Println("###################################################################### 3x check")
	bestHand = findBestHandInCardList(create3xCL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	fmt.Println("###################################################################### 2x2 check")
	bestHand = findBestHandInCardList(create2x2CL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	fmt.Println("###################################################################### 2x1 check")
	bestHand = findBestHandInCardList(create2x1CL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	fmt.Println("###################################################################### 2x1 check")
	bestHand = findBestHandInCardList(createHcCL())
	fmt.Println("best hand ALL: ", bestHand)
	fmt.Println()

	return err
}
