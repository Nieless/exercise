package main

import "fmt"

func main(){

	s := []int{5, 5, 3, 2}
	printSlice(s)

	s = s[2:]
	printSlice(s)


}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
