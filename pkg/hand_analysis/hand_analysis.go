package hand_analysis

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sort"
// 	"strconv"

// 	"github.com/ttudrej/pokertrainer/pkg/debugging"
// 	"github.com/ttudrej/pokertrainer/pkg/manage_table"
// 	"gonum.org/v1/gonum/stat/combin"
// )

// var (
// 	// Trace *log.Logger
// 	Info *log.Logger
// 	// Warning *log.Logger
// 	// Error   *log.Logger
// )

// // csrdIndex comment
// type cardIndex []manage_table.CardString

// var cardIndexS = cardIndex{manage_table.SA, manage_table.SK, manage_table.SQ, manage_table.SJ, manage_table.ST, manage_table.S9, manage_table.S8, manage_table.S7, manage_table.S6, manage_table.S5, manage_table.S4, manage_table.S3, manage_table.S2}
// var cardIndexC = cardIndex{manage_table.CA, manage_table.CK, manage_table.CQ, manage_table.CJ, manage_table.CT, manage_table.C9, manage_table.C8, manage_table.C7, manage_table.C6, manage_table.C5, manage_table.C4, manage_table.C3, manage_table.C2}
// var cardIndexH = cardIndex{manage_table.HA, manage_table.HK, manage_table.HQ, manage_table.HJ, manage_table.HT, manage_table.H9, manage_table.H8, manage_table.H7, manage_table.H6, manage_table.H5, manage_table.H4, manage_table.H3, manage_table.H2}
// var cardIndexD = cardIndex{manage_table.DA, manage_table.DK, manage_table.DQ, manage_table.DJ, manage_table.DT, manage_table.D9, manage_table.D8, manage_table.D7, manage_table.D6, manage_table.D5, manage_table.D4, manage_table.D3, manage_table.D2}

// var cardIndexFull = cardIndex{
// 	manage_table.SA, manage_table.HA, manage_table.CA, manage_table.DA,
// 	manage_table.SK, manage_table.HK, manage_table.CK, manage_table.DK,
// 	manage_table.SQ, manage_table.HQ, manage_table.CQ, manage_table.DQ,
// 	manage_table.SJ, manage_table.HJ, manage_table.CJ, manage_table.DJ,
// 	manage_table.ST, manage_table.HT, manage_table.CT, manage_table.DT,
// 	manage_table.S9, manage_table.H9, manage_table.C9, manage_table.D9,
// 	manage_table.S8, manage_table.H8, manage_table.C8, manage_table.D8,
// 	manage_table.S7, manage_table.H7, manage_table.C7, manage_table.D7,
// 	manage_table.S6, manage_table.H6, manage_table.C6, manage_table.D6,
// 	manage_table.S5, manage_table.H5, manage_table.C5, manage_table.D5,
// 	manage_table.S4, manage_table.H4, manage_table.C4, manage_table.D4,
// 	manage_table.S3, manage_table.H3, manage_table.C3, manage_table.D3,
// 	manage_table.S2, manage_table.H2, manage_table.C2, manage_table.D2}

// // fiveCardHandKind, Not using "fiveCardHandType", so as to NOT confuse it with "type" as a declaration directive.
type fiveCardHandKind string

type fiveCardHandKindRanking struct {
	handKind fiveCardHandKind
	typeRank int
}

const (
	sf  fiveCardHandKind = "Straight Flush"
	x4  fiveCardHandKind = "Four of a Kind"
	fh  fiveCardHandKind = "Full House"
	fl  fiveCardHandKind = "Flush"
	st  fiveCardHandKind = "Straight"
	x3  fiveCardHandKind = "Three of a Kind"
	x22 fiveCardHandKind = "Two Pair"
	x2  fiveCardHandKind = "Pair"
	hc  fiveCardHandKind = "High Card"
)

// FiveCardHandKindRankingFullList is a struct containing the ranking info. We assign integer rank to hand types. Used in comparing hand strength.
type FiveCardHandKindRankingFullList struct { // fchkr == fiveCardHandKindRanking
	SfInfo  fiveCardHandKindRanking
	X4Info  fiveCardHandKindRanking
	FhInfo  fiveCardHandKindRanking
	FlInfo  fiveCardHandKindRanking
	StInfo  fiveCardHandKindRanking
	X3Info  fiveCardHandKindRanking
	X22Info fiveCardHandKindRanking
	X2Info  fiveCardHandKindRanking
	HcInfo  fiveCardHandKindRanking
}

// /*
// // We're assuming that we'll be comparing ordered lists of cards only.
// type fiveCardHandOrderedList struct { // 0 index == highest rank; higher index == lower rank
// 	handType fiveCardHandKind
// 	handList [2598960]fiveCardHand // Num of combos of 5 pick 52

// 	// We now need to prepare data structures that will support "ranking" 2 or more 5 card hands.
// 	// Seems that a 2 level rannking system will work well, first, compare by fiveCardHandKind, then by an ordered list of all hands of that type.
// 	// This will allow us to have a "ranking scale" which avoids calculating and properly ordering all the 5 card combinations. We only
// 	// need to order within each hand kind.
// }
// */

// type fiveCardHandMetadata struct {
// 	index    int // The ssi index of hand, corresponds to the sscp index
// 	handKind fiveCardHandKind
// }

// var crm manage_table.CardRankMap

// // fiveCardList needs at most the max number of commuity cards + max num of the hole cards, so 7 for Texas NLH
// type fiveCardList []*manage_table.Card

// type equivalentFiveCardHand struct {
// 	// This is what we'll be collapsing equivalent hands to
// 	c1r   manage_table.CardRank // card 1 rank
// 	c2r   manage_table.CardRank // card 2 rank
// 	c3r   manage_table.CardRank
// 	c4r   manage_table.CardRank
// 	c5r   manage_table.CardRank
// 	info  fiveCardHandKindRanking
// 	count int // Keep track of how many equivalent hands we've found.
// }

// // Keep a sorted list, by hand rank, of equivalent hands with some additional info
// type equivalentFCHList []equivalentFiveCardHand

// type rankCtrMap map[manage_table.CardRank]int
// type suitCtrMap map[manage_table.CardSuit]int

// // rankCouter keeps track of counts of specific card ranks in a card list. This provides means for determining the
// // relative strength of a specific list of cards.
// type rankCounter struct {
// 	max         int
// 	uniqeRankCt int                   // count of all unique ranks in the list
// 	top4x1      manage_table.CardRank // rank of the top quads

// 	top3x1 manage_table.CardRank // rank of the top trips / three of a kind
// 	top3x2 manage_table.CardRank // rank of the 2nd trips / three of a kind

// 	top2x1 manage_table.CardRank // rank of the top pair
// 	top2x2 manage_table.CardRank // rank of the second pair ...
// 	top2x3 manage_table.CardRank

// 	top1x1 manage_table.CardRank
// 	top1x2 manage_table.CardRank
// 	top1x3 manage_table.CardRank
// 	top1x4 manage_table.CardRank
// 	top1x5 manage_table.CardRank

// 	rcm rankCtrMap // hold how many of each rank there are in a card list
// }

// // suitCounter keeps track of suit composition in a list of cards. Means for detecting flushes.
// type suitCounter struct {
// 	max int
// 	scm suitCtrMap
// }

// // orderedListOfPtrsToCard
// type orderedListOfPtrsToCards [52]*manage_table.Card

// // orderedListOfPtrsToCard uses 56 not 52 slots, to accomodate for the Aces in 5-A straights
// // Used for hand rank checks ONLY
// type orderedListFullOfPtrsToCards [56]*manage_table.Card

// var orderedList orderedListOfPtrsToCards
// var orderedListFull orderedListFullOfPtrsToCards

// var counterSF int  // Straight Flush
// var counter4x int  // 4 of a kind / quads
// var counterFH int  // Full House
// var counterFl int  // Flush
// var counterSt int  // Straight
// var counter3x int  // 3 of a kind / trips / set
// var counter2x2 int // 2x 2 of a kind / 2 pair
// var counter2x int  // 2 of a kind / pair
// var counterHC int  // High card

// // var scl []fiveCardList
// var sOfAllFCLs []fiveCardList

// // For use by the sorting fuctions
// // type fclByRank [5]*card
// // type sfIndexByFirstmanage_table.CardRank []int

// var fchkrFL FiveCardHandKindRankingFullList

// /*
// #########################################################################
// #########################################################################
// #########################################################################

// ######## ##     ## ##    ##  ######   ######
// ##       ##     ## ###   ## ##    ## ##    ##
// ##       ##     ## ####  ## ##       ##
// ######   ##     ## ## ## ## ##        ######
// ##       ##     ## ##  #### ##             ##
// ##       ##     ## ##   ### ##    ## ##    ##
// ##        #######  ##    ##  ######   ######

// #########################################################################
// #########################################################################
// #########################################################################
// */

// // #########################################################################

// // prepareHandAnalysisTools doc line
// func prepareHandAnalysisTools() error {

// 	// Info.Printf("%s\n\n", ThisFunc())
// 	// Give ourselves a private/unique deck to work with
// 	// deckPtr, _ := createDeck()
// 	deckPtr, _ := manage_table.CreateDeck()

// 	// Generate an integer representation of 52 choose 5 combos, so a list of all possible 5 card hands
// 	// We will later create a mapping, which will translate the integers to proper card representation.
// 	ssiPtr, numberOfCombos, _ := generateAll5CardIntegerCombos(52, 5)

// 	// Create a map for rank value look up, since constant structs are not supported, so we could
// 	// not have a rank struct with indexes, as a constant.
// 	// crm = manage_table.CreateCardRankMap()
// 	// crm, _ := manage_table.CardRankMapCreator.Create(manage_table.CardRankMapStruct{})

// 	// ###############################

// 	fmt.Println()

// 	// _ = printSSIAsCardStrings(ssiPtr)

// 	fmt.Println("Num of combos: ", numberOfCombos)
// 	fmt.Println()

// 	// ###############################
// 	// Generate a list of all hand pointers
// 	// sOfAllFCLsPtr := &sOfAllFCLs
// 	sOfAllFCLsPtr, _ := genSCLfromSSI(ssiPtr, deckPtr)
// 	sOfAllFCLs = *sOfAllFCLsPtr

// 	// fmt.Println(sOfAllFCLsPtr)

// 	_ = sortCardsInEach5CList()

// 	// ###############################
// 	// _ = printAllHandsInList(sscpPtr)
// 	// _ = printAllHandsInList()

// 	fchkrFL, _ = CreateFiveCardHandKindRankings()
// 	// fmt.Println(fchkrFL.SfInfo, fchkrFL.X4Info, FhInfo, FlInfo, StInfo, X3Info, X22Info, X2Info, HcInfo)

// 	fmt.Println(fchkrFL.SfInfo)

// 	// fmt.Printf("map len: %v\n", len(*deckPtr.cdmPtr))

// 	// _ = identifyAllSFs()

// 	// Generate lists of indexes by hand type
// 	sOfSFIndexesPtr, sOf4xIndexesPtr, sOfFHIndexesPtr, sOfFlIndexesPtr, sOfStIndexesPtr, sOf3xIndexesPtr, sOf2x2IndexesPtr, sOf2xIndexesPtr, sOfHCIndexesPtr, _ := genListsOfHandType()
// 	// _, _, _, _, _, sOf3xIndexesPtr, _, _, _, _ := genListsOfHandType()
// 	fmt.Printf("%T, %T, %T, %T, %T, %T, %T, %T, %T\n", sOfSFIndexesPtr, sOf4xIndexesPtr, sOfFHIndexesPtr, sOfFlIndexesPtr, sOfStIndexesPtr, sOf3xIndexesPtr, sOf2x2IndexesPtr, sOf2xIndexesPtr, sOfHCIndexesPtr)
// 	fmt.Printf("%v\n", sOfSFIndexesPtr)

// 	// searchOut3x(sOf3xIndexesPtr)

// 	// fiveCardHandOrderedListPtr, _ := genOrderedListOfHands()
// 	// fmt.Println("fiveCardHandOrderedListPtr...", fiveCardHandOrderedListPtr[0].handKind)

// 	// sOfTEST := []int{2095432, 2095665, 2061854, 2061871, 2062186, 2062211}
// 	// sOfTESTPtr := &sOfTEST

// 	// writeHandsAndIndexesToFile(sOfSFIndexesPtr, "sfList")
// 	// writeHandsAndIndexesToFile(sOf4xIndexesPtr, "4xList")
// 	// writeHandsAndIndexesToFile(sOfFHIndexesPtr, "fhList")
// 	// writeHandsAndIndexesToFile(sOfFlIndexesPtr, "flList")
// 	// writeHandsAndIndexesToFile(sOfStIndexesPtr, "stList")
// 	// writeHandsAndIndexesToFile(sOf3xIndexesPtr, "3xList")
// 	// writeHandsAndIndexesToFile(sOf2x2IndexesPtr, "2x2List")
// 	// writeHandsAndIndexesToFile(sOf2xIndexesPtr, "2xList")
// 	// writeHandsAndIndexesToFile(sOfHCIndexesPtr, "hcList")

// 	// _ = printSIAsCardStrings(sOfSFIndexesPtr, 40)
// 	// _ = printSIAsCardStrings(sOf4xIndexesPtr, 10)
// 	// _ = printSIAsCardStrings(sOfFHIndexesPtr, 10)
// 	// _ = printSIAsCardStrings(sOfFlIndexesPtr, 10)
// 	// _ = printSIAsCardStrings(sOfStIndexesPtr, 100)
// 	// _ = printSIAsCardStrings(sOf3xIndexesPtr, 100)
// 	// _ = printSIAsCardStrings(sOf2x2IndexesPtr, 10)
// 	// _ = printSIAsCardStrings(sOfTESTPtr, 10)
// 	/*
// 			   	Jc 5s 5d 2h 2d   : 517 : 2095432
// 			   	Jc 5c 5h 4s 4c   : 518 : 2095665

// 			   	Jc 9s 9h 4h 4d   : 115 : 2061854
// 			   Jc 9s 9h 3s 3c   : 116 : 2061871

// 			   Jc 9s 9d 5h 5d   : 151 : 2062186
// 		Jc 9s 9d 4s 4c   : 152 : 2062211
// 	*/

// 	// _ = printSIAsCardStrings(sOf2xIndexesPtr, 10)
// 	// _ = printSIAsCardStrings(sOfHCIndexesPtr, 10)

// 	_ = orderSFIndexesAsc(sOfSFIndexesPtr)
// 	_ = order4xIndexesAsc(sOf4xIndexesPtr)
// 	_ = orderFHIndexesAsc(sOfFHIndexesPtr)
// 	_ = orderFlIndexesAsc(sOfFlIndexesPtr)
// 	_ = orderStIndexesAsc(sOfStIndexesPtr)
// 	_ = order3xIndexesAsc(sOf3xIndexesPtr)
// 	_ = order2x2IndexesAsc(sOf2x2IndexesPtr)
// 	// _ = order2x2IndexesAsc(sOfTESTPtr)
// 	_ = order2xIndexesAsc(sOf2xIndexesPtr)
// 	_ = orderHCIndexesAsc(sOfHCIndexesPtr)

// 	sOfAllIndexesSorted := *sOfHCIndexesPtr
// 	sOfAllIndexesSorted = append(sOfAllIndexesSorted, *sOf2xIndexesPtr...)
// 	sOfAllIndexesSorted = append(sOfAllIndexesSorted, *sOf2x2IndexesPtr...)
// 	sOfAllIndexesSorted = append(sOfAllIndexesSorted, *sOf3xIndexesPtr...)
// 	sOfAllIndexesSorted = append(sOfAllIndexesSorted, *sOfStIndexesPtr...)
// 	sOfAllIndexesSorted = append(sOfAllIndexesSorted, *sOfFlIndexesPtr...)
// 	sOfAllIndexesSorted = append(sOfAllIndexesSorted, *sOfFHIndexesPtr...)
// 	sOfAllIndexesSorted = append(sOfAllIndexesSorted, *sOf4xIndexesPtr...)
// 	sOfAllIndexesSorted = append(sOfAllIndexesSorted, *sOfSFIndexesPtr...)
// 	// sOfAllIndexesSortedPtr := &sOfAllIndexesSorted

// 	// _ = printSIAsCardStrings(sOfSFIndexesPtr, 40)
// 	// _ = printSIAsCardStrings(sOf4xIndexesPtr, 1000)
// 	// _ = printSIAsCardStrings(sOfFHIndexesPtr, 4000)
// 	// _ = printSIAsCardStrings(sOfFlIndexesPtr, 1000)
// 	// _ = printSIAsCardStrings(sOfStIndexesPtr, 11000)
// 	// _ = printSIAsCardStrings(sOf3xIndexesPtr, 10000)
// 	// _ = printSIAsCardStrings(sOf2x2IndexesPtr, 110000)
// 	// _ = printSIAsCardStrings(sOfTESTPtr, 10)
// 	// _ = printSIAsCardStrings(sOf2xIndexesPtr, 500000)
// 	// _ = printSIAsCardStrings(sOfHCIndexesPtr, 200000)

// 	// writeHandsAndIndexesToFile(sOfSFIndexesPtr, "sfListSorted")
// 	// writeHandsAndIndexesToFile(sOf4xIndexesPtr, "4xListSorted")
// 	// writeHandsAndIndexesToFile(sOfFHIndexesPtr, "fhListSorted")
// 	// writeHandsAndIndexesToFile(sOfFlIndexesPtr, "flListSorted")
// 	// writeHandsAndIndexesToFile(sOfStIndexesPtr, "stListSorted")
// 	// writeHandsAndIndexesToFile(sOf3xIndexesPtr, "3xListSorted")
// 	// writeHandsAndIndexesToFile(sOf2x2IndexesPtr, "2x2ListSorted")
// 	// writeHandsAndIndexesToFile(sOf2xIndexesPtr, "2xListSorted")
// 	// writeHandsAndIndexesToFile(sOfHCIndexesPtr, "hcListSorted")
// 	// writeHandsAndIndexesToFile(sOfAllIndexesSortedPtr, "allListSorted")

// 	// Create an ordered listing of hands, after collapsing hands with equivalent values, and
// 	// counting them, perhaps assigning weight based on count.
// 	var sOfEqHandsAll equivalentFCHList
// 	sOfEqHandsAllPtr := &sOfEqHandsAll

// 	var sOfEqFch equivalentFCHList
// 	sOfEqFchPtr := &sOfEqFch
// 	fmt.Println(sOfEqFchPtr)

