package main

import "fmt"

// interface - set of methods
type I interface {
	m1()
}

type T struct{}

func (T) m1() {}

func main(){
	var i I = T{}
	fmt.Println(i)
}


