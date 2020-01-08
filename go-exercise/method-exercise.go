package main

import (
	"fmt"
	)

type rectangle struct{
	length int
	width int
}


// function
func area(r rectangle){
	fmt.Printf("Rectangle result,,,:%d\n", r.length * r.width)
}

// method
func (r rectangle)area(){
	fmt.Printf("Rectangle result:%d\n", r.length * r.width)
}

func perimeter(p *rectangle){
	fmt.Printf("Perimeter result:%d\n", 2*(p.length+p.width))
}

func (p *rectangle)perimeter(){
	fmt.Printf("Perimeter result:%d\n", 2*(p.length+p.width))
}

func main(){
	r := rectangle{
		length: 4,
		width: 3,
	}
	// calling function
	r.area()

	// calling method
	area(r)

	p := &r

	// it will throw an error as go does not allow to call function using pointer,
	// functions only accept value receiver
	//area(p)

	// it is allowed to use value receiver and pointer receiver both with method
	p.area()

	// if method has pointer receiver , then method argument must be pointer not value
	perimeter(p)

	// this is not allowed
	//perimeter(r)

	// same as method ,it allows to call pointer receiver with pointer and value
	p.perimeter()

	// it is same as (&r).perimeter()
	r.perimeter()

}