// 	sOfEqFchPtr = genEquivalentSFHandList(sOfSFIndexesPtr)
// 	sOfEqHandsAll = append(sOfEqHandsAll, sOfEqFch...)
// 	sOfEqFchPtr = genEquivalent4xHandList(sOf4xIndexesPtr)
// 	sOfEqHandsAll = append(sOfEqHandsAll, sOfEqFch...)

// 	fmt.Println(sOfEqHandsAll)

// 	_ = printSOEH(sOfEqHandsAllPtr)

// 	// os.Exit(0)

// 	return nil
// }

// // #########################################################################

// // generateAll5CardIntegerCombos doc line
// func generateAll5CardIntegerCombos(cardSetSize, chooseThisMany int) (ssiPtr *[][]int, numOfEl int, err error) {

// 	// Info.Println(ThisFunc())

// 	n := cardSetSize
// 	k := chooseThisMany

// 	// Keep things artificially sane for now
// 	// n = 7
// 	// k = 5

// 	// Get all possible combinations for k choose n
// 	sl_of_sl_of_ints := combin.Combinations(n, k)
// 	ssiPtr = &sl_of_sl_of_ints

// 	// Find how many of the above there are.
// 	bc := combin.Binomial(n, k)
// 	// This will be the same as len(sl_of_sl_of_ints)

// 	return ssiPtr, bc, nil
// }

// // #########################################################################

// // getCardStringByIndex returns the card value as string, like "As", based on index in an ordered deck. The Ascdh are 0,1,2,3 respectively.
// func getCardStringByIndex(i int) (card manage_table.CardString, err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	if i >= 0 && i <= 51 {
// 		return cardIndexFull[i], nil
// 	} else {
// 		return "no card", errors.New("card index provided is out of range")
// 	}
// }

// // #########################################################################

// // printSSIAsCardStrings maps integers to card names, then prints all combos
// // in the provided Slice of Slices of ints
// func printSSIAsCardStrings(ssiPtr *[][]int) error {

// 	// Info.Printf("%s\n\n", ThisFunc())

// 	ssi := *ssiPtr

// 	for i, _ := range ssi {
// 		for j, _ := range ssi[i] {
// 			c, _ := getCardStringByIndex(ssi[i][j])
// 			fmt.Printf("%v  ", c)
// 		}
// 		fmt.Printf("  : %d\n", i+1)
// 	}
// 	return nil
// }

// // #########################################################################
// // printFclAsStringByIndex prints out the cards in a five card hand pinted to by index number
// // func printFclAsStringBySCLIndex(i int) error {

// // 	Info.Printf("%s\n\n", ThisFunc())

// // 	for _, cPtr := range scl[i] {
// // 		fmt.Printf("%v  ", cPtr)
// // 	}
// // 	fmt.Printf("  : %d\n", i+1)

// // 	return nil
// // }

// // #########################################################################
// // printFclAsStringBySCLIndex takes an index of the hand within a slice of five card hands, and prints the
// // list of cards in that hand as strings.
// func printFclAsStringBySCLIndex(i int) error {

// 	// Info.Printf("%s\n\n", ThisFunc())
// 	// scl := *sOfAllFCLs
// 	// fcl := sOfAllFCLs[i]

// 	for _, cPtr := range sOfAllFCLs[i] {
// 		fmt.Printf("%v%v ", cPtr.Rank, cPtr.Suit)
// 	}

// 	return nil
// }

// // #########################################################################

// // printSIAsCardStrings (print Slice Index as Cards (string).
// // Takes a slice of indexes. For each index value, looks up the 5 card list in the
// // scl (the full listing of all hands), then prints all cards in that hand.
// func printSIAsCardStrings(siPtr *[]int, printUpTo int) error {
// 	// func printSIAsCardStrings(siPtr *[]int, printUpTo int) error {

// 	// printUpTo: if 0, print all
// 	// if > 0, print all entries up to that number entry in the list.
// 	Info.Printf("%s\n\n", debugging.ThisFunc())

// 	for i, fclIndex := range *siPtr {

// 		if printUpTo == 0 || i+1 <= printUpTo {

// 			fclPtr := &sOfAllFCLs[fclIndex]
// 			// for _, cPtr := range fcl {
// 			// 	fmt.Printf("%v%v ", cPtr.Rank, cPtr.Suit)
// 			// }
// 			printFiveCardListAsString(fclPtr)

// 			fmt.Printf("  : %d : %d\n", i+1, fclIndex)
// 		}
// 	}
// 	return nil
// }

// // #########################################################################

// // printSIAsCardStringsNot3x prints the hand if it's a an FH. Specially written jut to catch non 3x hands in the 3x list.
// func printSIAsCardStringsNot3x(siPtr *[]int) error {

// 	// printUpTo: if 0, print all
// 	// if > 0, print all entries up to that number entry in the list.
// 	Info.Printf("%s\n\n", debugging.ThisFunc())

// 	var handName string
// 	handNamePtr := &handName

// 	for fclIndex := range *siPtr {
// 		fcl := sOfAllFCLs[fclIndex]

// 		fclPtr := &fcl

// 		rankCtrPtr := countRanksInfiveCardList(fclPtr)
// 		suitCtrPtr := countSuitsInfiveCardList(fclPtr)

// 		// printFiveCardListAsString(fclPtr)

// 		if checkForFH_5c(fclPtr, 5, rankCtrPtr, suitCtrPtr, handNamePtr) == nil {
// 			printFiveCardListAsString(fclPtr)
// 			fmt.Println()
// 		}

// 	}
// 	return nil
// }

// // #########################################################################

// // printAllHandsInList prints out all the card combos from a supplied slice of fiveCardList
// // func printAllHandsInList(sOfAllFCLs *[]fiveCardList) error {
// func printAllHandsInList() error {

// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// scl := *sOfAllFCLs

// 	// for i, _ := range sOfAllFCLs {
// 	// 	for j, _ := range sOfAllFCLs[i] {
// 	for i := range sOfAllFCLs {
// 		for j := range sOfAllFCLs[i] {
// 			// fmt.Println("comboNum : cardIndex    ", i, j)

// 			fmt.Printf("%v%v ", sOfAllFCLs[i][j].Rank, sOfAllFCLs[i][j].Suit)
// 		}
// 		fmt.Printf("  : %d\n", i+1)
// 	}
// 	return nil
// }

// // #########################################################################

// // genSCLfromSSI generate a Slice of fiveCardList, from Slice of Sices of Integers.
// // Map integers in 5 integer slices, to cards, in 5 card slices.
// // func genSCLfromSSI(ssiPtr *[][]int, dPtr *cardDeck) (sfclPtr *[]fiveCardList, err error) {
// // func genSCLfromSSI(ssiPtr *[][]int, dPtr *manage_table.CardDeck) (sfclPtr *[]fiveCardList, err error) {

// // 	// Info.Printf("%s\n\n", ThisFunc())
// // 	// fmt.Println("len of ssi: ", len(*ssiPtr))
// // 	// fmt.Println()

// // 	ssi := *ssiPtr
// // 	// fmt.Println("ssi: ", ssi)

// // 	// sscp := make([][]*card, len(*ssiPtr))
// // 	sfcl := make([]fiveCardList, len(*ssiPtr))
// // 	sfclPtr = &sfcl

// // 	// fmt.Println("sfclPtr: ", sfclPtr)

// // 	for comboNum := range ssi {
// // 		// fmt.Println("comboNum: ", comboNum)

// // 		var cl = fiveCardList{manage_table.NoCardPtr, manage_table.NoCardPtr, manage_table.NoCardPtr, manage_table.NoCardPtr, manage_table.NoCardPtr}
// // 		sfcl[comboNum] = cl

// // 		for cardIndex, cNum := range ssi[comboNum] {
// // 			// fmt.Println("cardIndex: ", cardIndex, "card Num: ", cNum)
// // 			// sfcl[comboNum][cardIndex] = dPtr.oloptcPtr[ssi[comboNum][cardIndex]]
// // 			cl[cardIndex] = dPtr.oloptcPtr[cNum]
// // 			// fmt.Println("cl[cardIndex]: ", cl[cardIndex])

// // 			// YYY
// // 		}
// // 		// fmt.Println("eo combonum loop")
// // 	}

// // 	return sfclPtr, nil
// // }

// #########################################################################

// CreateFiveCardHandKindRankings maps a major hand type to an integer, representing
// the relative strength of that type, 1 being the best/highest.
// This will allow us to split the comparison and sorting of 5 card hands into 2 sub-tasks,
// 1) figure out the type of hand 1 and hand 2. This will be enough, most often, to know which hand is better.
// 2) If hand 1 and 2 are of the same type, then we must compare/sort within that type.
func CreateFiveCardHandKindRankings() (fchkrFL FiveCardHandKindRankingFullList, err error) {

	// Info.Printf("%s\n\n", ThisFunc())

	// Give ourselves a way to compare 5 card hands, with just integer values for type.
	fchkrFL.SfInfo = fiveCardHandKindRanking{sf, 1}
	fchkrFL.X4Info = fiveCardHandKindRanking{x4, 2}
	fchkrFL.FhInfo = fiveCardHandKindRanking{fh, 3}
	fchkrFL.FlInfo = fiveCardHandKindRanking{fl, 4}
	fchkrFL.StInfo = fiveCardHandKindRanking{st, 5}
	fchkrFL.X3Info = fiveCardHandKindRanking{x3, 6}
	fchkrFL.X22Info = fiveCardHandKindRanking{x22, 7}
	fchkrFL.X2Info = fiveCardHandKindRanking{x2, 8}
	fchkrFL.HcInfo = fiveCardHandKindRanking{hc, 9}

	return fchkrFL, err
}

// // #########################################################################

// // genOrderedListOfHands Not sure what to do with this yet. ...Sorts the sscp/ssi list and gives us the entire ordered listing
// // of all the possible 5 card hands
// func genOrderedListOfHands() (listPtr *[2598960]fiveCardHandMetadata, err error) {

// 	// Info.Printf("%s\n\n", ThisFunc())

// 	var list [2598960]fiveCardHandMetadata
// 	listPtr = &list

// 	return listPtr, nil
// }

// // #####################################################################

// // identifyAllSFs doc line
// func identifyAllSFs() error {

// 	// Info.Printf("%s\n\n", ThisFunc())

// 	var handName string
// 	handNamePtr := &handName

// 	// scl := *sOfAllFCLs

// 	var cl fiveCardList
// 	clPtr := &cl

// 	var i int

// 	var sSfFcl []fiveCardList
// 	sSfFclPtr := &sSfFcl

// 	fmt.Println("len sSfFcl : ", len(sSfFcl))
// 	fmt.Println("len sSfFclPtr : ", len(*sSfFclPtr))

// 	// Will store the
// 	var listOfSFIndexes []int

// 	// os.Exit(0)

// 	// for i, _ := range ssi {
// 	// for i, _ := range sOfAllFCLs {
// 	for i, cl = range sOfAllFCLs {

// 		// fmt.Printf("index: %d\n", i)
// 		// cl = sOfAllFCLs[i]

// 		// Figure out how many cards in list
// 		fiveCardListLength := getfiveCardListLength(clPtr)

// 		// fmt.Printf("  : %d\n", i+1)

// 		// fmt.Println()
// 		// fmt.Println("#### CARD LIST LEN: ", fiveCardListLength)
// 		// fmt.Println()

// 		rankCtrPtr := countRanksInfiveCardList(clPtr)
// 		suitCtrPtr := countSuitsInfiveCardList(clPtr)

// 		// fmt.Println()
// 		// fmt.Println("rankCtrPtr: ", rankCtrPtr)
// 		// fmt.Println("suitCtrPtr: ", suitCtrPtr)
// 		// fmt.Println("dumping the fiveCardList: ", clPtr[0], clPtr[1], clPtr[2], clPtr[3], clPtr[4], clPtr[5], clPtr[6])
// 		// fmt.Println()

// 		_ = fmt.Sprint(clPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr)

// 		if findSFsInList(clPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr, sSfFclPtr) == nil {
// 			// Found a SF
// 			listOfSFIndexes = append(listOfSFIndexes, i)
// 		}
// 	}

// 	for indx, fcl := range sSfFcl {
// 		for j := 0; j < 5; j++ {
// 			fmt.Printf("%v%v ", fcl[j].Rank, fcl[j].Suit)
// 		}
// 		fmt.Printf("  : %d\n", indx+1)
// 	}

// 	fmt.Println()
// 	// fmt.Println("len sSfFcl : ", len(sSfFcl))
// 	fmt.Println("len sSfFclPtr : ", len(*sSfFclPtr))
// 	// fmt.Printf("%v : ", *sSfFclPtr)

// 	os.Exit(0)

// 	// return handName, nil
// 	return nil
// }

// // #####################################################################

// // runCheckForHands exercises the findBestHandInfiveCardList func, to give us some confidence it's doing
// // the right thing.
// func runCheckForHands() (err error) {

// 	var bestHand string

// 	fmt.Println("###################################################################### SF check")
// 	bestHand = findBestHandInfiveCardList(createSFCL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	fmt.Println("###################################################################### 4x check")
// 	bestHand = findBestHandInfiveCardList(create4xCL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	fmt.Println("###################################################################### FH check")
// 	bestHand = findBestHandInfiveCardList(createFHCL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	fmt.Println("###################################################################### Fl check")
// 	bestHand = findBestHandInfiveCardList(createFlCL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	fmt.Println("###################################################################### St check")
// 	bestHand = findBestHandInfiveCardList(createStCL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	fmt.Println("###################################################################### 3x check")
// 	bestHand = findBestHandInfiveCardList(create3xCL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	fmt.Println("###################################################################### 2x2 check")
// 	bestHand = findBestHandInfiveCardList(create2x2CL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	fmt.Println("###################################################################### 2x1 check")
// 	bestHand = findBestHandInfiveCardList(create2x1CL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	fmt.Println("###################################################################### 2x1 check")
// 	bestHand = findBestHandInfiveCardList(createHcCL())
// 	fmt.Println("best hand ALL: ", bestHand)
// 	fmt.Println()

// 	return err
// }

// // #####################################################################

// // genListsOfHandType produces slices of index numbers, pointing at hands of the same type,
// // so, we end up with
// // a slice of ints that are the indexes of all SFs,
// // a slice of ints that are the indexs fo all the 4xs, ...
// func genListsOfHandType() (
// 	sOfSFIndexesPtr,
// 	sOf4xIndexesPtr,
// 	sOfFHIndexesPtr,
// 	sOfFlIndexesPtr,
// 	sOfStIndexesPtr,
// 	sOf3xIndexesPtr,
// 	sOf2x2IndexesPtr,
// 	sOf2xIndexesPtr,
// 	sOfHCIndexesPtr *[]int, err error) {
// 	Info.Printf("%s\n\n", debugging.ThisFunc())

// 	// XXXX

// 	// Flag for making new files, since this operation can be resouce/time consuming.
// 	writeFiles := false

// 	// Give ourselves "buckets" to catch the different types of hands into.
// 	var sOfSFIndexes []int
// 	sOfSFIndexesPtr = &sOfSFIndexes
// 	var sOf4xIndexes []int
// 	sOf4xIndexesPtr = &sOf4xIndexes
// 	var sOfFHIndexes []int
// 	sOfFHIndexesPtr = &sOfFHIndexes
// 	var sOfFlIndexes []int
// 	sOfFlIndexesPtr = &sOfFlIndexes
// 	var sOfStIndexes []int
// 	sOfStIndexesPtr = &sOfStIndexes
// 	var sOf3xIndexes []int
// 	sOf3xIndexesPtr = &sOf3xIndexes
// 	var sOf2x2Indexes []int
// 	sOf2x2IndexesPtr = &sOf2x2Indexes
// 	var sOf2xIndexes []int
// 	sOf2xIndexesPtr = &sOf2xIndexes
// 	var sOfHCIndexes []int
// 	sOfHCIndexesPtr = &sOfHCIndexes

// 	//
// 	var handName string
// 	handNamePtr := &handName

// 	fiveCardListLength := 5

// 	// Create a file for catching index numbers corresponding to "x type" hands.
// 	var all_l_f *os.File

// 	if writeFiles {
// 		all_l_f, err = os.Create("./tmp/allHands")
// 		defer all_l_f.Close()
// 	}

// 	// For each 5 card hand in list, figure out it's type, and keep count.
// 	for i, fcl := range sOfAllFCLs {

// 		fclPtr := &fcl

// 		rankCtrPtr := countRanksInfiveCardList(fclPtr)
// 		suitCtrPtr := countSuitsInfiveCardList(fclPtr)

// 		// printFiveCardListAsString(fclPtr)

// 		// os.Exit(0)
// 		if writeFiles {
// 			// Create a file wth ALL the hands. This can be used later to verify hand/index paring.
// 			sAll, _ := sprintFiveCardListAsString(fclPtr, i, i)
// 			_, _ = all_l_f.WriteString(sAll)
// 		}

// 		switch {
// 		case checkForSF_5c(fclPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 			sOfSFIndexes = append(sOfSFIndexes, i)
// 			counterSF++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = sf_l_f.WriteString(s)
// 			// }

// 		case checkFor4x_5c(fclPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 			sOf4xIndexes = append(sOf4xIndexes, i)
// 			counter4x++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = x4_l_f.WriteString(s)
// 			// }

// 		case checkForFH_5c(fclPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 			sOfFHIndexes = append(sOfFHIndexes, i)
// 			counterFH++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = fh_l_f.WriteString(s)
// 			// }

// 		case checkForFl_5c(fclPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 			sOfFlIndexes = append(sOfFlIndexes, i)
// 			counterFl++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = fl_l_f.WriteString(s)
// 			// }

// 		case checkForSt_5c(fclPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 			sOfStIndexes = append(sOfStIndexes, i)
// 			counterSt++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = st_l_f.WriteString(s)
// 			// }

// 		case checkFor3x_5c(fclPtr, fiveCardListLength, rankCtrPtr, handNamePtr) == nil:
// 			sOf3xIndexes = append(sOf3xIndexes, i)
// 			counter3x++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = x3_l_f.WriteString(s)
// 			// }

// 		case checkFor2x2_5c(fclPtr, fiveCardListLength, rankCtrPtr, handNamePtr) == nil:
// 			sOf2x2Indexes = append(sOf2x2Indexes, i)
// 			counter2x2++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = x2x2_l_f.WriteString(s)
// 			// }

