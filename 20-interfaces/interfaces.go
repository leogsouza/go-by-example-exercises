// Interfaces are named collections of methods signatures
package main

import "fmt"
import "math"

// Here's a basic interface for geografic shapes
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on rect and circle types
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

/* To implment an interface in Go, we just need to implment all the methods
in the interface. He we implment geomery in rects */
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/* If a variable has an interface type, then we can call methods that are in
the named interface. Here's a generic measure function taking advantages of
this to work on any geometry */
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	/* The circle and rect struct types both implment geometry interface so we
	can use intances of these structs as arguments to measure */
	measure(r)
	measure(c)
}
