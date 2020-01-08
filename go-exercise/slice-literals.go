package main

import "fmt"

func main(){
    q := []int{2,5,5,4}
    fmt.Println(q)

    r := []bool{true, true, true, true, false}
    fmt.Println(r)

    s := []struct {
	    i int
	    b bool
	}{
	    {2,true},
	    {3,false},
	    {4,true},
	}

    fmt.Println(s)
}