// 		case checkFor2x1_5c(fclPtr, fiveCardListLength, rankCtrPtr, handNamePtr) == nil:
// 			sOf2xIndexes = append(sOf2xIndexes, i)
// 			counter2x++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = x2_l_f.WriteString(s)
// 			// }

// 		// Here we're left with just the HC hands
// 		default:
// 			sOfHCIndexes = append(sOfHCIndexes, i)
// 			// fmt.Println("Foudn a ############################################# HC")
// 			// fmt.Println("added to NO-TYPE list")
// 			counterHC++

// 			// if writeFiles {
// 			// 	s, _ := sprintFiveCardListAsString(fclPtr, i)
// 			// 	_, _ = hc_l_f.WriteString(s)
// 			// }

// 		}

// 		/*
// 			// Quit based on count
// 			if counterHC == 1500000 {
// 				// 1.3M
// 				os.Exit(0)
// 			}
// 		*/
// 	}

// 	// Write the lists to files

// 	// Total all the slice lengths, OF each hand type
// 	total := len(sOfSFIndexes) +
// 		len(sOf4xIndexes) +
// 		len(sOfFHIndexes) +
// 		len(sOfFlIndexes) +
// 		len(sOfStIndexes) +
// 		len(sOf3xIndexes) +
// 		len(sOf2x2Indexes) +
// 		len(sOf2xIndexes) +
// 		len(sOfHCIndexes)

// 	fmt.Println()
// 	fmt.Println()
// 	// Total up all the hand counts, so we can make sure we got them all
// 	countsTotal := counterSF + counter4x + counterFH + counterFl + counterSt + counter3x + counter2x2 + counter2x + counterHC

// 	fmt.Printf("count SF : %v\n", counterSF)
// 	fmt.Printf("count 4x : %v\n", counter4x)
// 	fmt.Printf("count FH : %v\n", counterFH)
// 	fmt.Printf("count Fl : %v\n", counterFl)
// 	fmt.Printf("count St : %v\n", counterSt)
// 	fmt.Printf("count 3x : %v\n", counter3x)
// 	fmt.Printf("count 2x2: %v\n", counter2x2)
// 	fmt.Printf("count 2x : %v\n", counter2x)
// 	fmt.Printf("count Hc : %v\n", counterHC)
// 	fmt.Println("")

// 	fmt.Printf("len sOfSfInsdexe    : %v\t%.3f %%\n", len(sOfSFIndexes), percentOfTotal(len(sOfSFIndexes), total))
// 	fmt.Printf("len sOf4xIndexes    : %v\t%.3f %%\n", len(sOf4xIndexes), percentOfTotal(len(sOf4xIndexes), total))
// 	fmt.Printf("len sOfFHIndexes    : %v\t%.3f %%\n", len(sOfFHIndexes), percentOfTotal(len(sOfFHIndexes), total))
// 	fmt.Printf("len sOfFlIndexes    : %v\t%.3f %%\n", len(sOfFlIndexes), percentOfTotal(len(sOfFlIndexes), total))
// 	fmt.Printf("len sOfStIndexes    : %v\t%.3f %%\n", len(sOfStIndexes), percentOfTotal(len(sOfStIndexes), total))
// 	fmt.Printf("len sOf3xIndexes    : %v\t%.3f %%\n", len(sOf3xIndexes), percentOfTotal(len(sOf3xIndexes), total))
// 	fmt.Printf("len sOf2x2Indexe    : %v\t%.3f %%\n", len(sOf2x2Indexes), percentOfTotal(len(sOf2x2Indexes), total))
// 	fmt.Printf("len sOf2xIndexes    : %v\t%.3f %%\n", len(sOf2xIndexes), percentOfTotal(len(sOf2xIndexes), total))
// 	fmt.Printf("len sOfNoTypeInd    : %v\t%.3f %%\n", len(sOfHCIndexes), percentOfTotal(len(sOfHCIndexes), total))

// 	fmt.Println()
// 	fmt.Println("cnts total          : ", countsTotal)
// 	fmt.Println("len total           : ", total)
// 	fmt.Println("sOfAllFCLs          : ", len(sOfAllFCLs))
// 	fmt.Println()
// 	fmt.Println("diff                : ", len(sOfAllFCLs)-total)
// 	fmt.Println()
// 	fmt.Println()

// 	// Figure out the stats based purely on math, so we can make sure our counts match up / make sense.
// 	findNumOfFlushesIn5cHand()
// 	findNumOfStraightsIn5cHand()
// 	findNumOf3xIn5cHand()
// 	findNumOfFHIn5cHand()
// 	findNumOf2x2In5cHand()
// 	findNumOf2xIn5cHand()
// 	findNumOfHCIn5cHand()
// 	fmt.Println("")

// 	// Sync up the files before closing, deferred.

// 	if writeFiles {
// 		all_l_f.Sync()
// 	}

// 	return sOfSFIndexesPtr, sOf4xIndexesPtr, sOfFHIndexesPtr, sOfFlIndexesPtr, sOfStIndexesPtr, sOf3xIndexesPtr, sOf2x2IndexesPtr, sOf2xIndexesPtr, sOfHCIndexesPtr, nil
// }

// /*
// len sOfSfInsdexe    :        40    0.002 %   correct			   40
// len sOf4xIndexes    :       624    0.024 %   correct			  624
// len sOfFHIndexes    :     3,744    0.144 %   correct			3,744
// len sOfFlIndexes    :     5,108    0.197 %   correct,  			5,108 		5,148 - 40 SF hands
// len sOfStIndexes    :    10,200    0.392 %   correct,  		   10,200       10,240 - 40 SF hands
// len sOf3xIndexes    :    54,912    2.113 %   correct,  		   54,912
// len sOf2x2Indexe    :   123,552    4.754 %   correct,		  123,552
// len sOf2xIndexes    : 1,098,240   42.257 %
// len sOfNoTypeInd    : 1,302,540   50.118 %

// len total           :  2,598,960
// sfclPtr             :  2,598,960
// */

// // #####################################################################

// // writeHandsAndIndexesToFile takes a slice of indexes of hands, and writes the corresponding hands to a file ./tmp
// func writeHandsAndIndexesToFile(sOfIndexesPtr *[]int, fileName string) error {
// 	Info.Printf("%s\n\n", debugging.ThisFunc())

// 	fPath := "./tmp/" + fileName

// 	// Create a file handle
// 	fhPtr, err := os.Create(fPath)
// 	defer fhPtr.Close()

// 	for index, v := range *sOfIndexesPtr {

// 		fcl := sOfAllFCLs[v]
// 		fclPtr := &fcl

// 		s, _ := sprintFiveCardListAsString(fclPtr, index, v)
// 		_, _ = fhPtr.WriteString(s)
// 	}

// 	fhPtr.Sync()
// 	return err
// }

// // #####################################################################

// // findNumOfFlushesIn5cHand doc line
// func findNumOfFlushesIn5cHand() {

// 	// 4 suits * 13 choose 5

// 	bc := combin.Binomial(13, 5)
// 	fmt.Println("Total num of flushes in a 5 card hands is ", bc*4)
// }

// // #####################################################################

// // findNumOfStraightsIn5cHand doc line
// func findNumOfStraightsIn5cHand() {

// 	// 10 possible botom ranks, A, 2, 3, 4 ... T * 4 card of that rank(r0) * 4 of r+1 * 4 r+2 * 4 r+3 * 4 r+4
// 	fmt.Println("Total num of straights in a 5 card hands is ", 10*4*4*4*4*4)
// }

// // #####################################################################

// // findNumOf3xIn5cHand
// func findNumOf3xIn5cHand() {

// 	// 13 ranks * 4c3 * 12c2 * 4 * 4
// 	//
// 	bc := combin.Binomial(4, 3)
// 	bc2 := combin.Binomial(12, 2)

// 	fmt.Println("Total num of 3x in a 5 card hands is ", 13*bc*bc2*16)
// }

// // #####################################################################

// // findNumOfFHIn5cHand
// func findNumOfFHIn5cHand() {

// 	// 13 * 4c3 * 12 * 4c2
// 	fmt.Println("Total num of FHs in a 5 card hands is ", 4*13*12*6)
// }

// // #####################################################################

// // findNumOf2x2In5cHand
// func findNumOf2x2In5cHand() {

// 	// 13c2 * 6 * 6 * 44
// 	bc := combin.Binomial(13, 2)
// 	fmt.Println("Total num of 2x2 in a 5 card hands is ", bc*6*6*44)
// }

// // #####################################################################

// // findNumOf2xIn5cHand
// func findNumOf2xIn5cHand() {

// 	// 13 * 4c2 * 12c3 * 4 * 4 * 4

// 	// bc := combin.Binomial(12, 3)
// 	// fmt.Println("Total num of 2x in a 5 card hands is ", 13*6*bc*4*4*4)

// 	// [ 4c2 * 13 ranks ]  * [ 12 ranks c 3 * 4 cards per rank ]

// 	bc := combin.Binomial(12, 3)
// 	fmt.Println("Total num of 2x in a 5 card hands is ", 6*13*bc*4*4*4)
// }

// // #####################################################################

// // findNumOfHCIn5cHand
// func findNumOfHCIn5cHand() {

// 	// 13 * 4 * 12 * 4 * 11 * 4 * 10 * 4 * 9 * 4 -#SF -#St -#Fl
// 	fmt.Println("Total num of HC in a 5 card hands is ", 2598960-40-624-3744-5108-10200-54912-123552-1098240)
// }

// // #####################################################################

// // searchOut3x is used to verify presence of 3x hands in any list of indexes refrencing the full set of 5 card hands
// func searchOut3x(sOfIndexToCPtr *[]int) {

// 	count := 0
// 	sOfIndexToC := *sOfIndexToCPtr

// 	f, _ := os.Create("./tmp/x3list")
// 	// check(err)
// 	defer f.Close()

// 	for _, v := range sOfIndexToC {

// 		c1r := sOfAllFCLs[v][0].Rank
// 		c2r := sOfAllFCLs[v][1].Rank
// 		c3r := sOfAllFCLs[v][2].Rank
// 		c4r := sOfAllFCLs[v][3].Rank
// 		c5r := sOfAllFCLs[v][4].Rank

// 		c1s := sOfAllFCLs[v][0].Suit
// 		c2s := sOfAllFCLs[v][1].Suit
// 		c3s := sOfAllFCLs[v][2].Suit
// 		c4s := sOfAllFCLs[v][3].Suit
// 		c5s := sOfAllFCLs[v][4].Suit

// 		if (c1r == c2r && c1r == c3r) ||
// 			(c1r == c2r && c1r == c4r) ||
// 			(c1r == c2r && c1r == c5r) ||

// 			(c1r == c3r && c1r == c4r) ||
// 			(c1r == c3r && c1r == c5r) ||

// 			(c1r == c4r && c1r == c5r) ||

// 			(c1r == c3r && c1r == c5r) {

// 			// fmt.Println("Found 3x hand : ", c1r, c1s, " ", c2r, c2s, " ", c3r, c3s, " ", c4r, c4s, " ", c5r, c5s, " : ", v)
// 			count++

// 			// fmt.Println("index / value : ", i, v)

// 			// n3, err := f.WriteString("writes\n")
// 			_, _ = f.WriteString(fmt.Sprintln(c1r, c1s, " ", c2r, c2s, " ", c3r, c3s, " ", c4r, c4s, " ", c5r, c5s, " : ", v))
// 			// os.Exit(0)
// 		}

// 	}
// 	fmt.Println("3x hand count : ", count)

// 	f.Sync()

// }

// // #####################################################################

// // printFiveCardListAsString takes a five card hand Ptr and prints out the 5 cards as text
// func printFiveCardListAsString(fclPtr *fiveCardList) (err error) {

// 	fcl := *fclPtr

// 	// Show order pre sort
// 	for j := 0; j <= 4; j++ {
// 		fmt.Print(fcl[j].Rank, fcl[j].Suit, " ")
// 	}
// 	// fmt.Println()

// 	return nil
// }

// // #####################################################################

// // sprintFiveCardListAsString takes a five card hand Ptr, and an integer (usually the hands index), and prints out the 5
// // cards, and the integer, as text.
// func sprintFiveCardListAsString(fclPtr *fiveCardList, index1, index2 int) (s string, err error) {

// 	fcl := *fclPtr

// 	// Show order pre sort
// 	for j := 0; j <= 4; j++ {
// 		s += fmt.Sprint(fcl[j].Rank, fcl[j].Suit, " ")
// 	}
// 	s += ": " + strconv.Itoa(index1) + " : " + strconv.Itoa(index2) + "\n"

// 	return s, nil

// }

// // #####################################################################

// // percentOfTotal doc line
// func percentOfTotal(value, total int) (f float64) {

// 	v64 := float64(value)
// 	t64 := float64(total)

// 	f = (v64 * 100 / t64)

// 	return f
// }

// // #####################################################################

// // findBestHandInfiveCardList examines the hand, and looks for hand rank, top to bottom, quits as soon as it finds one.
// // By convention, index 0 and 1 are the hole cards.
// // All functions used herein, use countRanksInfiveCardList and countSuitsInfiveCardList to as support functions
// //
// // In this release / for now, we'll ONLY consider card lists up to 7 cards long, 2 hole cards and 5 community cards
// // In each case, we're looking for the best/top hand of it's type.
// func findBestHandInfiveCardList(clPtr *fiveCardList) (handName string) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	var handNamePtr *string
// 	handNamePtr = &handName

// 	cl := *clPtr

// 	// Figure out how many cards in list
// 	fiveCardListLength := getfiveCardListLength(clPtr)

// 	fmt.Println()
// 	fmt.Println("#### CARD LIST LEN: ", fiveCardListLength)
// 	fmt.Println()

// 	rankCtrPtr := countRanksInfiveCardList(clPtr)
// 	suitCtrPtr := countSuitsInfiveCardList(clPtr)

// 	fmt.Println()
// 	fmt.Println("randCtrPtr: ", rankCtrPtr)
// 	fmt.Println("suitCtrPtr: ", suitCtrPtr)
// 	// fmt.Println("dumping the fiveCardList: ", clPtr[0], clPtr[1], clPtr[2], clPtr[3], clPtr[4], clPtr[5], clPtr[6])
// 	fmt.Println("dumping the fiveCardList: ", cl[0], cl[1], cl[2], cl[3], cl[4])
// 	fmt.Println()

// 	switch {
// 	case checkForSF_5c(clPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 	case checkFor4x_5c(clPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 	case checkForFH_5c(clPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 	case checkForFl_5c(clPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 	case checkForSt_5c(clPtr, fiveCardListLength, rankCtrPtr, suitCtrPtr, handNamePtr) == nil:
// 	case checkFor3x_5c(clPtr, fiveCardListLength, rankCtrPtr, handNamePtr) == nil:
// 	case checkFor2x2_5c(clPtr, fiveCardListLength, rankCtrPtr, handNamePtr) == nil:
// 	case checkFor2x1_5c(clPtr, fiveCardListLength, rankCtrPtr, handNamePtr) == nil:
// 	case checkForHc_5c(clPtr, fiveCardListLength, rankCtrPtr, handNamePtr) == nil:
// 	default:
// 		*handNamePtr = "default case"
// 	}
// 	return handName
// }

// // #####################################################################

// // getfiveCardListLength doc line
// func getfiveCardListLength(clPtr *fiveCardList) (l int) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	cl := *clPtr

// 	// Figure out how many cards in list are defined / have been assigned
// 	l = 0

// 	for _, cPtr := range cl {
// 		if cPtr != nil {
// 			l++
// 		}
// 	}
// 	return l
// }

// // #####################################################################

// // countRanksInfiveCardList updates the rcm with the number of cards of each rank in the card list.
// func countRanksInfiveCardList(clPtr *fiveCardList) (rcPtr *rankCounter) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// fmt.Println("clPtr0: ", clPtr[0])

// 	cl := *clPtr

// 	var myRcm rankCtrMap
// 	myRcm = make(rankCtrMap)
// 	var rc rankCounter
// 	rcPtr = &rc
// 	rcPtr.rcm = myRcm

// 	rc.max = 0
// 	rc.uniqeRankCt = 0

// 	// We don't calculate these values here, as it would be wasteful, since not all will be needed all the time,
// 	// but, some will be needed sometimes.
// 	rc.top1x1 = manage_table.RX // give it initial "unset" value
// 	rc.top1x2 = manage_table.RX
// 	rc.top1x3 = manage_table.RX
// 	rc.top1x4 = manage_table.RX
// 	rc.top1x5 = manage_table.RX
// 	rc.top2x1 = manage_table.RX
// 	rc.top2x2 = manage_table.RX
// 	rc.top2x3 = manage_table.RX
// 	rc.top3x1 = manage_table.RX
// 	rc.top3x2 = manage_table.RX
// 	rc.top4x1 = manage_table.RX

// 	for _, cPtr := range cl {

// 		// work only on the defined cards in the list
// 		if cPtr != nil {
// 			// fmt.Printf("%v\n", *cPtr)
// 			rcPtr.rcm[cPtr.Rank]++

// 			// fmt.Printf("c.Rank: %v; count: %v\n", cPtr.Rank, rcPtr.rcm[cPtr.Rank])

// 			if rcPtr.rcm[cPtr.Rank] > rcPtr.max {
// 				rcPtr.max = rcPtr.rcm[cPtr.Rank]
// 			}
// 		}
// 		// fmt.Println("max: ", rcPtr.max)
// 	}

// 	// Count up how many different unique ranks there are in the list
// 	for _, rank := range manage_table.RankList {
// 		if rcPtr.rcm[rank] > 0 {
// 			rcPtr.uniqeRankCt++
// 		}
// 	}

// 	// fmt.Println()
// 	// fmt.Println("Counts:")
// 	// fmt.Printf("rcPtr %v: \n", rcPtr)

// 	return rcPtr
// }

// // #####################################################################

// // countSuitsInfiveCardList counts the cards in the 4 souits
// func countSuitsInfiveCardList(clPtr *fiveCardList) (scPtr *suitCounter) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// fmt.Println("clPtr0: ", clPtr[0])

// 	cl := *clPtr

// 	var myScm suitCtrMap
// 	myScm = make(suitCtrMap)
// 	var sc suitCounter
// 	scPtr = &sc
// 	scPtr.scm = myScm

// 	for _, cPtr := range cl {

