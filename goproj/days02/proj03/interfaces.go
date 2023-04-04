package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectc struct {
	width,height float64
}
type circle struct {
	radius float64
}

func(r rectc) area() float64{
	return r.width * r.height
}

func (r rectc) perim() float64{
	return 2*r.width + 2*r.height
}

func(c circle) area() float64{
	return math.Pi *c.radius *c.radius
}

func (c circle) perim() float64{
	return 2* math.Pi *c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println("面积是:",g.area())
	fmt.Println("周长是:",g.perim())
}

func main() {
	r:= rectc{width: 3,height: 4}
	c:= circle{radius: 5}

	measure(r)

	measure(c)

}
