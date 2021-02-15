package main

import "fmt"

func main() {
	const A int = 10
	fmt.Println(A) // 常量不可以取地址
	var b int = 11
	var c int = 12
	d := A + c + b
	fmt.Println(d)
	e := "aa"
	fmt.Printf("%p\n", &e)
	e += "bb"
	fmt.Printf("%p\n", &e)
}
