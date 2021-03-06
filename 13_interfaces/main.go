package main

import (
	"fmt"
	"math"
)

// Define interface

// Shape asdfasdf
type Shape interface {
	area() float64
}

// Circle asdfasf
type Circle struct {
	x, y, radius float64
}

// Rectangle asdf as
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))

}
