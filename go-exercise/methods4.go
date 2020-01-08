package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	x, y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.x*v.y)
}

func main(){
	v := Vertex{
		x: 3,
		y: 5,
	}

	fmt.Println(v.Abs())
}
