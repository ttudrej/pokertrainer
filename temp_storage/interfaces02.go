package main

import (
	"fmt"
)

type WriterItf interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// ###########################################
func (cw ConsoleWriter) Write(bs []byte) (int, error) {
	n, err := fmt.Println(string(bs))
	return n, err
}

// ###########################################
func main() {
	fmt.Println("Hello")

	var w WriterItf = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

}
