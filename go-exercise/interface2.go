package main

import (
	"fmt"
	)

type I interface {
	M1()
}

type T int


func (t T) M1(){
	fmt.Println(t)
}

func main(){
	var i I = T(20)
	i.M1()
	i = T(25)
	i.M1()
}
