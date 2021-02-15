package main

import "fmt"

func main() {
	a := 10.0
	b := 20.0
	c := a/b
	fmt.Println(c)
	// 取余运算符只能用于整数
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
}