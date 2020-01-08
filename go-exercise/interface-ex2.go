package main

import "fmt"

type SalaryCalculator interface{
	DisplaySalary()
}

type LeaveCalculator interface{
	LeaveCalc() int
}

type Employee struct{
	name string
	basicpay int
	pf int
	totalleaves int
	leavestaken int
}

func (e Employee)DisplaySalary(){
	fmt.Printf("%s %d\n", e.name, e.basicpay + e.pf)
}

func (e Employee)LeaveCalc()int{
	return e.totalleaves - e.leavestaken
}

func main(){
	var s SalaryCalculator
	e := Employee{
		name:"Neel",
		basicpay:6000,
		pf:500,
		totalleaves:25,
		leavestaken:12,
	}
	s = e
	s.DisplaySalary()
	var l LeaveCalculator
	l = e
	fmt.Printf("your left leaves:%d\n",l.LeaveCalc())
}
