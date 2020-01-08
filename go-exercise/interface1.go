package main

import "fmt"

type I interface{
	M()
}

type T1 struct{}

type T2 struct{}

func (T1) M(){
	fmt.Println("T1")
}

func (T2) M(){
	fmt.Println("T2")
}

func main(){
	// here I is static type of variable i and T1 is dynamic type
	var i I = T1{}
	i.M()
	i = T2{}
	i.M()
	_ = i
	fmt.Println(i)
}
