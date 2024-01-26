package main

import (
	"fmt"
	"math"
)

// an interface is a collection of method signatures that define a set of behaviors
// unlike some other object-oriented languages, Go interfaces are implicitly implemented
// meaning a type satisfies an interface if it implements all the methods 
// declared by that interface, without explicitly declaring to do so

type shape interface {
	area() float64
}

func getArea(f shape) {
	fmt.Printf("The area of the shape is %0.2f\n", f.area())
}

type rectangle struct {
	height  float64
	width float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := rectangle{10, 15}
	getArea(r)

	c := circle{10}
	getArea(c)

}