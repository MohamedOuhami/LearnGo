package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectv2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectv2) area() float64 {
	return r.width * r.height
}

func (r rectv2) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle with Radius ", c.radius)
	} else {
		fmt.Println("Not a circle")
	}
}

func learnInterface() {
	r1 := rectv2{2, 4}
	c1 := circle{2}

	measure(r1)
	measure(c1)
	detectCircle(r1)
	detectCircle(c1)
}