// 		// work only on the defined cards in the list
// 		if cPtr != nil {
// 			// fmt.Printf("%v\n", *cPtr)
// 			scPtr.scm[cPtr.Suit]++

// 			// fmt.Printf("cPtr.Suit: %v; count: %v\n", cPtr.Suit, scPtr.scm[cPtr.Suit])

// 			if scPtr.scm[cPtr.Suit] > scPtr.max {
// 				scPtr.max = scPtr.scm[cPtr.Suit]
// 			}
// 		}
// 		// fmt.Println("max: ", scPtr.max)
// 	}

// 	// fmt.Println("Counts:")
// 	// fmt.Println("scPtr[s]: ", scPtr.scm[s])
// 	// fmt.Println("scPtr[c]: ", scPtr.scm[c])
// 	// fmt.Println("scPtr[h]: ", scPtr.scm[h])
// 	// fmt.Println("scPtr[d]: ", scPtr.scm[d])
// 	// fmt.Println("scPtr.max: ", scPtr.max)
// 	// fmt.Println()

// 	return scPtr
// }

// // #####################################################################

// // createClPtr - doc line
// func createClPtr() (clPtr *fiveCardList) {

// 	cl := make(fiveCardList, 5)
// 	clPtr = &cl

// 	return clPtr
// }

// // #####################################################################

// // find5OrMoreOfSameSuitInfiveCardList takes a card list and returns true, if there were
// // 5 or more cards of any one suit in it, and the list of cards of that suit.
// // Otherwise assigns creates an error.
// func find5OrMoreOfSameSuitInfiveCardList(fclPtr *fiveCardList, cll int, scPtr *suitCounter) (resultingFclPtr *fiveCardList, topSuit manage_table.CardSuit, err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	fcl := *fclPtr
// 	// fmt.Println("in find5OrMoreOfSameSuitInfiveCardList; cl[...]: ", cl[0], cl[1])

// 	resultingFclPtr = createClPtr()
// 	resultingFcl := *resultingFclPtr
// 	// fmt.Println("resultingFcl: ",Fcl)
// 	// fmt.Println("scPtr.max: ", scPtr.max, "cll: ", cll)

// 	// Run prelim checks
// 	if scPtr.max < 5 || cll < 5 {
// 		err = errors.New("Failed prelim checks, returning emtpy resultingFclPtr")
// 		return resultingFclPtr, manage_table.X, err
// 	}

// 	// Figure out which suit is most common in the list.
// 	topSuit = manage_table.X
// 	for _, cs := range manage_table.SuitList {
// 		// fmt.Println("in find5OrMoreOfSameSuitInfiveCardList; cs: ", cs)
// 		// Info.Printf("%s; cs: %v\n\n", ThisFunc(), cs)
// 		if scPtr.scm[cs] > 4 {
// 			topSuit = cs
// 		}
// 	}
// 	// fmt.Println("top suit: ", topSuit)

// 	// os.Exit(0)

// 	// Grab just the cards of the topSuit
// 	i := 0
// 	for _, cPtr := range fcl {
// 		// fmt.Println("i: ", i)
// 		// fmt.Println("cPtr.Suit: ", cPtr.Suit)

// 		if cPtr.Suit == topSuit {
// 			// fmt.Println("i inside if: ", i)
// 			resultingFcl[i] = cPtr
// 			i++
// 		}
// 	}

// 	// fmt.Println("resulting fcl: ", resultingFcl)
// 	// printFiveCardListAsString(resultingFclPtr)
// 	// fmt.Println("err: ", err)
// 	// os.Exit(0)

// 	return resultingFclPtr, topSuit, err
// }

// // #####################################################################

// // orderCardsOfSameSuit takes 5 to 7 cards of the same suit, and returns a 5 card ordered list of ranks, high to low.
// // The significance of "same suit" in this context is that there can be no 2 or more cards of the same rank.
// //
// //	Checks for a 5 high straight as well,
// //
// // to differentiate between an arbitrary A high flush and a 5 high Straight Flush, in which case it will return the A as the last card, not first.
// func orderCardsOfSameSuit(clPtr *fiveCardList, cs manage_table.CardSuit) (manage_table.CardRank, manage_table.CardRank, manage_table.CardRank, manage_table.CardRank, manage_table.CardRank) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	cl := *clPtr

// 	var c manage_table.Card

// 	var myFlCard1, myFlCard2, myFlCard3, myFlCard4, myFlCard5, myFlCard6, myFlCard7 manage_table.Card
// 	myFlCard1 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard2 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard3 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard4 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard5 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard6 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard7 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}

// 	for _, cPtr := range cl {

// 		c = *cPtr

// 		if c.Suit == cs { // if we found a card of requested suit...

// 			if myFlCard1.Rank == manage_table.RX { // 1st card in the suit of the Fl
// 				myFlCard1 = *cPtr

// 			} else if myFlCard2.Rank == manage_table.RX { // 2nd card in the suit of the Fl
// 				if crm[c.Rank] > crm[myFlCard1.Rank] { // if new card is higher than the the first
// 					myFlCard2 = myFlCard1 // move the first to the second position
// 					myFlCard1 = c         // Set the first/top pos to c
// 				} else {
// 					myFlCard2 = c
// 				}

// 			} else if myFlCard3.Rank == manage_table.RX { // 3rd card in the suit of the Fl
// 				if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 2
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = myFlCard1
// 					myFlCard1 = c
// 				} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = c
// 				} else {
// 					myFlCard3 = c
// 				}

// 			} else if myFlCard4.Rank == manage_table.RX {
// 				if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 3
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = myFlCard1
// 					myFlCard1 = c
// 				} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = c
// 				} else if crm[c.Rank] > crm[myFlCard3.Rank] {
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = c
// 				} else {
// 					myFlCard4 = c
// 				}

// 			} else if myFlCard5.Rank == manage_table.RX {
// 				if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 4
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = myFlCard1
// 					myFlCard1 = c
// 				} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = c
// 				} else if crm[c.Rank] > crm[myFlCard3.Rank] {
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = c
// 				} else if crm[c.Rank] > crm[myFlCard4.Rank] {
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = c
// 				} else {
// 					myFlCard5 = c
// 				}

// 			} else if myFlCard6.Rank == manage_table.RX {
// 				if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 5
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = myFlCard1
// 					myFlCard1 = c
// 				} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = c
// 				} else if crm[c.Rank] > crm[myFlCard3.Rank] {
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = c
// 				} else if crm[c.Rank] > crm[myFlCard4.Rank] {
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = c
// 				} else if crm[c.Rank] > crm[myFlCard5.Rank] {
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = c
// 				} else {
// 					myFlCard6 = c
// 				}

// 			} else if myFlCard7.Rank == manage_table.RX {
// 				if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 6
// 					myFlCard7 = myFlCard6
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = myFlCard1
// 					myFlCard1 = c
// 				} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 					myFlCard7 = myFlCard6
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = myFlCard2
// 					myFlCard2 = c
// 				} else if crm[c.Rank] > crm[myFlCard3.Rank] {
// 					myFlCard7 = myFlCard6
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = myFlCard3
// 					myFlCard3 = c
// 				} else if crm[c.Rank] > crm[myFlCard4.Rank] {
// 					myFlCard7 = myFlCard6
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = myFlCard4
// 					myFlCard4 = c
// 				} else if crm[c.Rank] > crm[myFlCard5.Rank] {
// 					myFlCard7 = myFlCard6
// 					myFlCard6 = myFlCard5
// 					myFlCard5 = c
// 				} else if crm[c.Rank] > crm[myFlCard6.Rank] {
// 					myFlCard7 = myFlCard6
// 					myFlCard6 = c
// 				} else {
// 					myFlCard7 = c
// 				}
// 			}
// 		}
// 	}
// 	return myFlCard1.Rank, myFlCard2.Rank, myFlCard3.Rank, myFlCard4.Rank, myFlCard5.Rank
// }

// // #####################################################################

// // orderCardsOfSameSuit2 takes 5 to 7 cards of the same suit, and returns a 5 card ordered list of ranks, high to low.
// // The significance of "same suit" in this context is that there can be no 2 cards of the same rank.
// //
// // Checks for a 5 high straight as well,
// // to differentiate between an arbitrary A high flush and a 5 high Straight Flush, in which case it will return the
// // A as the last manage_table.Card, not first.
// func orderCardsOfSameSuit2(clPtr *fiveCardList, cs manage_table.CardSuit) (resultingClPtr *fiveCardList, err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	cl := *clPtr

// 	// fmt.Println("in orderCardsOfSameSuit2; clPtr: ", clPtr)

// 	resultingClPtr = createClPtr()
// 	resultingCl := *resultingClPtr

// 	var c manage_table.Card

// 	var myFlCard1, myFlCard2, myFlCard3, myFlCard4, myFlCard5, myFlCard6, myFlCard7 manage_table.Card
// 	myFlCard1 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard2 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard3 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard4 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard5 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard6 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myFlCard7 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}

// 	myFlCard1Ptr := &myFlCard1
// 	myFlCard2Ptr := &myFlCard2
// 	myFlCard3Ptr := &myFlCard3
// 	myFlCard4Ptr := &myFlCard4
// 	myFlCard5Ptr := &myFlCard5
// 	// myFlCard6Ptr := &myFlCard6
// 	// myFlCard7Ptr := &myFlCard7

// 	for _, cPtr := range cl {

// 		if cPtr != nil {
// 			c = *cPtr
// 			// fmt.Println("in orderCardsOfSameSuit2, manage_table.Card from range: ", c)

// 			if c.Suit == cs { // if we found a card of requested suit...

// 				if myFlCard1.Rank == manage_table.RX { // 1st card in the suit of the Fl
// 					myFlCard1 = *cPtr

// 				} else if myFlCard2.Rank == manage_table.RX { // 2nd card in the suit of the Fl
// 					if crm[c.Rank] > crm[myFlCard1.Rank] { // if new card is higher than the the first
// 						myFlCard2 = myFlCard1 // move the first to the second position
// 						myFlCard1 = c         // Set the first/top pos to c
// 					} else {
// 						myFlCard2 = c
// 					}

// 				} else if myFlCard3.Rank == manage_table.RX { // 3rd card in the suit of the Fl
// 					if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 2
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = myFlCard1
// 						myFlCard1 = c
// 					} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = c
// 					} else {
// 						myFlCard3 = c
// 					}

// 				} else if myFlCard4.Rank == manage_table.RX {
// 					if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 3
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = myFlCard1
// 						myFlCard1 = c
// 					} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = c
// 					} else if crm[c.Rank] > crm[myFlCard3.Rank] {
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = c
// 					} else {
// 						myFlCard4 = c
// 					}

// 				} else if myFlCard5.Rank == manage_table.RX {
// 					if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 4
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = myFlCard1
// 						myFlCard1 = c
// 					} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = c
// 					} else if crm[c.Rank] > crm[myFlCard3.Rank] {
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = c
// 					} else if crm[c.Rank] > crm[myFlCard4.Rank] {
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = c
// 					} else {
// 						myFlCard5 = c
// 					}

// 				} else if myFlCard6.Rank == manage_table.RX {
// 					if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 5
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = myFlCard1
// 						myFlCard1 = c
// 					} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = c
// 					} else if crm[c.Rank] > crm[myFlCard3.Rank] {
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = c
// 					} else if crm[c.Rank] > crm[myFlCard4.Rank] {
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = c
// 					} else if crm[c.Rank] > crm[myFlCard5.Rank] {
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = c
// 					} else {
// 						myFlCard6 = c
// 					}

// 				} else if myFlCard7.Rank == manage_table.RX {
// 					if crm[c.Rank] > crm[myFlCard1.Rank] { // higher than the other 6
// 						myFlCard7 = myFlCard6
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = myFlCard1
// 						myFlCard1 = c
// 					} else if crm[c.Rank] > crm[myFlCard2.Rank] {
// 						myFlCard7 = myFlCard6
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = myFlCard2
// 						myFlCard2 = c
// 					} else if crm[c.Rank] > crm[myFlCard3.Rank] {
// 						myFlCard7 = myFlCard6
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = myFlCard3
// 						myFlCard3 = c
// 					} else if crm[c.Rank] > crm[myFlCard4.Rank] {
// 						myFlCard7 = myFlCard6
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = myFlCard4
// 						myFlCard4 = c
// 					} else if crm[c.Rank] > crm[myFlCard5.Rank] {
// 						myFlCard7 = myFlCard6
// 						myFlCard6 = myFlCard5
// 						myFlCard5 = c
// 					} else if crm[c.Rank] > crm[myFlCard6.Rank] {
// 						myFlCard7 = myFlCard6
// 						myFlCard6 = c
// 					} else {
// 						myFlCard7 = c
// 					}
// 				}
// 			}
// 		}
// 	}

// 	resultingCl[0] = myFlCard1Ptr
// 	resultingCl[1] = myFlCard2Ptr
// 	resultingCl[2] = myFlCard3Ptr
// 	resultingCl[3] = myFlCard4Ptr
// 	resultingCl[4] = myFlCard5Ptr

// 	// switch {
// 	// case myFlCard6.Rank != manage_table.RX:
// 	// 	resultingClPtr[5] = myFlCard6Ptr
// 	// case myFlCard7.Rank != manage_table.RX:
// 	// 	resultingClPtr[6] = myFlCard7Ptr
// 	// }

// 	// fmt.Println("dumping the resultingClPtr: ", resultingClPtr[0], resultingClPtr[1], resultingClPtr[2], resultingClPtr[3], resultingClPtr[4], resultingClPtr[5], resultingClPtr[6])
// 	// fmt.Println("dumping the resultingClPtr: ", resultingClPtr[0], resultingClPtr[1], resultingClPtr[2], resultingClPtr[3], resultingClPtr[4])

// 	/*
// 		for i, _ := range resultingClPtr {
// 			fmt.Printf("%v%v ", resultingClPtr[i].Rank, resultingClPtr[i].Suit)
// 		}
// 		fmt.Printf("\n")
// 	*/

// 	return resultingClPtr, err
// }

// // #####################################################################

// // orderCardsOfMixedSuit takes 5 to 7 cards, irrespective of suit, and orders them by rank.
// // Supports checkFor2x2_5c, find2x, and checkForHc_5c, only.
// // Also, fills in clPtr values, when possible, for the 2x... vars.
// func orderCardsOfMixedSuit(clPtr *fiveCardList) (resultingClPtr *fiveCardList, err error) {
// 	// fmt.Println("### Starting orderCardsOfMixedSuit ###")
// 	// fmt.Println("in orderCardsOfMixedSuit; clPtr: ", clPtr)

// 	cl := *clPtr

// 	resultingClPtr = createClPtr()
// 	resultingCl := *resultingClPtr

// 	var c manage_table.Card

// 	// var myCard1, myCard2, myCard3, myCard4, myCard5, myCard6, myCard7 manage_table.Card
// 	var myCard1, myCard2, myCard3, myCard4, myCard5 manage_table.Card
// 	myCard1 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myCard2 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myCard3 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myCard4 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	myCard5 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	// myCard6 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}
// 	// myCard7 = manage_table.Card{manage_table.RX, manage_table.X, false, false, false, false, 0}

// 	myCard1Ptr := &myCard1
// 	myCard2Ptr := &myCard2
// 	myCard3Ptr := &myCard3
// 	myCard4Ptr := &myCard4
// 	myCard5Ptr := &myCard5
// 	// myCard6Ptr := &myCard6
// 	// myCard7Ptr := &myCard7

// 	for _, cPtr := range cl {

// 		if cPtr != nil {
// 			c = *cPtr
// 			// fmt.Println("in orderCardsOfMixedSuit, manage_table.Card from range: ", c)

// 			if myCard1.Rank == manage_table.RX { // 1st card in the suit of the Fl
// 				myCard1 = *cPtr

// 			} else if myCard2.Rank == manage_table.RX { // 2nd card in the suit of the Fl
// 				if crm[c.Rank] > crm[myCard1.Rank] { // if new card is higher than the the first
// 					myCard2 = myCard1 // move the first to the second position
// 					myCard1 = c       // Set the first/top pos to c
// 				} else {
// 					myCard2 = c
// 				}

// 			} else if myCard3.Rank == manage_table.RX { // 3rd card in the suit of the Fl
// 				if crm[c.Rank] > crm[myCard1.Rank] { // higher than the other 2
// 					myCard3 = myCard2
// 					myCard2 = myCard1
// 					myCard1 = c
// 				} else if crm[c.Rank] > crm[myCard2.Rank] {
// 					myCard3 = myCard2
// 					myCard2 = c
// 				} else {
// 					myCard3 = c
// 				}

// 			} else if myCard4.Rank == manage_table.RX {
// 				if crm[c.Rank] > crm[myCard1.Rank] { // higher than the other 3
// 					myCard4 = myCard3
// 					myCard3 = myCard2
// 					myCard2 = myCard1
// 					myCard1 = c
// 				} else if crm[c.Rank] > crm[myCard2.Rank] {
// 					myCard4 = myCard3
// 					myCard3 = myCard2
// 					myCard2 = c
// 				} else if crm[c.Rank] > crm[myCard3.Rank] {
// 					myCard4 = myCard3
// 					myCard3 = c
// 				} else {
// 					myCard4 = c
// 				}

// 			} else if myCard5.Rank == manage_table.RX {
// 				if crm[c.Rank] > crm[myCard1.Rank] { // higher than the other 4
// 					myCard5 = myCard4
// 					myCard4 = myCard3
// 					myCard3 = myCard2
// 					myCard2 = myCard1
// 					myCard1 = c
// 				} else if crm[c.Rank] > crm[myCard2.Rank] {
// 					myCard5 = myCard4
// 					myCard4 = myCard3
// 					myCard3 = myCard2
// 					myCard2 = c
// 				} else if crm[c.Rank] > crm[myCard3.Rank] {
// 					myCard5 = myCard4
// 					myCard4 = myCard3
// 					myCard3 = c
// 				} else if crm[c.Rank] > crm[myCard4.Rank] {
// 					myCard5 = myCard4
// 					myCard4 = c
// 				} else {
// 					myCard5 = c
// 				}

