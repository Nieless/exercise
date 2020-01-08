package main

import "fmt"

type Describer interface{
	Describe()
}

type Person struct{
	name string
	age int
}

// value receiver
func (p Person)Describe(){
	fmt.Printf("%s %d\n", p.name, p.age)
}

type Address struct{
	city string
	state string
}

// pointer receiver
func (a *Address)Describe(){
	fmt.Printf("%s %s\n", a.city, a.state)
}

func main(){
	// for value receiver method
	fmt.Println("for value receiver method")
	var d1 Describer
	p1 := Person{"Nil",25}
	d1 = p1
	d1.Describe()

	p2 := Person{"Neel",24}
	d1 = &p2
	d1.Describe() 

	// for pointer receiver method
	fmt.Println("for pointer receiver method")
	var d2 Describer

	/*
	-- This is not allowed, type implement interface here as method has pointer reciever, 
	a1 := Address{"ahm","gj"}
	d2 = a1
	d2.Describe()
	*/

	a2 := Address{"hmt","gj"}
	d2 = &a2
	d2.Describe()
}

