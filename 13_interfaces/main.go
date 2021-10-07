package main

import (
	"fmt"
	"math"
)

// Interface: list of methods that can be implemented on different structs
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

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
	circle := Circle{x: 0, y: 0, radius: 10}
	rectangle := Rectangle{height: 18, width: 20}

	fmt.Println("Circle Area is: ", getArea(circle))
	fmt.Println("Rectangle Area is: ", getArea(rectangle))
}