package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

// implement he interface for circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Define another struct type
type Rectangle struct {
	Width, Height float64
}

// Implement the interface for rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printShapeInfo(s Shape) {
	switch shape := s.(type) {
	case Circle:
		fmt.Println("Circle with radius is:", shape.Radius)
	case Rectangle:
		fmt.Println("Rectangle with width", shape.Width, "And Height", shape.Height)
	default:
		fmt.Println("Unknown shape")
	}
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	printShapeInfo(c)
	printShapeInfo(r)
}