// 				/*
// 					} else if myCard6.Rank == manage_table.RX {
// 						if crm[c.Rank] > crm[myCard1.Rank] { // higher than the other 5
// 							myCard6 = myCard5
// 							myCard5 = myCard4
// 							myCard4 = myCard3
// 							myCard3 = myCard2
// 							myCard2 = myCard1
// 							myCard1 = c
// 						} else if crm[c.Rank] > crm[myCard2.Rank] {
// 							myCard6 = myCard5
// 							myCard5 = myCard4
// 							myCard4 = myCard3
// 							myCard3 = myCard2
// 							myCard2 = c
// 						} else if crm[c.Rank] > crm[myCard3.Rank] {
// 							myCard6 = myCard5
// 							myCard5 = myCard4
// 							myCard4 = myCard3
// 							myCard3 = c
// 						} else if crm[c.Rank] > crm[myCard4.Rank] {
// 							myCard6 = myCard5
// 							myCard5 = myCard4
// 							myCard4 = c
// 						} else if crm[c.Rank] > crm[myCard5.Rank] {
// 							myCard6 = myCard5
// 							myCard5 = c
// 						} else {
// 							myCard6 = c
// 						}

// 					} else if myCard7.Rank == manage_table.RX {
// 						if crm[c.Rank] > crm[myCard1.Rank] { // higher than the other 6
// 							myCard7 = myCard6
// 							myCard6 = myCard5
// 							myCard5 = myCard4
// 							myCard4 = myCard3
// 							myCard3 = myCard2
// 							myCard2 = myCard1
// 							myCard1 = c
// 						} else if crm[c.Rank] > crm[myCard2.Rank] {
// 							myCard7 = myCard6
// 							myCard6 = myCard5
// 							myCard5 = myCard4
// 							myCard4 = myCard3
// 							myCard3 = myCard2
// 							myCard2 = c
// 						} else if crm[c.Rank] > crm[myCard3.Rank] {
// 							myCard7 = myCard6
// 							myCard6 = myCard5
// 							myCard5 = myCard4
// 							myCard4 = myCard3
// 							myCard3 = c
// 						} else if crm[c.Rank] > crm[myCard4.Rank] {
// 							myCard7 = myCard6
// 							myCard6 = myCard5
// 							myCard5 = myCard4
// 							myCard4 = c
// 						} else if crm[c.Rank] > crm[myCard5.Rank] {
// 							myCard7 = myCard6
// 							myCard6 = myCard5
// 							myCard5 = c
// 						} else if crm[c.Rank] > crm[myCard6.Rank] {
// 							myCard7 = myCard6
// 							myCard6 = c
// 						} else {
// 							myCard7 = c
// 						}

// 				*/
// 			}
// 		}
// 	}

// 	resultingCl[0] = myCard1Ptr
// 	resultingCl[1] = myCard2Ptr
// 	resultingCl[2] = myCard3Ptr
// 	resultingCl[3] = myCard4Ptr
// 	resultingCl[4] = myCard5Ptr

// 	// switch {
// 	// case myCard6.Rank != manage_table.RX:
// 	// 	resultingClPtr[5] = myCard6Ptr
// 	// case myCard7.Rank != manage_table.RX:
// 	// 	resultingClPtr[6] = myCard7Ptr
// 	// }

// 	// fmt.Println("dumping the resultingClPtr: ", resultingClPtr[0], resultingClPtr[1], resultingClPtr[2], resultingClPtr[3], resultingClPtr[4], resultingClPtr[5], resultingClPtr[6])
// 	// fmt.Println("dumping the resultingClPtr: ", resultingClPtr[0], resultingClPtr[1], resultingClPtr[2], resultingClPtr[3], resultingClPtr[4])

// 	return resultingClPtr, err
// }

// // #####################################################################

// /*
// #####################################################################

// ######## #### ##    ## ########     ##     ##    ###    ##    ## ########     ######## ##    ## ########  ########
// ##        ##  ###   ## ##     ##    ##     ##   ## ##   ###   ## ##     ##       ##     ##  ##  ##     ## ##
// ##        ##  ####  ## ##     ##    ##     ##  ##   ##  ####  ## ##     ##       ##      ####   ##     ## ##
// ######    ##  ## ## ## ##     ##    ######### ##     ## ## ## ## ##     ##       ##       ##    ########  ######
// ##        ##  ##  #### ##     ##    ##     ## ######### ##  #### ##     ##       ##       ##    ##        ##
// ##        ##  ##   ### ##     ##    ##     ## ##     ## ##   ### ##     ##       ##       ##    ##        ##
// ##       #### ##    ## ########     ##     ## ##     ## ##    ## ########        ##       ##    ##        ########

// #####################################################################
// */

// // findSFsInList - doc line
// // Scans a list of hands, a slice of five card hands, and finds the SFs within
// func findSFsInList(clPtr *fiveCardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string, sSfFclPtr *[]fiveCardList) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// sSfFcl := *sSfFclPtr

// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)
// 	// fmt.Println("scPtr.max: ", scPtr.max)

// 	// fmt.Printf("findSFsInList\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
// 	// First check, the basics
// 	if cll < 5 || rcPtr.max > 4 || scPtr.max < 5 {
// 		err = errors.New("Failed prelim SF checks, exiting findSFsInList")
// 		return err
// 	} else {
// 		// Info.Printf("%s got past first screen \n\n", ThisFunc())
// 	}

// 	flushClPtr, topSuit, err01 := find5OrMoreOfSameSuitInfiveCardList(clPtr, cll, scPtr)

// 	if err01 != nil {
// 		fmt.Println(err01)
// 		err = errors.New("Trouble in find5OrMoreOfSameSuitInfiveCardList")
// 		return err
// 	}

// 	flushClOrderedPtr, err02 := orderCardsOfSameSuit2(flushClPtr, topSuit)
// 	flushClOrdered := *flushClOrderedPtr

// 	if err02 != nil {
// 		fmt.Println(err02)
// 		err = errors.New("Trouble in orderCardsOfSameSuit2")
// 		return err
// 	}

// 	rcFlPtr := countRanksInfiveCardList(flushClOrderedPtr)
// 	scFlPtr := countSuitsInfiveCardList(flushClOrderedPtr)

// 	err03 := checkForSt_5c(flushClOrderedPtr, getfiveCardListLength(flushClOrderedPtr), rcFlPtr, scFlPtr, handNamePtr)

// 	if err03 != nil {
// 		// fmt.Println("err03: ", err03)
// 		err = errors.New("Did not find SF")
// 	} else {
// 		// fmt.Println("Found SF XXXXXXXX")

// 		fmt.Println("flushClOrdered: ")

// 		for i, _ := range flushClOrdered {
// 			fmt.Printf("%v%v ", flushClOrdered[i].Rank, flushClOrdered[i].Suit)
// 		}

// 		fmt.Printf("\n")
// 		fmt.Println(flushClOrdered)
// 		fmt.Printf("\n\n")

// 		fmt.Println("5 card list: ")
// 		fmt.Println(flushClOrderedPtr)

// 		*sSfFclPtr = append(*sSfFclPtr, flushClOrdered)

// 		// fmt.Println("Appended another SF: ")
// 		// fmt.Println(*sSfFclPtr)

// 		*handNamePtr = string(flushClOrdered[0].Rank) + " high Straight Flush"
// 	}
// 	return err
// }

// // #####################################################################

// // find4xInList looks for quads - doc line
// func find4xInList(clPtr *fiveCardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string, sSfFclPtr *[]fiveCardList) (err error) {

// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)
// 	// fmt.Println("scPtr.max: ", scPtr.max)
// 	// fmt.Println("")

// 	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
// 	// First check, the basics
// 	if cll < 4 || rcPtr.max < 4 || rcPtr.uniqeRankCt > 4 {
// 		err = errors.New("Failed prelim 4x checks, exiting checkFor 4x")
// 		return err
// 	}

// L01:
// 	for _, rank := range manage_table.RankList {
// 		if rcPtr.rcm[rank] == 4 {
// 			rcPtr.top4x1 = rank
// 			break L01
// 		}
// 	}

// 	switch {
// 	case rcPtr.max == 4:
// 		*handNamePtr = "Four of a kind, " + string(rcPtr.top4x1) + "s"
// 		// fmt.Println("Found 4x ########")
// 	default:
// 		*handNamePtr = "did not find 4x"
// 		err = errors.New("err: Did not find 4x")
// 	}

// 	return err
// }

// /*

//  */

// // #####################################################################

// // checkForSF_5c - doc line
// // Looks for 5+ cards of same suit, then orders them, then looks for a straight
// func checkForSF_5c(clPtr *fiveCardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {

// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// os.Exit(0)

// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)
// 	// fmt.Println("scPtr.max: ", scPtr.max)

// 	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", *clPtr[0], *clPtr[1], *clPtr[2], *clPtr[3], *clPtr[4])
// 	// printFiveCardListAsString(clPtr)
// 	// First check, the basics
// 	if cll < 5 || rcPtr.max > 4 || scPtr.max < 5 {
// 		// fmt.Println("Failed prelim SF checks, exiting checkForSF_5c")
// 		err = errors.New("Failed prelim SF checks, exiting checkForSF_5c")

// 		return err
// 	}

// 	flushClPtr, topSuit, err01 := find5OrMoreOfSameSuitInfiveCardList(clPtr, cll, scPtr)

// 	if err01 != nil {
// 		// fmt.Println(err01)
// 		err = errors.New("Trouble in find5OrMoreOfSameSuitInCardList")
// 		return err
// 	}

// 	flushClOrderedPtr, err02 := orderCardsOfSameSuit2(flushClPtr, topSuit)
// 	flushClOrdered := *flushClOrderedPtr

// 	if err02 != nil {
// 		// fmt.Println(err02)
// 		err = errors.New("Trouble in orderCardsOfSameSuit2")
// 		return err
// 	}

// 	// fmt.Println("err 02 is nil")
// 	// os.Exit(0)

// 	rcFlPtr := countRanksInfiveCardList(flushClOrderedPtr)
// 	scFlPtr := countSuitsInfiveCardList(flushClOrderedPtr)

// 	err03 := checkForSt_5c(flushClOrderedPtr, getfiveCardListLength(flushClOrderedPtr), rcFlPtr, scFlPtr, handNamePtr)

// 	if err03 != nil {
// 		// fmt.Println("err03: ", err03)
// 		err = errors.New("Did not find SF")
// 	} else {
// 		// fmt.Println("Found SF XXXXXXXX")
// 		*handNamePtr = string(flushClOrdered[0].Rank) + " high Straight Flush"
// 	}

// 	return err
// }

// // #####################################################################

// // checkFo4x looks for quads - doc line
// func checkFor4x_5c(clPtr *fiveCardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)
// 	// fmt.Println("scPtr.max: ", scPtr.max)
// 	// fmt.Println("rcPtr.uniqRC: ", rcPtr.uniqeRankCt)
// 	// fmt.Println()

// 	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
// 	// First check, the basics
// 	if cll < 4 || rcPtr.max < 4 || rcPtr.uniqeRankCt > 4 {
// 		err = errors.New("Failed prelim 4x checks, exiting checkFor 4x")
// 		// fmt.Println("4x prelim failed")
// 		return err
// 	}

// L01:
// 	for _, rank := range manage_table.RankList {
// 		if rcPtr.rcm[rank] == 4 {
// 			rcPtr.top4x1 = rank
// 			break L01
// 		}
// 	}

// 	switch {
// 	case rcPtr.max == 4:
// 		*handNamePtr = "Four of a kind, " + string(rcPtr.top4x1) + "s"
// 		// fmt.Println("Found 4x ########")
// 	default:
// 		*handNamePtr = "did not find 4x"
// 		err = errors.New("err: Did not find 4x")
// 	}

// 	// fmt.Println("4x ret err: ", err)

// 	return err
// }

// // #####################################################################

// // checkFor3x_5c looks for 3 of a kind - doc line
// func checkFor3x_5c(clPtr *fiveCardList, cll int, rcPtr *rankCounter, handNamePtr *string) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)

// 	// fmt.Printf("inFindSF\n%v\n%v\n%v\n%v\n%v\n", lPtr[0], lPtr[1], lPtr[2], lPtr[3], lPtr[4])
// 	// First check, the basics
// 	if cll < 3 || rcPtr.max < 3 || rcPtr.uniqeRankCt > 5 {
// 		err = errors.New("Failed prelim 3x checks, exiting checkFor3x_5c")
// 		return err
// 	}

// L03:
// 	for _, rank := range manage_table.RankList {
// 		if rcPtr.rcm[rank] == 3 {
// 			rcPtr.top3x1 = rank
// 			break L03
// 		}
// 	}

// 	switch {
// 	case rcPtr.max == 3:
// 		*handNamePtr = "Three of a kind, " + string(rcPtr.top3x1) + "s"
// 		// fmt.Println("Found a ######################################## 3x")
// 	default:
// 		*handNamePtr = "did not find 3x"
// 		err = errors.New("err: Did not find 3x")
// 	}

// 	// fmt.Println("3x return err: ", err)
// 	return err
// }

// // #####################################################################

// // checkForFH_5c looks for a Full House - doc line
// func checkForFH_5c(fclPtr *fiveCardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {

// 	// We start the process with a couple of assumptions:
// 	// 1) This hand is definitely NOT a SF of 4x, since we checked for those already. This allow us to simplify the match conditions.

// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// printFiveCardListAsString(fclPtr, 0)

// 	// fmt.Println("### Looking for a FH #######")
// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)
// 	// fmt.Println("scPtr.max: ", scPtr.max)

// 	// Set default behavior
// 	*handNamePtr = "did not find a FH"
// 	err = errors.New("Did not find a FH")

// 	// First check, the basics
// 	if cll < 3 || rcPtr.max >= 6 {
// 		err = errors.New("Failed prelim FH checks, exiting checkForFH_5c")
// 		// fmt.Println("prelim check err: ", err)
// 		return err
// 	}

// 	// os.Exit(0)

// 	if rcPtr.max == 3 {
// 		// Opting to NOT use (>= 3), so that we'll catch issues with 4x detection, if there is a 4x, and we got here anyway.
// 		// The idea is that 4x checker should have found 4x already, and we should NOT be in this section of code, if
// 		// that's the case.
// 		//
// 		// Possible configs of the 7 cards at this point can be:
// 		// 3x, 3x, 1x     *
// 		// 3x, 2x, 2x     *
// 		// 3x, 2x, 1x, 1x *
// 		// 3x, 1x, 1x, 1x, 1x
// 		// Got an FH, if we have one of the *-ed situations

// 		// Work through the possible configurations
// 		ct3x := 0
// 		ct2x := 0

// 		for _, rank := range manage_table.RankList {

// 			switch {
// 			case rcPtr.rcm[rank] == 3: // found a 3x
// 				ct3x++

// 				switch {
// 				case ct3x == 1:
// 					rcPtr.top3x1 = rank
// 				case ct3x == 2: // 3x, 3x   // We houild never get here with a 5 card hand
// 					rcPtr.top3x2 = rank
// 					// found = true
// 				default:
// 				}

// 			case rcPtr.rcm[rank] == 2: // found a 2x
// 				ct2x++

// 				switch {
// 				case ct2x == 1:
// 					rcPtr.top2x1 = rank
// 					// found = true
// 				default:
// 				}

// 			default:
// 			} // End switch

// 		} // Since manage_table.RankList is arranged from highest to lowest rank, the top3x1 and top3x2, ..., are properly set.

// 		/*
// 			fmt.Println()
// 			fmt.Println("counter3x    : ", counter3x)
// 			fmt.Println("counter2x    : ", counter2x)
// 			fmt.Println("rcPtr.top3x1 : ", rcPtr.top3x1)
// 			fmt.Println("rcPtr.top3x2 : ", rcPtr.top3x2)
// 			fmt.Println("rcPtr.top2x1 : ", rcPtr.top2x1)
// 			fmt.Println()

// 		*/

// 		/*
// 			switch { // Considering 5 card hands only !!!
// 			// Only for 6+ cards:
// 			// case 3x, 3x, with 7 cards, there can't be a 2x here
// 			// case rcPtr.top3x2 != manage_table.RX:
// 			// 	*handNamePtr = "FH, " + string(rcPtr.top3x1) + "s full of " + string(rcPtr.top3x2) + "s."
// 			// 	// case 3x, 2x, 2x OR 3x, 2x, 1x, 1x
// 			// 	fmt.Println("Found FH ############## A")

// 			case rcPtr.top3x1 != manage_table.RX:
// 				*handNamePtr = "FH, " + string(rcPtr.top3x1) + "s full of " + string(rcPtr.top3x2) + "s."
// 				// case 3x, 2x
// 				fmt.Println("Found FH ############## A")

// 			case rcPtr.top2x1 != manage_table.RX:
// 				*handNamePtr = "FH, " + string(rcPtr.top3x1) + "s full of " + string(rcPtr.top2x1) + "s."
// 				fmt.Println("Found FH ############## B")

// 			default:
// 				*handNamePtr = "did not find a FH"
// 				err = errors.New("Did not find a FH")
// 			}

// 		*/

// 		if rcPtr.top2x1 != manage_table.RX {
// 			*handNamePtr = "FH, " + string(rcPtr.top3x1) + "s full of " + string(rcPtr.top2x1) + "s."
// 			// case 3x, 2x
// 			// fmt.Println("Found FH ############## A")

// 			// Make sure we return nil when FH is found
// 			err = nil
// 			// fmt.Println("FH count: ", counterFH)

// 		}

// 	}

// 	// fmt.Println("FH ret err: ", err)
// 	// os.Exit(0)
// 	return err

// }

// // #####################################################################

// // checkForFl_5c identifies the flush, and determines the the best 5 ranks that make it up. This allows us to compare flushes later.
// func checkForFl_5c(fclPtr *fiveCardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	var mySuit manage_table.CardSuit

// 	var cr1, cr2, cr3, cr4, cr5 manage_table.CardRank
// 	cr1 = manage_table.RX
// 	cr2 = manage_table.RX
// 	cr3 = manage_table.RX
// 	cr4 = manage_table.RX
// 	cr5 = manage_table.RX

// 	// fmt.Println("### Looking for a Flush #######")
// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)
// 	// fmt.Println("scPtr.max: ", scPtr.max)

// 	// First check, the basics
// 	if cll < 5 || scPtr.max < 5 {
// 		err = errors.New("Failed prelim Flush checks, exiting checkForFl_5c")
// 		return err
// 	}

// 	switch {
// 	case scPtr.scm[manage_table.S] >= 5:
// 		mySuit = manage_table.S
// 	case scPtr.scm[manage_table.C] >= 5:
// 		mySuit = manage_table.C
// 	case scPtr.scm[manage_table.H] >= 5:
// 		mySuit = manage_table.H
// 	case scPtr.scm[manage_table.D] >= 5:
// 		mySuit = manage_table.D
// 	default:
// 		*handNamePtr = "did not find a Flush"
// 		err = errors.New("Did not find a Fl")
// 		return err
// 	}

