package main

import (
	"fmt"
	"runtime"
)


func main(){
	fmt.Print("runs on")
	os := runtime.GOOS
	switch os {
	case "x":
		fmt.Println("os x")
	//case "linux":
	//	fmt.Println("linux")
	default:
		fmt.Printf("%s.", os)
	}
}
