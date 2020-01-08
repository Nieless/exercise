package main

import (
	"fmt"
)


type Employee struct {
	name string
	salary int
	currency string
}

func (e Employee) displaySalary(){
	fmt.Println(e)
}

func main() {
	emp1 := Employee {
		name: "xyz",
		salary: 25000,
		currency : "rs",
	}
	emp1.displaySalary()
	emp1.displayName()

}

func (e Employee) displayName(){
	fmt.Println(e.name)
}