// 	cr1, cr2, cr3, cr4, cr5 = orderCardsOfSameSuit(fclPtr, mySuit)

// 	*handNamePtr = "Flush, " + string(cr1) + string(cr2) + string(cr3) + string(cr4) + string(cr5)

// 	// fmt.Println("FL ret err: ", err)
// 	// fmt.Println("Found a ####################### FLUSH #######")

// 	// printFiveCardListAsString(fclPtr, 0)

// 	/*
// 			if counterFl
// 		 == 1000 {
// 				os.Exit(0)
// 			}
// 	*/
// 	return err
// }

// // #####################################################################

// // checkForSt_5c looks for straights - doc line
// func checkForSt_5c(clPtr *fiveCardList, cll int, rcPtr *rankCounter, scPtr *suitCounter, handNamePtr *string) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	cl := *clPtr
// 	/*
// 		fmt.Println("card list len, cll: ", cll)
// 		fmt.Println("rcPtr.uniqeRankCt: ", rcPtr.uniqeRankCt)
// 		fmt.Println("rcPtr.max: ", rcPtr.max)
// 		fmt.Println("scPtr.max: ", scPtr.max)
// 	*/

// 	// default state
// 	*handNamePtr = "did not find a St"
// 	err = errors.New("Did not find a St")
// 	// fmt.Println("did not find a St, err: ", err)

// 	// First check, the basics
// 	// Since St is getting to ge rather far in the hand classification, some of the prelim checks are becoming "absolete", ie. certain configurations
// 	// should never bee seen this far dowin in the hand rank check list.
// 	// if cll < 5 || rcPtr.uniqeRankCt < 5 || scPtr.max > 4 {
// 	if cll < 5 || rcPtr.uniqeRankCt < 5 {
// 		err = errors.New("Failed prelim St checks, exiting checkForSt_5c")
// 		// Info.Printf("%s; FAILED prelim check.\n\n", ThisFunc())
// 		return err
// 	} else {
// 		// Info.Printf("%s; Got past prelim check.\n\n", ThisFunc())
// 		// os.Exit(0)
// 	}

// L02:
// 	for i := 0; i <= 9; i++ {
// 		// fmt.Println("i: ", i)
// 		// fmt.Printf("in checkForSt_5c manage_table.RankListFull of i is: %v\n\n", manage_table.RankListFull[i])
// 		// fmt.Printf("clPtr 0 : %v; 1: %v; 2: %v; 3: %v; 4: %v; 5: %v; 6: %v\n", clPtr[0].Rank, clPtr[1], clPtr[2], clPtr[3], clPtr[4], clPtr[5], clPtr[6])
// 		// fmt.Printf("rankLFi : %v;+1: %v;+2: %v;+3: %v;+4: %v\n\n", manage_table.RankListFull[i], manage_table.RankListFull[i+1], manage_table.RankListFull[i+2], manage_table.RankListFull[i+3], manage_table.RankListFull[i+4])

// 		// if (clPtr[0].Rank == manage_table.RankListFull[i] || clPtr[1].Rank == manage_table.RankListFull[i] || clPtr[2].Rank == manage_table.RankListFull[i] || clPtr[3].Rank == manage_table.RankListFull[i] || clPtr[4].Rank == manage_table.RankListFull[i] || clPtr[5].Rank == manage_table.RankListFull[i] || clPtr[6].Rank == manage_table.RankListFull[i]) &&
// 		// 	(clPtr[0].Rank == manage_table.RankListFull[i+1] || clPtr[1].Rank == manage_table.RankListFull[i+1] || clPtr[2].Rank == manage_table.RankListFull[i+1] || clPtr[3].Rank == manage_table.RankListFull[i+1] || clPtr[4].Rank == manage_table.RankListFull[i+1] || clPtr[5].Rank == manage_table.RankListFull[i+1] || clPtr[6].Rank == manage_table.RankListFull[i+1]) &&
// 		// 	(clPtr[0].Rank == manage_table.RankListFull[i+2] || clPtr[1].Rank == manage_table.RankListFull[i+2] || clPtr[2].Rank == manage_table.RankListFull[i+2] || clPtr[3].Rank == manage_table.RankListFull[i+2] || clPtr[4].Rank == manage_table.RankListFull[i+2] || clPtr[5].Rank == manage_table.RankListFull[i+2] || clPtr[6].Rank == manage_table.RankListFull[i+2]) &&
// 		// 	(clPtr[0].Rank == manage_table.RankListFull[i+3] || clPtr[1].Rank == manage_table.RankListFull[i+3] || clPtr[2].Rank == manage_table.RankListFull[i+3] || clPtr[3].Rank == manage_table.RankListFull[i+3] || clPtr[4].Rank == manage_table.RankListFull[i+3] || clPtr[5].Rank == manage_table.RankListFull[i+3] || clPtr[6].Rank == manage_table.RankListFull[i+3]) &&
// 		// 	(clPtr[0].Rank == manage_table.RankListFull[i+4] || clPtr[1].Rank == manage_table.RankListFull[i+4] || clPtr[2].Rank == manage_table.RankListFull[i+4] || clPtr[3].Rank == manage_table.RankListFull[i+4] || clPtr[4].Rank == manage_table.RankListFull[i+4] || clPtr[5].Rank == manage_table.RankListFull[i+4] || clPtr[6].Rank == manage_table.RankListFull[i+4]) {

// 		if (cl[0].Rank == manage_table.RankListFull[i] || cl[1].Rank == manage_table.RankListFull[i] || cl[2].Rank == manage_table.RankListFull[i] || cl[3].Rank == manage_table.RankListFull[i] || cl[4].Rank == manage_table.RankListFull[i]) &&
// 			(cl[0].Rank == manage_table.RankListFull[i+1] || cl[1].Rank == manage_table.RankListFull[i+1] || cl[2].Rank == manage_table.RankListFull[i+1] || cl[3].Rank == manage_table.RankListFull[i+1] || cl[4].Rank == manage_table.RankListFull[i+1]) &&
// 			(cl[0].Rank == manage_table.RankListFull[i+2] || cl[1].Rank == manage_table.RankListFull[i+2] || cl[2].Rank == manage_table.RankListFull[i+2] || cl[3].Rank == manage_table.RankListFull[i+2] || cl[4].Rank == manage_table.RankListFull[i+2]) &&
// 			(cl[0].Rank == manage_table.RankListFull[i+3] || cl[1].Rank == manage_table.RankListFull[i+3] || cl[2].Rank == manage_table.RankListFull[i+3] || cl[3].Rank == manage_table.RankListFull[i+3] || cl[4].Rank == manage_table.RankListFull[i+3]) &&
// 			(cl[0].Rank == manage_table.RankListFull[i+4] || cl[1].Rank == manage_table.RankListFull[i+4] || cl[2].Rank == manage_table.RankListFull[i+4] || cl[3].Rank == manage_table.RankListFull[i+4] || cl[4].Rank == manage_table.RankListFull[i+4]) {

// 			// fmt.Printf("in checkForSt_5c inside IF\n%v %v %v %v %v\n", manage_table.RankListFull[i], manage_table.RankListFull[i+1], manage_table.RankListFull[i+2], manage_table.RankListFull[i+3], manage_table.RankListFull[i+4])

// 			*handNamePtr = string(manage_table.RankListFull[i]) + " high Straight"
// 			err = nil // since it could have been set to "not nil"
// 			break L02
// 		}
// 	}
// 	if err == nil {
// 		// fmt.Println("Foudn a ################################ STRAIGHT")

// 	}
// 	// os.Exit(0)
// 	return err
// }

// // #####################################################################

// // checkFor2x2_5c looks for two pair hands - doc line
// func checkFor2x2_5c(clPtr *fiveCardList, cll int, rcPtr *rankCounter, handNamePtr *string) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// We need a local copy of rc, just so we can do a range operation.
// 	// All assignements will be done with rcPtr that was passed to us.
// 	rc := *rcPtr

// 	pair_count := 0

// 	// Look for 2 ranks with count of 2, in the rank counter map, and update the top2x1 and top2x2 vars.
// 	for cr, count := range rc.rcm {
// 		// fmt.Printf("rank:[%s] count[%d]\n", cr, count)

// 		switch count {
// 		case 2:

// 			pair_count++

// 			if rcPtr.top2x1 == manage_table.RX {
// 				rcPtr.top2x1 = cr
// 			} else {
// 				if crm[cr] > crm[rcPtr.top2x1] {
// 					rcPtr.top2x2 = rcPtr.top2x1
// 					rcPtr.top2x1 = cr
// 				}
// 			}
// 		case 1:
// 			rcPtr.top1x1 = cr
// 		default:

// 		}
// 	}

// 	// orderedClPtr, err := orderCardsOfMixedSuit(clPtr)
// 	// err04 := findPairs(orderedClPtr, cll, rcPtr) // Assigns the 2x var values in the rcPtr

// 	// if err04 == nil && rcPtr.top2x2 != manage_table.RX {
// 	if pair_count == 2 {
// 		*handNamePtr = "Two pair, " + string(rcPtr.top2x1) + "s and " + string(rcPtr.top2x2) + "s, " + string(rcPtr.top1x1) + " kicker"
// 	} else {
// 		err = errors.New("did not find a 2 pair hand")
// 		// fmt.Println("did not find 2x2; err: ", err)
// 	}

// 	if err == nil {
// 		// fmt.Println("Foudn a ################################ 2x2")
// 	}
// 	// os.Exit(0)

// 	return err
// }

// // #####################################################################

// // findPairs examines the card list and if it finds pairs, assigns appropriate ranks to rcCouter.topXXXX vars.
// func findPairs(clPtr *fiveCardList, cll int, rcPtr *rankCounter) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	cl := *clPtr

// 	// fmt.Println("cll-2: ", cll-2)
// 	// fmt.Println("dump rcPtr: ", rcPtr)

// 	for i := 0; i < cll-2; i++ {
// 		// fmt.Println("i: ", i)
// 		// fmt.Println("clPtr[i]: ", clPtr[i], "clPtr[i+1]: ", clPtr[i+1], "rcPtr.top2x1: ", rcPtr.top2x1)

// 		if cl[i].Rank == cl[i+1].Rank {
// 			if rcPtr.top2x1 == manage_table.RX {
// 				// fmt.Println("in findPairs, inside IF, found a pair")
// 				rcPtr.top2x1 = cl[i].Rank
// 			} else if rcPtr.top2x2 == manage_table.RX {
// 				// fmt.Println("in findPairs, inside IF, found 2nd pair")
// 				rcPtr.top2x2 = cl[i].Rank
// 			}
// 			i = i + 1 // if we encountered the first card in a pair, then skip the second
// 		} else {
// 			switch {
// 			case rcPtr.top1x1 == manage_table.RX:
// 				rcPtr.top1x1 = cl[i].Rank
// 			case rcPtr.top1x2 == manage_table.RX:
// 				rcPtr.top1x2 = cl[i].Rank
// 			case rcPtr.top1x3 == manage_table.RX:
// 				rcPtr.top1x3 = cl[i].Rank
// 			case rcPtr.top1x4 == manage_table.RX:
// 				rcPtr.top1x4 = cl[i].Rank
// 			case rcPtr.top1x5 == manage_table.RX:
// 				rcPtr.top1x5 = cl[i].Rank
// 			default:
// 				fmt.Println("reached default case in findPairs; i: ", i)
// 			}
// 		}
// 	}

// 	if rcPtr.top2x1 == manage_table.RX {
// 		err = errors.New("did not find any pairs")
// 		// fmt.Println("did not find any pairs; err: ", err)
// 	}

// 	// fmt.Println("rcPtr.top2x1: ", rcPtr.top2x1, "rcPtr.top2x2: ", rcPtr.top2x2)

// 	return err
// }

// // #####################################################################

// // checkFor2x1_5c looks for a single pair
// // So, checkFor2x2_5c ran the findPairs, which set up the rcPtr to the end. We can just use that info here
// // and do not need to figure anything out.
// func checkFor2x1_5c(clPtr *fiveCardList, cll int, rcPtr *rankCounter, handNamePtr *string) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)

// 	// orderedClPtr, err := orderCardsOfMixedSuit(clPtr)
// 	// err05 := findPairs(orderedClPtr, cll, rcPtr) // Assigns the 2x var values in the rcPtr

// 	if rcPtr.top2x1 != manage_table.RX {
// 		*handNamePtr = "Pair, " + string(rcPtr.top2x1) + "s, " + string(rcPtr.top1x1) + string(rcPtr.top1x2) + string(rcPtr.top1x3) + " kicker."
// 	} else {
// 		err = errors.New("did not find any pairs")
// 		// fmt.Println("did not find 2x1; err: ", err)
// 	}

// 	if err == nil {
// 		// fmt.Println("Foudn a ################################ 2x1")
// 	}
// 	// os.Exit(0)

// 	return err
// }

// // #####################################################################

// // checkForHc_5c looks for a high card hand.
// // Well, at this stage any other hand should have been caught by the pervious findXXX functions.
// // So, checkFor2x2_5c ran the findPairs, which set up the rcPtr to the end. We can just use that info here
// // and do not need to figure anything out.
// func checkForHc_5c(clPtr *fiveCardList, cll int, rcPtr *rankCounter, handNamePtr *string) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// fmt.Println("card list len, cll: ", cll)
// 	// fmt.Println("rcPtrmax: ", rcPtr.max)

// 	if rcPtr.top2x1 == manage_table.RX {
// 		*handNamePtr = "High card: " + string(rcPtr.top1x1) + ", " + string(rcPtr.top1x2) + string(rcPtr.top1x3) + string(rcPtr.top1x4) + string(rcPtr.top1x5) + " kicker."
// 	} else {
// 		err = errors.New("did not find any pairs")
// 		// fmt.Println("did not find 2x1; err: ", err)
// 	}

// 	return err
// }

// /*

//  ######  ########  ########    ###    ######## ########    ##     ##    ###    ##    ## ########     ######## ##    ## ########  ########
// ##    ## ##     ## ##         ## ##      ##    ##          ##     ##   ## ##   ###   ## ##     ##       ##     ##  ##  ##     ## ##
// ##       ##     ## ##        ##   ##     ##    ##          ##     ##  ##   ##  ####  ## ##     ##       ##      ####   ##     ## ##
// ##       ########  ######   ##     ##    ##    ######      ######### ##     ## ## ## ## ##     ##       ##       ##    ########  ######
// ##       ##   ##   ##       #########    ##    ##          ##     ## ######### ##  #### ##     ##       ##       ##    ##        ##
// ##    ## ##    ##  ##       ##     ##    ##    ##          ##     ## ##     ## ##   ### ##     ##       ##       ##    ##        ##
//  ######  ##     ## ######## ##     ##    ##    ########    ##     ## ##     ## ##    ## ########        ##       ##    ##        ########

// */

// // #####################################################################

// // !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// // !!! As a convention, any fiveCardList MUST be ordered by rank, A to 2.
// // Functions that process RELY on the order being correct
// // This will simpify some code down the line, as certain assumptions will
// // be valid, and certain manipulations no longer necessary.
// // !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// // #####################################################################

// // createSFCL - doc line
// func createSFCL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr

// 	j := 39 // Legal index range is 0-39; 39 + 4*4(16) = 55(56-1), 5h sf
// 	j = 0   // 0-3: A high sf, 4-7: K high SF, ... 36-39: 5h SF
// 	cl[0] = orderedListFull[j]
// 	cl[1] = orderedListFull[j+4]
// 	cl[2] = orderedListFull[j+8]
// 	cl[3] = orderedListFull[j+12]
// 	cl[4] = orderedListFull[j+16]
// 	// need 2 more for 7 cards total
// 	// clPtr[5] = orderedListFull[j+30]
// 	// clPtr[6] = orderedListFull[j+40]

// 	return clPtr
// }

// // #####################################################################

// // create4xCL - doc line
// func create4xCL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr

// 	cl[0] = orderedList[0]
// 	cl[1] = orderedList[1]
// 	cl[2] = orderedList[2]
// 	cl[3] = orderedList[3]
// 	cl[4] = orderedList[4]
// 	// need 2 more for 7 cards total
// 	// clPtr[5] = orderedList[18]
// 	// clPtr[6] = orderedList[45]

// 	return clPtr
// }

// // #####################################################################

// // createFHCL - doc line
// func createFHCL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr

// 	// FH 3x, 3x, 1x
// 	cl[0] = orderedList[0]
// 	cl[1] = orderedList[1]
// 	cl[2] = orderedList[2]
// 	cl[3] = orderedList[4]

// 	// clPtr[4] = orderedList[5]
// 	// clPtr[5] = orderedList[6]

// 	// clPtr[6] = orderedList[45]

// 	return clPtr
// }

// // #####################################################################

// // createFlCL - doc line
// func createFlCL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr

// 	j := 0 //

// 	cl[0] = orderedList[j]
// 	cl[1] = orderedList[j+4]
// 	cl[2] = orderedList[j+8]
// 	cl[3] = orderedList[j+12]
// 	cl[4] = orderedList[j+24]

// 	// clPtr[5] = orderedList[j+30]
// 	// clPtr[6] = orderedList[j+45]

// 	return clPtr
// }

// // #####################################################################

// // createStCL - doc line
// func createStCL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr

// 	cl[0] = orderedListFull[0]
// 	cl[1] = orderedListFull[42]
// 	cl[2] = orderedListFull[47]
// 	cl[3] = orderedListFull[50]
// 	cl[4] = orderedListFull[39]

// 	// clPtr[5] = orderedListFull[5]
// 	// clPtr[6] = orderedListFull[9]

// 	return clPtr
// }

// // #####################################################################

// // create3xCL - doc line
// func create3xCL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr

// 	// 3x
// 	// j := 39 // Legal index range is 0-39; 39 + 4*4(16) = 55(56-1), 5h sf
// 	j := 0 // 0-3: A high sf, 4-7: K high SF, ... 36-39: 5h SF

// 	cl[0] = orderedList[j]
// 	cl[1] = orderedList[j+1]
// 	cl[2] = orderedList[j+3]

// 	cl[3] = orderedList[20]
// 	cl[4] = orderedList[25]
// 	// clPtr[5] = orderedList[30]
// 	// clPtr[6] = orderedList[35]

