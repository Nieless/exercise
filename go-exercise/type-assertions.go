package main

import "fmt"

func main(){
	var i interface{} = "hello"

	s := i.(string)
	fmt.Printf(s)

	var a interface{} = 50
	b := a.(int)
	fmt.Printf(b)
}
