package main

import "fmt"

type I interface{
	walk()
}

type A struct {}

func (a A) walk() {}

type B struct {}

func (b B) walk() {}

func main(){
	var i I
	i = A{}
	fmt.Printf("%T\n", i.(A)) //dynamic type of i is A

	

}