// 	return clPtr
// }

// // #####################################################################

// // create2x2CL - doc line
// func create2x2CL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr

// 	// Two pair
// 	cl[0] = orderedList[0]
// 	cl[1] = orderedList[1]
// 	cl[2] = orderedList[4]
// 	cl[3] = orderedList[5]
// 	cl[4] = orderedList[39]

// 	// clPtr[5] = orderedList[30]
// 	// clPtr[6] = orderedList[40]

// 	return clPtr
// }

// // #####################################################################

// // create2x1CL - doc line
// func create2x1CL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr

// 	// Two pair
// 	cl[0] = orderedList[0]
// 	cl[1] = orderedList[1]
// 	cl[2] = orderedList[10]
// 	cl[3] = orderedList[15]
// 	cl[4] = orderedList[39]

// 	// clPtr[5] = orderedList[30]
// 	// clPtr[6] = orderedList[40]

// 	return clPtr
// }

// // #####################################################################

// // createHcCL - doc line
// func createHcCL() (clPtr *fiveCardList) {
// 	clPtr = createClPtr()
// 	cl := *clPtr
// 	// Two pair
// 	cl[0] = orderedList[0]
// 	cl[1] = orderedList[5]
// 	cl[2] = orderedList[10]
// 	cl[3] = orderedList[15]
// 	cl[4] = orderedList[39]
// 	// clPtr[5] = orderedList[30]
// 	// clPtr[6] = orderedList[40]

// 	return clPtr
// }

// // #####################################################################
// func getRankOfNthCardForGivenIndexOfscl(i int, n int) manage_table.CardRank {
// 	// func getRankOfNthCardForGivenIndexOfscl(i int, n int) int {

// 	return sOfAllFCLs[i][n].Rank
// 	// return crm[sOfAllFCLs[i][n].Rank]
// }

// // #####################################################################
// // sortCardsInEach5CList sorts cards by rank within each 5 card hand provided in a hand list.
// func sortCardsInEach5CList() (err error) {

// 	for _, fcl := range sOfAllFCLs {

// 		sort.SliceStable(fcl, func(i, j int) bool {
// 			return crm[fcl[i].Rank] > crm[fcl[j].Rank]
// 		})
// 	}
// 	return nil
// }

// // #####################################################################
// // orderSFIndexesAsc orders the indexes according to hand strength, low to high
// func orderSFIndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false
// 		// If our fcl starts with A5, then we have a wheel straight, the lowest straight, so
// 		// this index must end up at the opposite end from the AK hands.
// 		if sOfAllFCLs[si[i]][0].Rank == manage_table.RA && sOfAllFCLs[si[i]][1].Rank == manage_table.R5 {
// 			iltj = true
// 		} else if sOfAllFCLs[si[j]][0].Rank == manage_table.RA && sOfAllFCLs[si[j]][1].Rank == manage_table.R5 {
// 			iltj = false
// 		} else if crm[sOfAllFCLs[si[i]][0].Rank] < crm[sOfAllFCLs[si[j]][0].Rank] {
// 			iltj = true
// 		}
// 		return iltj
// 	})
// 	return nil
// }

// // #####################################################################
// // orderSFIndexesDes orders the indexes according to hand strength, high to low.
// func orderSFIndexesDes(siPtr *[]int) (err error) {
// 	Info.Printf("%s\n\n", debugging.ThisFunc())

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false
// 		// If our fcl starts with A5, then we have a wheel straight, the lowest straight, so
// 		// this index must end up at the opposite end from the AK hands.
// 		if sOfAllFCLs[si[i]][0].Rank == manage_table.RA && sOfAllFCLs[si[i]][1].Rank == manage_table.R5 {
// 			iltj = false
// 		} else if sOfAllFCLs[si[j]][0].Rank == manage_table.RA && sOfAllFCLs[si[j]][1].Rank == manage_table.R5 {
// 			iltj = true
// 		} else if crm[sOfAllFCLs[si[i]][0].Rank] > crm[sOfAllFCLs[si[j]][0].Rank] {
// 			iltj = true
// 		}
// 		return iltj
// 	})
// 	return nil
// }

// // #####################################################################
// // orderSorder4xIndexesAsc orders the indexes according to hand strength, low to high.
// func order4xIndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	// sort.SliceStable(si, func(i, j int) bool {
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false

// 		// First, let's figure out which index, 0 or 4, our kicker sits at.
// 		ikickerIndexIsAt0 := false // If it's not 0, it must be 4
// 		if crm[sOfAllFCLs[si[i]][0].Rank] != crm[sOfAllFCLs[si[i]][1].Rank] {
// 			ikickerIndexIsAt0 = true
// 		}
// 		jkickerIndexIsAt0 := false // If it's not 0, it must be 4
// 		if crm[sOfAllFCLs[si[j]][0].Rank] != crm[sOfAllFCLs[si[j]][1].Rank] {
// 			jkickerIndexIsAt0 = true
// 		}

// 		switch {

// 		// Consider i and j hands with kiker at index 0
// 		case ikickerIndexIsAt0 && jkickerIndexIsAt0:
// 			// If same rank of 4x, compare kickers
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 				if crm[sOfAllFCLs[si[i]][0].Rank] < crm[sOfAllFCLs[si[j]][0].Rank] {
// 					iltj = true
// 				}
// 				// For different rank of 4x, kicker does not matter
// 			} else {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			}

// 		// Consider i and j hands with kiker at index 4
// 		// 4x rank at index 0123
// 		case !ikickerIndexIsAt0 && !jkickerIndexIsAt0:
// 			// If same rank of 4x, compare kickers
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 				// if sOfAllFCLs[si[i]][4].Rank < sOfAllFCLs[si[j]][4].Rank {
// 				if crm[sOfAllFCLs[si[i]][4].Rank] < crm[sOfAllFCLs[si[j]][4].Rank] {
// 					iltj = true
// 				}
// 				// For different rank of 4x, kicker does not matter
// 			} else {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			}

// 		// Consider i kicker at 0 and j kicker at 4
// 		// => 4x i at 1234 and 4x j at 0123
// 		case ikickerIndexIsAt0 && !jkickerIndexIsAt0:
// 			// If same rank of 4x, compare kickers
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 				// if sOfAllFCLs[si[i]][0].Rank < sOfAllFCLs[si[j]][4].Rank {
// 				if crm[sOfAllFCLs[si[i]][0].Rank] < crm[sOfAllFCLs[si[j]][4].Rank] {
// 					iltj = true
// 				}
// 				// For different rank of 4x, kicker does not matter
// 			} else {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			}

// 		// Consider i kicker at 4 and j kicker at 0
// 		// => 4x i at 0123 and 4x j at 1234
// 		case !ikickerIndexIsAt0 && jkickerIndexIsAt0:
// 			// If same rank of 4x, compare kickers
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 				// if sOfAllFCLs[si[i]][4].Rank < sOfAllFCLs[si[j]][0].Rank {
// 				if crm[sOfAllFCLs[si[i]][4].Rank] < crm[sOfAllFCLs[si[j]][0].Rank] {
// 					iltj = true
// 				}
// 				// For different rank of 4x, kicker does not matter
// 			} else {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			}

// 		default:
// 			err = errors.New("default case reached: ")
// 			Info.Printf("%s\n\n", debugging.ThisFunc())
// 		}

// 		return iltj
// 	})

// 	return err
// }

// // #####################################################################
// func orderFHIndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	// sort.SliceStable(si, func(i, j int) bool {
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false

// 		// First, let's figure out which side the 2x cards are on.
// 		// In either case, card index 2 will always belong to the 3x.
// 		i2xAt0and1 := false
// 		if crm[sOfAllFCLs[si[i]][0].Rank] != crm[sOfAllFCLs[si[i]][2].Rank] {
// 			i2xAt0and1 = true
// 		}
// 		j2xAt0and1 := false
// 		if crm[sOfAllFCLs[si[j]][0].Rank] != crm[sOfAllFCLs[si[j]][2].Rank] {
// 			j2xAt0and1 = true
// 		}

// 		switch {

// 		// Consider i and j hands with 2x cards at index 0-1
// 		case i2xAt0and1 && j2xAt0and1:
// 			// If same rank of 3x, compare kickers
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 				if crm[sOfAllFCLs[si[i]][0].Rank] < crm[sOfAllFCLs[si[j]][0].Rank] {
// 					iltj = true
// 				}
// 				// For different rank of 3x, 2x cards ds not matter.
// 			} else {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			}

// 		// Consider i and j hands with 2x cards at index 3-4
// 		// 3x rank at index 0,1,2
// 		case !i2xAt0and1 && !j2xAt0and1:
// 			// If same rank of 3x, compare 2x cards
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 				if crm[sOfAllFCLs[si[i]][4].Rank] < crm[sOfAllFCLs[si[j]][4].Rank] {
// 					iltj = true
// 				}
// 				// For different rank of 3x, kicker does not matter
// 			} else {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			}

// 		// Consider i 2x at 0,1 and j 2x at 3,4
// 		// => i3x at 234 and -3x j at 012
// 		case i2xAt0and1 && !j2xAt0and1:
// 			// If same rank of 4x, compare kickers
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 				if crm[sOfAllFCLs[si[i]][0].Rank] < crm[sOfAllFCLs[si[j]][4].Rank] {
// 					iltj = true
// 				}
// 				// For different rank of 3x, 2x cards do not matter
// 			} else {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			}

// 		// Consider i 2x at 3,4 and j 2x at 0,1
// 		// => 3x i at 012 and 3x j at 234
// 		case !i2xAt0and1 && j2xAt0and1:
// 			// If same rank of 3x, compare 2x cards
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 				if crm[sOfAllFCLs[si[i]][4].Rank] < crm[sOfAllFCLs[si[j]][0].Rank] {
// 					iltj = true
// 				}
// 				// For different rank of 3x, 2x cards do not matter
// 			} else {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			}

// 		default:
// 			err = errors.New("default case reached: ")
// 			Info.Printf("%s\n\n", debugging.ThisFunc())
// 		}

// 		return iltj
// 	})

// 	return err
// }

// // #####################################################################
// func orderFlIndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false
// 		//
// 		if crm[sOfAllFCLs[si[i]][0].Rank] == crm[sOfAllFCLs[si[j]][0].Rank] {
// 			if crm[sOfAllFCLs[si[i]][1].Rank] == crm[sOfAllFCLs[si[j]][1].Rank] {
// 				if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[j]][2].Rank] {
// 					if crm[sOfAllFCLs[si[i]][3].Rank] == crm[sOfAllFCLs[si[j]][3].Rank] {
// 						if crm[sOfAllFCLs[si[i]][4].Rank] < crm[sOfAllFCLs[si[j]][4].Rank] {
// 							iltj = true
// 						}
// 					} else if crm[sOfAllFCLs[si[i]][3].Rank] < crm[sOfAllFCLs[si[j]][3].Rank] {
// 						iltj = true
// 					}
// 				} else if crm[sOfAllFCLs[si[i]][2].Rank] < crm[sOfAllFCLs[si[j]][2].Rank] {
// 					iltj = true
// 				}
// 			} else if crm[sOfAllFCLs[si[i]][1].Rank] < crm[sOfAllFCLs[si[j]][1].Rank] {
// 				iltj = true
// 			}
// 		} else if crm[sOfAllFCLs[si[i]][0].Rank] < crm[sOfAllFCLs[si[j]][0].Rank] {
// 			iltj = true
// 		}

// 		return iltj
// 	})
// 	return nil
// }

// // #####################################################################
// // orderStIndexesAsc is the same as order SF
// func orderStIndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	_ = orderSFIndexesAsc(siPtr)
// 	return nil
// }

// // #####################################################################
// // See if hand at index i is LESS than hand at index j
// func cmpare3xHandsWithKickers(siPtr *[]int, iIdx, jIdx, ikHi, ikLi, jkHi, jkLi int) bool {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	iltj := false

// 	// If same rank of 3x, compare kickers
// 	if crm[sOfAllFCLs[si[iIdx]][2].Rank] == crm[sOfAllFCLs[si[jIdx]][2].Rank] {

// 		// If
// 		if crm[sOfAllFCLs[si[iIdx]][ikHi].Rank] < crm[sOfAllFCLs[si[jIdx]][jkHi].Rank] {
// 			iltj = true
// 			return iltj
// 		}

// 		// If H kickers are the same, compare the L kickers.
// 		if crm[sOfAllFCLs[si[iIdx]][ikHi].Rank] == crm[sOfAllFCLs[si[jIdx]][jkHi].Rank] {

// 			if crm[sOfAllFCLs[si[iIdx]][ikLi].Rank] < crm[sOfAllFCLs[si[jIdx]][jkLi].Rank] {
// 				iltj = true
// 			}
// 		}
// 		// For different rank of 3x, kicker cards do not matter.
// 	} else {
// 		if crm[sOfAllFCLs[si[iIdx]][2].Rank] < crm[sOfAllFCLs[si[jIdx]][2].Rank] {
// 			iltj = true
// 		}
// 	}
// 	return iltj
// }

// // #####################################################################
// func order3xIndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	var iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex int

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	// sort.SliceStable(si, func(i, j int) bool {
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false

// 		// First, let's figure out where the kicker cards are. We can have:
// 		// i	 j
// 		// xxxab xxxab
// 		// axxxb axxxb
// 		// abxxx abxxx
// 		// 9 possible scenarios

// 		i2KickHighAt0 := false // high kicer at 3
// 		if crm[sOfAllFCLs[si[i]][0].Rank] != crm[sOfAllFCLs[si[i]][2].Rank] {
// 			i2KickHighAt0 = true
// 		}
// 		i2KickLowhAt1 := false // low kicker at 4
// 		if crm[sOfAllFCLs[si[i]][1].Rank] != crm[sOfAllFCLs[si[i]][2].Rank] {
// 			i2KickLowhAt1 = true
// 		}
// 		j2KickHighAt0 := false // high kicer at 3
// 		if crm[sOfAllFCLs[si[j]][0].Rank] != crm[sOfAllFCLs[si[j]][2].Rank] {
// 			j2KickHighAt0 = true
// 		}
// 		j2KickLowhAt1 := false // low kicker at 4
// 		if crm[sOfAllFCLs[si[j]][1].Rank] != crm[sOfAllFCLs[si[j]][2].Rank] {
// 			j2KickLowhAt1 = true
// 		}

// 		switch {
// 		// #################
// 		// abxxx abxxx
// 		// Consider i and j hands with kicker kicker high at 0 and kicker low at 1
// 		case i2KickHighAt0 && i2KickLowhAt1 && j2KickHighAt0 && j2KickLowhAt1:
// 			iKickerHighIndex = 0
// 			iKickerLowIndex = 1
// 			jKickerHighIndex = 0
// 			jKickerLowIndex = 1
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)
// 		// abxxx axxxb
// 		// i kicker h at 0, i kicker low at 1, j kicker high at 0, j kicker low at 4
// 		case i2KickHighAt0 && i2KickLowhAt1 && j2KickHighAt0 && !j2KickLowhAt1:
// 			iKickerHighIndex = 0
// 			iKickerLowIndex = 1
// 			jKickerHighIndex = 0
// 			jKickerLowIndex = 4
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)
// 		/// abxxx xxxab
// 		case i2KickHighAt0 && i2KickLowhAt1 && !j2KickHighAt0 && !j2KickLowhAt1:
// 			iKickerHighIndex = 0
// 			iKickerLowIndex = 1
// 			jKickerHighIndex = 3
// 			jKickerLowIndex = 4
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)

// 		// #################
// 		// axxxb abxxx
// 		case i2KickHighAt0 && !i2KickLowhAt1 && j2KickHighAt0 && j2KickLowhAt1:
// 			iKickerHighIndex = 0
// 			iKickerLowIndex = 4
// 			jKickerHighIndex = 0
// 			jKickerLowIndex = 1
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)
// 		// axxxb axxxb
// 		case i2KickHighAt0 && !i2KickLowhAt1 && j2KickHighAt0 && !j2KickLowhAt1:
// 			iKickerHighIndex = 0
// 			iKickerLowIndex = 4
// 			jKickerHighIndex = 0
// 			jKickerLowIndex = 4
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)
// 		/// axxxb xxxab
// 		case i2KickHighAt0 && !i2KickLowhAt1 && !j2KickHighAt0 && !j2KickLowhAt1:
// 			iKickerHighIndex = 0
// 			iKickerLowIndex = 4
// 			jKickerHighIndex = 3
// 			jKickerLowIndex = 4
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)

// 		// #################
// 		// xxxab abxxx
// 		case !i2KickHighAt0 && !i2KickLowhAt1 && j2KickHighAt0 && j2KickLowhAt1:
// 			iKickerHighIndex = 3
// 			iKickerLowIndex = 4
// 			jKickerHighIndex = 0
// 			jKickerLowIndex = 1
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)
// 		// xxxab axxxb
// 		case !i2KickHighAt0 && !i2KickLowhAt1 && j2KickHighAt0 && !j2KickLowhAt1:
// 			iKickerHighIndex = 3
// 			iKickerLowIndex = 4
// 			jKickerHighIndex = 0
// 			jKickerLowIndex = 4
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)
// 		/// xxxab xxxab
// 		case !i2KickHighAt0 && !i2KickLowhAt1 && !j2KickHighAt0 && !j2KickLowhAt1:
// 			iKickerHighIndex = 3
// 			iKickerLowIndex = 4
// 			jKickerHighIndex = 3
// 			jKickerLowIndex = 4
// 			iltj = cmpare3xHandsWithKickers(siPtr, i, j, iKickerHighIndex, iKickerLowIndex, jKickerHighIndex, jKickerLowIndex)

// 		// #################
// 		default:
// 			err = errors.New("default case reached: ")
// 			Info.Printf("%s\n\n", debugging.ThisFunc())
// 		}

// 		return iltj
// 	})

// 	return err
// }

// // #####################################################################
// func order2x2IndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false

// 		// Determine which spot the kicker is in
// 		// 1) K aa bb
// 		// 2) aa K bb
// 		// 3) aa bb K

// 		// Give kicker position intial values, start at option 1
// 		iKIndex := 0 // 0-5
// 		jKIndex := 0 // 0-5

