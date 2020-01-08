package main

import "fmt"

type person struct {
    name string
    age int
}

func main(){
	fmt.Println(person{"Nilesh", 20})
	fmt.Println(person{name: "Nilesh", age:25})
	s := person{"Nilesh",25}
	d := &s
	fmt.Println(s.name, d.age)
}
