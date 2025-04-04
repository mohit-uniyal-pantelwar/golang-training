package main

import (
	"fmt"
	"math"
)

// Interfaces: A type that contains method signatures.
// Any value may belong to interface if it implements the methods signature mentioned in interface.
type ShapeMethodsImpl interface {
	getArea() float32
	getPerimeter() float32
}

type rectangle struct {
	length  float32
	breadth float32
}

type circle struct {
	radius float32
}

func (r rectangle) getArea() float32 {
	return r.length * r.breadth
}

func (r rectangle) getPerimeter() float32 {
	return 2 * (r.length + r.breadth)
}

func (c circle) getArea() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getPerimeter() float32 {
	return 2 * math.Pi * c.radius
}

func printArea(s ShapeMethodsImpl) {
	fmt.Printf("Area: %v\n", s.getArea())
}

func printPerimeter(s ShapeMethodsImpl) {
	fmt.Printf("Perimeter: %v\n", s.getPerimeter())
}

func main() {

	circle1 := circle{
		radius: 2,
	}

	rectangle1 := rectangle{
		length:  4,
		breadth: 5,
	}

	fmt.Println("Rectangle")
	printArea(rectangle1)
	printPerimeter(rectangle1)
	fmt.Println("")
	fmt.Println("circle")
	printArea(circle1)
	printPerimeter(circle1)

}