// 		if crm[sOfAllFCLs[si[i]][0].Rank] == crm[sOfAllFCLs[si[i]][1].Rank] {
// 			// mst be in scenatio 2 or 3
// 			if crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[i]][3].Rank] {
// 				// scenatio 3
// 				iKIndex = 4
// 			} else {
// 				iKIndex = 2
// 			}
// 		}
// 		if crm[sOfAllFCLs[si[j]][0].Rank] == crm[sOfAllFCLs[si[j]][1].Rank] {
// 			// mst be in scenatio 2 or 3
// 			if crm[sOfAllFCLs[si[j]][2].Rank] == crm[sOfAllFCLs[si[j]][3].Rank] {
// 				// scenatio 3
// 				jKIndex = 4
// 			} else {
// 				jKIndex = 2
// 			}
// 		}

// 		// i	 	j
// 		// K aa bb	K aa bb
// 		// aa K bb  aa K bb
// 		// aa bb K  aa bb K
// 		// 9 possible scenarios, permutations of i and j

// 		var iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx int // indexes of the first card of each pair

// 		switch {
// 		// #################
// 		// K aa bb - K aa bb
// 		case iKIndex == 0 && jKIndex == 0:
// 			iA2xIdx = 1
// 			iB2xIdx = 3
// 			jA2xIdx = 1
// 			jB2xIdx = 3
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)
// 		// #################
// 		// K aa bb - aa K bb
// 		case iKIndex == 0 && jKIndex == 2:
// 			iA2xIdx = 1
// 			iB2xIdx = 3
// 			jA2xIdx = 0
// 			jB2xIdx = 3
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)
// 		// #################
// 		// K aa bb - aa bb K
// 		case iKIndex == 0 && jKIndex == 4:
// 			iA2xIdx = 1
// 			iB2xIdx = 3
// 			jA2xIdx = 0
// 			jB2xIdx = 2
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)

// 		// #################
// 		// aa K bb - K aa bb
// 		case iKIndex == 2 && jKIndex == 0:
// 			iA2xIdx = 0
// 			iB2xIdx = 3
// 			jA2xIdx = 1
// 			jB2xIdx = 3
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)
// 		// #################
// 		// aa K bb - aa K bb
// 		case iKIndex == 2 && jKIndex == 2:
// 			iA2xIdx = 0
// 			iB2xIdx = 3
// 			jA2xIdx = 0
// 			jB2xIdx = 3
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)
// 		// #################
// 		// aa K bb - aa bb K
// 		case iKIndex == 2 && jKIndex == 4:
// 			iA2xIdx = 0
// 			iB2xIdx = 3
// 			jA2xIdx = 0
// 			jB2xIdx = 2
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)

// 		// #################
// 		// aa bb K - K aa bb
// 		case iKIndex == 4 && jKIndex == 0:
// 			iA2xIdx = 0
// 			iB2xIdx = 2
// 			jA2xIdx = 1
// 			jB2xIdx = 3
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)
// 		// #################
// 		// aa bb K - aa K bb
// 		case iKIndex == 4 && jKIndex == 2:
// 			iA2xIdx = 0
// 			iB2xIdx = 2
// 			jA2xIdx = 0
// 			jB2xIdx = 3
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)
// 		// #################
// 		// aa bb K - aa bb K
// 		case iKIndex == 4 && jKIndex == 4:
// 			iA2xIdx = 0
// 			iB2xIdx = 2
// 			jA2xIdx = 0
// 			jB2xIdx = 2
// 			iltj = cmpare2x2HandsWithKickers(siPtr, i, j, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIndex, jKIndex)

// 		// #################
// 		default:
// 			err = errors.New("default case reached: ")
// 			// Info.Printf("%s\n\n", ThisFunc())
// 		}

// 		return iltj
// 	})

// 	return err
// }

// // #####################################################################

// func cmpare2x2HandsWithKickers(siPtr *[]int, iIdx, jIdx, iA2xIdx, iB2xIdx, jA2xIdx, jB2xIdx, iKIdx, jKIdx int) bool {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	iltj := false

// 	// If iA and jA are the same, and iB and jB are the same, compare kickers,
// 	// otherwise the rank of A pairs decides
// 	if crm[sOfAllFCLs[si[iIdx]][iA2xIdx].Rank] == crm[sOfAllFCLs[si[jIdx]][jA2xIdx].Rank] &&
// 		crm[sOfAllFCLs[si[iIdx]][iB2xIdx].Rank] == crm[sOfAllFCLs[si[jIdx]][jB2xIdx].Rank] {
// 		if crm[sOfAllFCLs[si[iIdx]][iKIdx].Rank] < crm[sOfAllFCLs[si[jIdx]][jKIdx].Rank] {
// 			iltj = true
// 		}
// 	} else if crm[sOfAllFCLs[si[iIdx]][iA2xIdx].Rank] == crm[sOfAllFCLs[si[jIdx]][jA2xIdx].Rank] {
// 		if crm[sOfAllFCLs[si[iIdx]][iB2xIdx].Rank] < crm[sOfAllFCLs[si[jIdx]][jB2xIdx].Rank] {
// 			iltj = true
// 		}
// 	} else if crm[sOfAllFCLs[si[iIdx]][iA2xIdx].Rank] < crm[sOfAllFCLs[si[jIdx]][jA2xIdx].Rank] {
// 		iltj = true
// 	}

// 	return iltj
// }

// // #####################################################################
// func order2xIndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false

// 		// Determine which spots the kickers are in
// 		// 1) NN k1 k1 k3
// 		// 2) k1 NN k2 k3
// 		// 3) k1 k2 NN k3
// 		// 4) k1 k2 k3 NN
// 		//
// 		// k1 > k2 > k3

// 		var i2x_idx, j2x_idx, iK1_idx, iK2_idx, iK3_idx, jK1_idx, jK2_idx, jK3_idx int

// 		switch {
// 		case crm[sOfAllFCLs[si[i]][0].Rank] == crm[sOfAllFCLs[si[i]][1].Rank]:
// 			iK1_idx = 2 // 0-5
// 			iK2_idx = 3 // 0-5
// 			iK3_idx = 4 // 0-5
// 			i2x_idx = 0 // indexes of the first card of the pairs
// 		case crm[sOfAllFCLs[si[i]][1].Rank] == crm[sOfAllFCLs[si[i]][2].Rank]:
// 			iK1_idx = 0
// 			iK2_idx = 3
// 			iK3_idx = 4
// 			i2x_idx = 1
// 		case crm[sOfAllFCLs[si[i]][2].Rank] == crm[sOfAllFCLs[si[i]][3].Rank]:
// 			iK1_idx = 0
// 			iK2_idx = 1
// 			iK3_idx = 4
// 			i2x_idx = 2
// 		case crm[sOfAllFCLs[si[i]][3].Rank] == crm[sOfAllFCLs[si[i]][4].Rank]:
// 			iK1_idx = 0
// 			iK2_idx = 1
// 			iK3_idx = 2
// 			i2x_idx = 3
// 		default:
// 			err = errors.New("default case reached: ")
// 		}

// 		switch {
// 		case crm[sOfAllFCLs[si[j]][0].Rank] == crm[sOfAllFCLs[si[j]][1].Rank]:
// 			jK1_idx = 2 // 0-5
// 			jK2_idx = 3 // 0-5
// 			jK3_idx = 4 // 0-5
// 			j2x_idx = 0
// 		case crm[sOfAllFCLs[si[j]][1].Rank] == crm[sOfAllFCLs[si[j]][2].Rank]:
// 			jK1_idx = 0
// 			jK2_idx = 3
// 			jK3_idx = 4
// 			j2x_idx = 1
// 		case crm[sOfAllFCLs[si[j]][2].Rank] == crm[sOfAllFCLs[si[j]][3].Rank]:
// 			jK1_idx = 0
// 			jK2_idx = 1
// 			jK3_idx = 4
// 			j2x_idx = 2
// 		case crm[sOfAllFCLs[si[j]][3].Rank] == crm[sOfAllFCLs[si[j]][4].Rank]:
// 			jK1_idx = 0
// 			jK2_idx = 1
// 			jK3_idx = 2
// 			j2x_idx = 3
// 		default:
// 			err = errors.New("default case reached: ")
// 		}

// 		iltj = cmpare2xHandsWithKickers(siPtr, i, j, i2x_idx, j2x_idx, iK1_idx, jK1_idx, iK2_idx, jK2_idx, iK3_idx, jK3_idx)

// 		return iltj
// 	})

// 	return err
// }

// // #####################################################################
// // cmpare2xHandsWithKickers info line
// func cmpare2xHandsWithKickers(siPtr *[]int, iIdx, jIdx, i2x_idx, j2x_idx, iK1_idx, jK1_idx, iK2_idx, jK2_idx, iK3_idx, jK3_idx int) bool {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	iltj := false

// 	// If i2x and j2x are the same, compare kickers,
// 	// otherwise the rank of 2x decides
// 	if crm[sOfAllFCLs[si[iIdx]][i2x_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][j2x_idx].Rank] {

// 		if crm[sOfAllFCLs[si[iIdx]][iK1_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK1_idx].Rank] &&
// 			crm[sOfAllFCLs[si[iIdx]][iK2_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK2_idx].Rank] {
// 			if crm[sOfAllFCLs[si[iIdx]][iK3_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][jK3_idx].Rank] {
// 				iltj = true
// 			}
// 		} else if crm[sOfAllFCLs[si[iIdx]][iK1_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK1_idx].Rank] {
// 			if crm[sOfAllFCLs[si[iIdx]][iK2_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][jK2_idx].Rank] {
// 				iltj = true
// 			}
// 		} else if crm[sOfAllFCLs[si[iIdx]][iK1_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][jK1_idx].Rank] {
// 			iltj = true
// 		}
// 	} else if crm[sOfAllFCLs[si[iIdx]][i2x_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][j2x_idx].Rank] {
// 		iltj = true
// 	}
// 	return iltj
// }

// // #####################################################################
// func orderHCIndexesAsc(siPtr *[]int) (err error) {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	// Remember, this is the "is i less than j" function for sort, sorting in ascending order.
// 	sort.SliceStable(si, func(i, j int) bool {
// 		iltj := false

// 		iltj = cmpareHCHands(siPtr, i, j, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4)

// 		return iltj
// 	})

// 	return err
// }

// // #####################################################################
// // cmpare2xHandsWithKickers info line
// func cmpareHCHands(siPtr *[]int, iIdx, jIdx, iK1_idx, jK1_idx, iK2_idx, jK2_idx, iK3_idx, jK3_idx, iK4_idx, jK4_idx, iK5_idx, jK5_idx int) bool {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	si := *siPtr // Slice of Int

// 	iltj := false

// 	if crm[sOfAllFCLs[si[iIdx]][iK1_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK1_idx].Rank] &&
// 		crm[sOfAllFCLs[si[iIdx]][iK2_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK2_idx].Rank] &&
// 		crm[sOfAllFCLs[si[iIdx]][iK3_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK3_idx].Rank] &&
// 		crm[sOfAllFCLs[si[iIdx]][iK4_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK4_idx].Rank] {
// 		if crm[sOfAllFCLs[si[iIdx]][iK5_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][jK5_idx].Rank] {
// 			iltj = true
// 		}
// 	} else if crm[sOfAllFCLs[si[iIdx]][iK1_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK1_idx].Rank] &&
// 		crm[sOfAllFCLs[si[iIdx]][iK2_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK2_idx].Rank] &&
// 		crm[sOfAllFCLs[si[iIdx]][iK3_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK3_idx].Rank] {
// 		if crm[sOfAllFCLs[si[iIdx]][iK4_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][jK4_idx].Rank] {
// 			iltj = true
// 		}
// 	} else if crm[sOfAllFCLs[si[iIdx]][iK1_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK1_idx].Rank] &&
// 		crm[sOfAllFCLs[si[iIdx]][iK2_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK2_idx].Rank] {
// 		if crm[sOfAllFCLs[si[iIdx]][iK3_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][jK3_idx].Rank] {
// 			iltj = true
// 		}
// 	} else if crm[sOfAllFCLs[si[iIdx]][iK1_idx].Rank] == crm[sOfAllFCLs[si[jIdx]][jK1_idx].Rank] {
// 		if crm[sOfAllFCLs[si[iIdx]][iK2_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][jK2_idx].Rank] {
// 			iltj = true
// 		}
// 	} else if crm[sOfAllFCLs[si[iIdx]][iK1_idx].Rank] < crm[sOfAllFCLs[si[jIdx]][jK1_idx].Rank] {
// 		iltj = true
// 	}

// 	return iltj
// }

// // #########################################################################
// // genEquivalentSFHandList returns a listing of five card hands, excluding suits, sorted by hand rank,
// // with counts and hand kind name and kind rank.
// func genEquivalentSFHandList(siPtr *[]int) (efchlPtr *equivalentFCHList) {

// 	// var efchl equivalentFCHList
// 	efchl := make(equivalentFCHList, 10)
// 	efchlPtr = &efchl

// 	si := *siPtr // Slice of Int
// 	j := 0       // inde for efchl
// 	var efch equivalentFiveCardHand

// 	// Load up the first entry
// 	efch.c1r = sOfAllFCLs[si[0]][0].Rank
// 	efch.c2r = sOfAllFCLs[si[0]][1].Rank
// 	efch.c3r = sOfAllFCLs[si[0]][2].Rank
// 	efch.c4r = sOfAllFCLs[si[0]][3].Rank
// 	efch.c5r = sOfAllFCLs[si[0]][4].Rank
// 	efch.count = 1
// 	efch.info.handKind = fchkrFL.SfInfo.handKind
// 	efch.info.typeRank = fchkrFL.SfInfo.typeRank

// 	// var efchl[j] equivalentFiveCardHand
// 	efchl[j] = efch

// 	// For entry 1-end, take the value (index within the All Five Card List), strip the suits, then count.
// 	for i, iAfcl := range si[1:] {
// 		i++
// 		fmt.Println("i: ", i, " iAfcl: ", iAfcl)

// 		// Compare this han to previous. If same, just count up. If not, start next entry.
// 		if sOfAllFCLs[iAfcl][0].Rank == sOfAllFCLs[si[i-1]][0].Rank &&
// 			sOfAllFCLs[iAfcl][1].Rank == sOfAllFCLs[si[i-1]][1].Rank &&
// 			sOfAllFCLs[iAfcl][2].Rank == sOfAllFCLs[si[i-1]][2].Rank &&
// 			sOfAllFCLs[iAfcl][3].Rank == sOfAllFCLs[si[i-1]][3].Rank &&
// 			sOfAllFCLs[iAfcl][4].Rank == sOfAllFCLs[si[i-1]][4].Rank {

// 			efchl[j].count++
// 		} else {
// 			j++

// 			efch.c1r = sOfAllFCLs[iAfcl][0].Rank
// 			efch.c2r = sOfAllFCLs[iAfcl][1].Rank
// 			efch.c3r = sOfAllFCLs[iAfcl][2].Rank
// 			efch.c4r = sOfAllFCLs[iAfcl][3].Rank
// 			efch.c5r = sOfAllFCLs[iAfcl][4].Rank
// 			efch.count = 1
// 			efch.info.handKind = fchkrFL.SfInfo.handKind
// 			efch.info.typeRank = fchkrFL.SfInfo.typeRank

// 			efchl[j] = efch
// 		}
// 	}
// 	return efchlPtr
// }

// // #########################################################################
// // printSOEH info line
// func printSOEH(soehPtr *equivalentFCHList) error {
// 	// Info.Printf("%s\n\n", ThisFunc())

// 	// ssi := *ssiPtr
// 	// soeh := *soehPtr

// 	for i, v := range *soehPtr {

// 		fmt.Println("i: ", i, " ", v)
// 	}
// 	return nil
// }

// // #########################################################################
// // genEquivalent4xHandList returns a listing of five card hands, excluding suits, sorted by hand rank,
// // with counts and hand kind name and kind rank.
// func genEquivalent4xHandList(siPtr *[]int) (efchlPtr *equivalentFCHList) {

// 	// var efchl equivalentFCHList
// 	efchl := make(equivalentFCHList, 10)
// 	efchlPtr = &efchl

// 	si := *siPtr // Slice of Int
// 	j := 0       // inde for efchl
// 	var efch equivalentFiveCardHand

// 	// Load up the first entry
// 	efch.c1r = sOfAllFCLs[si[0]][0].Rank
// 	efch.c2r = sOfAllFCLs[si[0]][1].Rank
// 	efch.c3r = sOfAllFCLs[si[0]][2].Rank
// 	efch.c4r = sOfAllFCLs[si[0]][3].Rank
// 	efch.c5r = sOfAllFCLs[si[0]][4].Rank
// 	efch.count = 1
// 	efch.info.handKind = fchkrFL.X4Info.handKind
// 	efch.info.typeRank = fchkrFL.X4Info.typeRank

// 	// var efchl[j] equivalentFiveCardHand
// 	efchl[j] = efch

// 	// For entry 1-end, take the value (index within the All Five Card List), strip the suits, then count.
// 	for i, iAfcl := range si[1:] {
// 		i++
// 		fmt.Println("i: ", i, " iAfcl: ", iAfcl)

// 		// Compare this han to previous. If same, just count up. If not, start next entry.
// 		if sOfAllFCLs[iAfcl][0].Rank == sOfAllFCLs[si[i-1]][0].Rank &&
// 			sOfAllFCLs[iAfcl][1].Rank == sOfAllFCLs[si[i-1]][1].Rank &&
// 			sOfAllFCLs[iAfcl][2].Rank == sOfAllFCLs[si[i-1]][2].Rank &&
// 			sOfAllFCLs[iAfcl][3].Rank == sOfAllFCLs[si[i-1]][3].Rank &&
// 			sOfAllFCLs[iAfcl][4].Rank == sOfAllFCLs[si[i-1]][4].Rank {

// 			efchl[j].count++
// 		} else {
// 			// j++

// 			efch.c1r = sOfAllFCLs[iAfcl][0].Rank
// 			efch.c2r = sOfAllFCLs[iAfcl][1].Rank
// 			efch.c3r = sOfAllFCLs[iAfcl][2].Rank
// 			efch.c4r = sOfAllFCLs[iAfcl][3].Rank
// 			efch.c5r = sOfAllFCLs[iAfcl][4].Rank
// 			efch.count = 1
// 			efch.info.handKind = fchkrFL.X4Info.handKind
// 			efch.info.typeRank = fchkrFL.X4Info.typeRank

// 			// efchl[j] = efch
// 			efchl = append(efchl, efch)

// 		}
// 	}
// 	return efchlPtr
// }
