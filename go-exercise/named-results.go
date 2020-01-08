package main

import "fmt"

func split(sum int) (x, y int) {
    fmt.Println(sum)

    x = sum * 4 / 9
    fmt.Pinntln(x)
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
