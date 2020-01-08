package main

import "fmt"

func main() {
    primes := [6]int{2,6,9,4,3}

    var s []int = primes[1:4]
    fmt.Println(s)
}
