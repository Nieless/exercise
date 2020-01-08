package main

import "fmt"

type IpAddr [4]byte

func main(){
	hosts := map[string]IpAddr{
		"local": {127, 0, 0, 1},
		"DNS": {8, 8, 8, 8},
	}
	fmt.Println(hosts)
}

func (ip IpAddr) String() string{
	for name, ip := range hosts{
                return fmt.Printf("%v: %v \n", name, ip)
        }
}
