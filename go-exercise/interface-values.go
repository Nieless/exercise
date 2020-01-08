package main

import "fmt"

type I interface{
	M() string
}

type T struct{
	S string
}

func (t *T)M() string{
	fmt.Println(t.S)
	return t.S
}

func main(){
	var i I = &T{"Nilesh"}
	describe(i)
}

func describe(i I){
	i.M()
	fmt.Printf("%T \n", i)
}
