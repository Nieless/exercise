package main

import "fmt"

type I1 interface {
	M1()
}

type I2 interface {
	M2()
}

type T struct{}

func (t T) M1(){
	fmt.Println("t.m1")
}

func (t T) M2(){
	fmt.Println("t.m2")
}

func f1(i I1){
	i.M1()
}
func f2(i I2){
	i.M2()
}

func main(){
	t := T{}
	//t.M1()
	//t.M2()
	//f1(t)
	//f2(t)
	f1(t)
	f2(t)
}
