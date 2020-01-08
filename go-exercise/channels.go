package main

import (
	"fmt"
	//"time"
	)


func main(){
	n := 3

	// create channel
	out := make(chan int)

	go multiply2(n, out)
	//time.Sleep(time.Second)

	// Once any output is received on this channel, print it to the console and proceed
	fmt.Println(<-out)
}

// This function now accepts a channel as its second argument...
func multiply2(n int, out chan<- int) {
	result := n * 2
	//fmt.Println(result)
	//return result
	// Send result to channel.
	out <- result
}
