package main

import "fmt"

func main() {
	number := 589

	// channels initialized
	sqrch := make(chan int)
	cubech := make(chan int)

	// goroutines started with channel
	go calSquares(number, sqrch)
	go calCubes(number, cubech)

	// read data
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("final output", squares+cubes)
}

func calSquares(number int, sqrch chan int){
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
        	number /= 10
	    }
	//write data to channel
	sqrch <- sum
}

func calCubes(number int, cubech chan int){
	sum := 0
    	for number != 0 {
        	digit := number % 10
        	sum += digit * digit * digit
        	number /= 10
    	}
	//write data to channel
	cubech <- sum
}
