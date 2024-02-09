package shapes

import (
	"math"
)

// Shape represents a geometrical shape
type Shape interface {
	area() float64
}

// Rectangle is a struct that contains height and width
type Rectangle struct {
	Height  float64
	Width float64
}

// Area returns the rectangle's area
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle represents a struct that contains radius
type Circle struct {
	Radius float64
}

// Area returns the circles's area
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}