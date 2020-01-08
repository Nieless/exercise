package main

import (
	"fmt"
)

func main(){
	x, y := 4, 5
	var f float32 = float32(x+y)
        fmt.Println(f)
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
