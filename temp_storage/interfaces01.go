package main

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
func (r rectangle) area() (float64, error) {
	return r.height * r.width, nil
}

// #################################
func (r rectangle) perimeter() (float64, error) {
	return 2*r.height + 2*r.width, nil
}

// #################################
func (c circle) area() (float64, error) {
	return math.Pi * c.radius * c.radius, nil
}

// #################################
func (c circle) perimeter() (float64, error) {
	return 2 * math.Pi * c.radius, nil
}

// #################################
// func performAllGeometricMeasurements(gmsItf geometricMeasurements) {
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

	performAllGeometricMeasurements(r)
	performAllGeometricMeasurements(c)

}
