package main

import (
	"fmt"
	"math"
)

// Define an Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct type
type Circle struct {
	Radius float64
}

// Implement the interface for circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Define another struct type
type Rectangle struct {
	Width, Height float64
}

// Impelement the interfaces for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	//Create instances of circle and Rectangle
	c := Circle{Radius: 5}
	fmt.Println("C is:", c)
	r := Rectangle{Width: 4, Height: 6}
	fmt.Println("R is:", r)

	//Create a slice of Shape interface
	shapes := []Shape{c, r}
	fmt.Println("Shapes is:", shapes)

	//Iterate over the slice and print the Area and Perimeter
	for _, shape := range shapes {
		fmt.Printf("Shape: %T, Area: %.2f, Perimeter: %.2f\n", shape, shape.Area(), shape.Perimeter())
	}
}
