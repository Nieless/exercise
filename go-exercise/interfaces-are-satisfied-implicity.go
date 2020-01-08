package main

import "fmt"

type I interface{
	M()
}

type T struct{
	s string
}

func (t T)M(){
	fmt.Println(t.s)
}

func main(){
	var s I = T{"string"}
	s.M()
}
