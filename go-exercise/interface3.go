package main

import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
}

type rect struct{
	length, width float64
}

type circle struct{
	radius float64
}

func main(){
	r := rect{
		width: 2.0,
		length: 5.0,
		}
	c := circle{radius:5.0}

	measure(r)
	measure(c)
}

func (r rect) area() float64{
	return r.width * r.length
}

func (r rect) perim() float64{
	return 2*r.width + 2*r.length
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64{
	return 2* math.Pi * c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
