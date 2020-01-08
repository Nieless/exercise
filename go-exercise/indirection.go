package main

import "fmt"

type Vertex struct {
	x, y float64
}


// function
func (v *Vertex) Scale(f float64){
	v.x = v.x + f
	v.y = v.y + f
}

// method
func ScaleFunc(v *Vertex, f float64){
	v.x = v.x + f
	v.y = v.y + f
}

func main(){
	v := Vertex{3, 4}
	v.Scale(2)
	fmt.Println(v)

	// pointer argument must take a pointer
	ScaleFunc(&v, 2)

	
	p := &Vertex{4, 3}
	p.Scale(2)
	ScaleFunc(p, 8)
	fmt.Println(v)
}
