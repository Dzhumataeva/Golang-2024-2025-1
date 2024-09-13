package main

import (
    "fmt"
    "math"
)

// Shape interface with a method Area
type Shape interface {
    Area() float64
}
// Circle struct with a radius
type Circle struct {
    Radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Rectangle struct with width and height
type Rectangle struct {
    Width, Height float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
// PrintArea function takes a Shape and prints its area
func PrintArea(s Shape) {
    fmt.Printf("The area is: %.2f\n", s.Area())
}

func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 3}

    PrintArea(c) 
    PrintArea(r) 
}

