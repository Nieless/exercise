package main

import "fmt"

func main(){
	var s []int
	printSlice(s)

	s = append(s, 1)
	printSlice(s)
}

func printSlice(s []int){
	fmt.Printf("length=%d capacity=%d %v", len(s), cap(s), s)
}
