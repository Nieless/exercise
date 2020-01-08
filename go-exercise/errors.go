package main

import (
	"fmt"
	"errors"
)

func f1(arg int) (int, error){
	if arg < 0{
		return 0, errors.New("cant pass negative arg")
	}
	return arg + 3, nil
}

type argError struct{
	err_code int
	msg string
}

func (e *argError) Error(){
	return fmt.Sprintf("%d - %s", e.err_code, e.msg)
}

func f2(arg int)(int, error){
	if arg < 0{
                return 0, &argError{err_code: 500, msg:"cant pass negative arg"}
        }
        return arg + 3, nil
}

func main(){
	for _, i := range []int{-5,10}{
		if r, e := f2(i); e !=nil{
			fmt.Println("error:", e)
		} else {
			fmt.Println(r, e)
		}
	}
}
