package main

import "fmt"

type Person struct{
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.name, p.age)
}

func main(){
	a := Person{"neel",25}
	fmt.Println(a)
}
