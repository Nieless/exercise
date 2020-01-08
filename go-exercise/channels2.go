package main

import (
	"fmt"
	//"time"
	)

func main(){
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	// read channel data
	<-done

}

func hello(done chan bool){
	fmt.Println("hello go routine is going to sleep")
	fmt.Println("hello go routine awake and going to write to done")
	// write to channnel 
	done <- true
}
