// package main

import (
	"fmt"
	"math"
)

// #################################
type areaCalculator interface {
	area() (float64, error)
}

type perimeterCalculator interface {
	perimeter() (float64, error)
}

type shapeMeasurementCalculator interface {
	areaCalculator
	perimeterCalculator
}

/*
type geometricMeasurements interface {
	area() (float64, error)
	perimeter() (float64, error)
}
*/

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

// #################################
func (rPtr *rectangle) area() (float64, error) {
	return rPtr.height * rPtr.width, nil
}

// #################################
func (rPtr *rectangle) perimeter() (float64, error) {
	return 2*rPtr.height + 2*rPtr.width, nil
}

// #################################
func (cPtr *circle) area() (float64, error) {
	return math.Pi * cPtr.radius * cPtr.radius, nil
}

// #################################
func (cPtr *circle) perimeter() (float64, error) {
	return 2 * math.Pi * cPtr.radius, nil
}

// #################################
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
