package main

import "fmt"

type address struct{
	city string
	state string
}

type Person struct{
	firstname string
	lastname string
	address
}

func (a Person) fullAddress(){
	fmt.Printf("%s %s\n", a.firstname, a.address.state)
}

func main(){
	p := Person{
		firstname: "Nilesh",
		lastname: "Patel",
		address : address{
			city: "Ahemedabad",
			state: "Gujarat",
		},
	}
	p.fullAddress()
}
