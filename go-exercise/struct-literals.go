package main

import "fmt"

type person struct {
    x, y int
}

var (
    p1 = person{1,2}
    p2 = person{}
    p3 = person{x:8,y:10}
    p4 = &person{5,4}
    p5 = &p4
)

func main() {
    fmt.Println(p1,p2,p3,p4,p5)
}
