package main

import "fmt"

func main(){
    names := [4]string{
		"h",
		"e",
		"l",
		"l",
	}

    fmt.Println(names)

    a := names[0:2]
    b := names[1:3]

    fmt.Println(a,b)

    b[0] = "x"
    fmt.Println(a, b)
    fmt.Println(names)

}
