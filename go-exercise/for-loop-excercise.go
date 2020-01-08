package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, d := 1.0, 1.0
	fmt.Println(1e-15)
	for d > 1e-15 {
	    z0 := z
	    fmt.Print("dsadasd")
	    z = z - (z*z -x ) /(2*z)
	    //fmt.Printf('c %v', z)
	    d = math.Abs(z - z0)
	    //fmt.Printf('d %v', d)
	}
	fmt.Println(z)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

