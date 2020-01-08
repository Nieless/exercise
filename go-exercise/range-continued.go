package main

import "fmt"

func main(){
	pow := make([]int, 10)

	fmt.Println(pow, len(pow), cap(pow))

	for i:= range pow {
		pow[i] = 1 << uint(i)
	}

	fmt.Println(pow)

	for i, value := range pow {
		fmt.Printf("%d=%d\n", i, value)
	}
}
