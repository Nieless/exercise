package main

import "fmt"

type Describer interface{
	Describe()
}

func main(){
	var d1 Describer
	if d1 == nil {
		fmt.Printf("type of interface initializer %T and value %v\n",d1, d1)
	}
}
