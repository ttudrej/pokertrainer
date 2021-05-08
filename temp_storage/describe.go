// package main

import (
	"fmt"
	"math"
)

// ######################################
type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

// ######################################

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ######################################
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ######################################
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ######################################
func main() {

	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type. For example,
	// fmt.Print takes any number of arguments of type interface{}.

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	b := v

	fmt.Println(a.Abs())
	fmt.Println(b.Abs())

}
