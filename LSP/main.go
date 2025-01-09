package main

import "fmt"

// Shape interface defines a common behavior for geometric shapes.
type Shape interface {
	Area() float64
}

// Rectangle shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Square shape
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Function to print the area of any shape
func printArea(shape Shape) {
	fmt.Printf("Area: %.2f\n", shape.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	square := Square{Side: 4}

	printArea(rect)   // Works with Rectangle
	printArea(square) // Works with Square
}
