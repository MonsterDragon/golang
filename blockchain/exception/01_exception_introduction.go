package main

import (
	"errors"
	"fmt"
)

/*
	error接口，它是Go语言内建的接口类型，该接口的定义如下：
	type error interface {
	Error() string
	}
*/

func main() {
	err := errors.New("this is normal err")
	fmt.Println("err = ", err.Error())

	err1 := fmt.Errorf("%s\n", "this is normal errl")
	fmt.Println("errl = ", err1)
}