package main

import (
	"fmt"
	"math"
)

type circle struct {
	x, y, r float64
}

type rectangle struct {
	x1, x2, y1, y2 float64
}

c := circle{0, 0, 5}
r := rectangle{0, 0, 10, 10}

func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *rectangle) area() float64 {
	l := distance(r.x1, r.x2, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y2)
	return l * w
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	fmt.Println(с.area())
	fmt.Println(r.area())
}
