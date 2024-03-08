package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"
)

// type VSlice []int
type VSlice [][]int

func (vs VSlice) String() string {
	var str string
	for index, v := range vs {
		// str += fmt.Sprintf("%d\n", i)
		str += fmt.Sprintf("%v : %d\n", v, index+1)
	}
	return str
}

func main() {
	fmt.Println("hello")

	n := 7
	k := 5

	// Get all possible combinations for n,k
	sl_of_sl_of_ints := combin.Combinations(n, k)

	// Find how many of the above there are.
	bc := combin.Binomial(n, k)

	fmt.Printf("sl of sl of ints: %T\n", sl_of_sl_of_ints)
	fmt.Printf("sl of sl of ints: %v\n\n", sl_of_sl_of_ints)
	fmt.Printf("binomial coeff: %v\n", bc)

	// slice := []int{1, 2, 3}
	// fmt.Print(VSlice(slice))
	fmt.Print(VSlice(sl_of_sl_of_ints))
}
