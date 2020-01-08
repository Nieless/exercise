package main

import "fmt"

type vertex struct {
	lat, long float64
}

var m = map[string]vertex{
	"a":{40.5, -80.11},
	"b":{80.60,11.90},
}

func main(){
	fmt.Println(m)
	fmt.Println(m["a"])

	d := make(map[string]int)
	d["a"] = 42
	d["a"] = 40
	fmt.Println(d)

	delete(d, "a")
	fmt.Println(d)

	v, k := d["a"]
	fmt.Println(v,k)
}
