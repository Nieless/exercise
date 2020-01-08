package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	x, y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.x * v.y)
}

func (v *Vertex) Scale(f float64){
	v.x = v.x + f
	v.y = v.y +f
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v)
	(&v).Scale(10)

	// pointer used to update values of struct variable
	p := &v
	p.Scale(10)

	fmt.Println(v.Abs())
	fmt.Println(v)
}
