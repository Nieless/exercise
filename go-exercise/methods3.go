package main

import "fmt"

type rectangle struct {
	length int
	width int
}

func area(r rectangle){
	fmt.Println(r.length*r.width)
}

func (r rectangle) area(){
	fmt.Println(r.length*r.width)
}

func main(){
	r := rectangle{
		length: 10,
		width: 10,
	}

	area(r)
	r.area()

	p := &r

	p.area()
}
