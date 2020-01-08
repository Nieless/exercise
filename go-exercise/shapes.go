package main

import (
	"fmt"
	"math"
)

type rect struct{
	a, b float64
}

type triangle struct{
	a, b, c float64
}

type circle struct{
	radius float64
}

func (r rect) Area() float64{
	return r.a * r.b
}

func (t triangle) Area() float64{
	return math.Pi * t.a + t.b + t.c
}

func (c circle) Area() float64{
	return 2 * c.radius
}

type Shape interface {
	Area() float64
}

func (t triangle) Angles() []float64{
	return []float64{angle(t.a, t.c, t.a), angle(t.a, t.c, t.b), angle(t.a, t.b, t.c)}
}

func angle(a, b, c float64) float64{
	return math.Acos((a*a+b*b-c*c)/(2*a*b)) * 180.0 / math.Pi
}

func main(){
	shapes := []Shape{
		rect{5,4},
		triangle{5,4,3},
		circle{1},
	}
	for _,v := range shapes {
		fmt.Printf("Area of %T: %f\n",v, v.Area())
		if t, ok := v.(triangle); ok {
			fmt.Println(t.Area(), t.Angles())
		}
	}
}
