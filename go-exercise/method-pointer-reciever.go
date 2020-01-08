package main

import "fmt"


type Employee struct{
	name string
	id int
	role string
}

func (e Employee) changeName(newName string){
	e.name = newName
}

func (e *Employee) changeNname(newName string){
	e.name = newName
}

func main(){
	e := Employee{
		name: "N",
		id: 1,
		role: "kk",
	}

	e.changeName("NN")
	fmt.Printf("%s\n",e.name)
	e.changeNname("NNN")
	fmt.Printf("%s\n",e.name)
}
