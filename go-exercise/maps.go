package main

import "fmt"

type vertex struct {
	lat, long float64
}

func main(){
	m := make(map[string]vertex)

	m["i"] = vertex{
		-4.2, 73.24,
	}
	fmt.Println(m["i"])

}
