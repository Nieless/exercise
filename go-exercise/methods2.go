package main

import (
	"fmt"
)


type Employee struct{
	name string
	id int
	role string
}

func displayId(e Employee){
	fmt.Printf("%d\n",e.id)
}

func main(){
	emp := Employee{
		name: "N",
		id: 1,
		role: "Programmer",
	}

	displayId(emp)
}
