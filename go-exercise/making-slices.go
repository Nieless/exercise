package main

import "fmt"

func main(){
	a := make([]int, 5)

	fmt.Println(len(a), cap(a), a)

	b := make([]int, 0, 5)

	fmt.Println(len(b), cap(b), b)

	c := b[:2]

	fmt.Println(len(c), cap(c), c)

	d := c[:2]

	fmt.Println(len(d), cap(d), d)
}
