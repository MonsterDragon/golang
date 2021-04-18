package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Println("lambda function a ", a)
		fmt.Println("lambda function b ", b)
	}(a, b)
	a = 100
	b = 200
	fmt.Println("main function a ", a)
	fmt.Println("main function b ", b)
}
