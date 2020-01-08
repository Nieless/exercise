package main

import "fmt"

type SalaryCalculator interface{
	DisplaySalary()
}

type LeaveCalculator interface{
	LeavesCalc() int
}

type EmployeeOperations interface{
	SalaryCalculator
	LeaveCalculator
}

type Employee struct{
	name string
	basicpay int
	pf int
	totalleaves int
	leavestaken int
}

func (e Employee) DisplaySalary(){
	fmt.Printf("%s=%d\n", e.name, e.basicpay + e.pf)
}

func (e Employee) LeavesCalc()int {
	return e.totalleaves - e.leavestaken
}

func main(){
	var empOp EmployeeOperations
	e := Employee {
		name: "Neel",
		basicpay: 5000,
		pf: 200,
		totalleaves: 24,
		leavestaken: 5,
	}
	empOp = e
	empOp.DisplaySalary()

	fmt.Println("\nleaves left=", empOp.LeavesCalc())
}
