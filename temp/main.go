package main

import (
	"fmt"
	"math"
)

// #######################
type Abser interface {
	Abs() float64
}

// #######################
const (
	A float64 = 1.1111111
	B float64 = 2.2222222222
	C float64 = 3.333333
	D float64 = 4.4444444
	E float64 = 5.5555555
)

// #######################
type MyFloat float64

type Vertex struct {
	X, Y float64
}

type Cnst struct {
	a, b, c, d, e, f float64
}

// #######################
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// #######################
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// #######################
func (c Cnst) Abs() float64 {
	return math.Sqrt(c.a + c.b + c.c + c.d)
}

/*
####################################################################################
####################################################################################
####################################################################################
*/

// #######################
func main() {
	var f Abser
	var v Abser
	var x Abser

	f = MyFloat(-math.Sqrt2)
	v = Vertex{3, 4}
	x = Cnst{A, B, C, D, E, E}

	// a = f  // a MyFloat implements Abser, since func Abs has a fnction definition with MyFloat as input
	// a = &v // a *Vertex implements Abser, since func Abs is also defined with *Vertex as input

	fmt.Println(f.Abs())
	fmt.Println(v.Abs())
	fmt.Println(x.Abs())
}
