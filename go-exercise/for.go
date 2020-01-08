package main

import "fmt"

func main() {
	sum := 0
	for i := 0; sum < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
