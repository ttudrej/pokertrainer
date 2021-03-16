package main

import "fmt"

type IntCounter int

type Incrementer interface {
	Increment() int
}

func (icPtr *IntCounter) Increment() int {

	*icPtr++
	return int(*icPtr)
}

// ########################################
func main() {

	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

}
