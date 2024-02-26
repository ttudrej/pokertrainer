// package main

import (
	"fmt"
	"math"
)

// Interface names
// By convention, one-method interfaces are named by the method name plus
// an -er suffix or similar modification to construct an agent noun: Reader,
// Writer, Formatter, CloseNotifier etc.

// #################################
type rectangle struct { // (the value type)
	height float64
	width  float64
}
type circle struct { // (the value type)
	radius float64
}

// #################################
// An interface type is defined as a set of method (func?) signatures.
// A value of interface type can hold any value that implements those methods.
// Interface types implement methods
// Interface types implement methods of the same name but with different inputs signatures but same output signature?
// Single method Interface types implement method
// This is a one-method interface
type areaCalculator interface { // can hold any value(funs are values too) that implements any of the methods
	area() (float64, error)
}

// A receiver in a method signature is a references to a struct( type somename struct {} )
// You can have many methods with the same name, for referencing various structs.
// It's the methods job to know how to perform the operation (suggested by the method name, area),
// on the receiver type/struct it is given.

// Functions are values too. They can be passed around just like other values.
// Function values may be used as function arguments and return values.

// You could say that the "method signatures" wrap "functin signatures".
// A method IS a functin, but with a "receiver argument".

// funcs that serve an interface, use the same func name, "area"

// A method is a function with a special receiver argument.
func (rPtr *rectangle) area() (float64, error) { // area method (func with a receiver argument) has a receiver of *rectangle type
	return rPtr.height * rPtr.width, nil // We can say, this method operates on type *rectangle
}
func (cPtr *circle) area() (float64, error) { // area method/func is defined on *circle type
	return math.Pi * cPtr.radius * cPtr.radius, nil
}

// Methods with pointer receivers can modify the value to which the
// receiver points (as Scale does here). Since methods often need to
// modify their receiver, pointer receivers are more common than
// value receivers.

// 4 things are needed to interfaces instead of functins directly:
//
// 1) One or more structs, sa, sb, sc, ... (type mystructname struct {})
// 2) One or more methods (functins with a rceiver argument), and same functin name, where the reciver
//		references one of the structs, sa, sb, sc, .... (func (<receiver-structref>) funcname(<inputs>) (<outputs>))
// 3) An Interface definition, which pools/ties the functions with the same name but
// 		different signatures.
// 4) Something that uses the interface.

// There are two reasons to use a pointer receiver.
// The first is so that the method can modify the value that its receiver
// points to.
// The second is to avoid copying the value on each method call. This can
// be more efficient if the receiver is a large struct, for example.
// In general, all methods on a given type should have either value or
// pointer receivers, but not a mixture of both.

// Terminology
// For any set of methods, so funcs with receivers, we can say that:
// "method's receiver implements an interface", if/when that interface's definition
// includes a function name and signature of our method.
//
// So for "func (rPtr *rectangle) area() (float64, error)" method, we say:
// "*rectangle" receiver implements interface areaCalculator
// because areCalculator interface contains "area() (float64, error)" func
// signature.

// #################################
type perimeterCalculator interface {
	perimeter() (float64, error)
}

func (rPtr *rectangle) perimeter() (float64, error) {
	return 2*rPtr.height + 2*rPtr.width, nil
}
func (cPtr *circle) perimeter() (float64, error) {
	return 2 * math.Pi * cPtr.radius, nil
}

type shapeMeasurementCalculator interface {
	areaCalculator
	perimeterCalculator
}

// ################################# smcItf: shape measurement calculator interface type float?
func performAllGeometricMeasurements(smcItf shapeMeasurementCalculator) {
	describe(smcItf)
	fmt.Println(smcItf.area())
	fmt.Println(smcItf.perimeter())
}

// #################################
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// #################################
func main() {

	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	rPtr := &r
	cPtr := &c

	performAllGeometricMeasurements(rPtr)
	performAllGeometricMeasurements(cPtr)

	ra, _ := rPtr.area()

	fmt.Println("rec area: ", ra)

}
