package main

import "fmt"

func main() {
	var x interface{} = 7
	a, ok := x.(int)
	fmt.Println(a, ok)
}
