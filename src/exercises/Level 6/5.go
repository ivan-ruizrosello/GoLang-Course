package main

import (
	"fmt"
	"math"
)

type square struct {
	l, w float64
}

type circle struct {
	rad float64
}

func (c circle) getArea() float64 {
	return math.Pi * (c.rad * c.rad)
}
func (s square) getArea() float64 {
	return s.l * s.w
}

type shape interface {
	getArea() float64
}

func main() {
	circle := circle{1}
	square := square{5,5}

	info(circle)
	info(square)
}

func info(s shape) {
	fmt.Printf("The area is -> %v\n", s.getArea())
}
