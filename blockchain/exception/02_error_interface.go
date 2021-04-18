package main

import (
	"errors"
	"fmt"
)

func Test(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("can not be 0")
	} else {
		result = a / b
	}

	return
}

func main() {
	result, err := Test(3, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println(result)
	}
}
