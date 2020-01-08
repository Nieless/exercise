package main

import "fmt"

type I interface{}

func main(){
	var i I
	describe(i)

	i = 42
	describe(i)

	i = "aa"
	describe(i)
}

func describe(i I){
	fmt.Printf("%v\n", i)
}
