package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	withd, height float64
}
type circle struct {
	radius float64
}

func (r rect)area() float64  {
	return r.withd * r.height
}
func (c circle)area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64  {
	return r.height * 2 + r.withd * 2
}
func (c circle) perim() float64  {
	return 2 * math.Pi * c.radius
}
func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main()  {
	r := rect{3, 5}
	c := circle{3.4}

	measure(r)
	measure(c)

}