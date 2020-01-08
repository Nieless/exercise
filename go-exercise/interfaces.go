package main

import (
	"fmt"
	"math"
	)

type Absolute interface{
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs()float64{
	if f < 0{
		return float64(-f)
	}
	return float64(-f)
}

type Vertex struct{
	x, y float64
}

func (v *Vertex) Abs()float64{
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func calc(a Absolute){
	fmt.Println(a.Abs())
}

func main(){
	mf := MyFloat(-math.Sqrt2)
	v := Vertex{5,4}
	calc(&v)
	calc(mf)
}
