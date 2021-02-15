package main

import "fmt"

func main01() {
	var a int = 10
	fmt.Printf("%d\n", a)
	var b float64 = 10
	fmt.Printf("%f\n", b)
	var c bool = true
	fmt.Printf("%t\n", c)
	var d byte = 'a'
	fmt.Printf("%c\n", d)
	var e string = "hello"
	fmt.Printf("%s\n", e)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	d = 's'
	fmt.Printf("%T\n", d)
}

func main() {
	a := 123 // 十进制
	b := 0123 // 八进制
	c := 0x123 // 十六进制
	fmt.Println(a, b, c)
	// go语言中不能直接识别二进制
	fmt.Printf("%b\n", a) // 二进制
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)
	fmt.Printf("%o\n", a) // 八进制
	fmt.Printf("%o\n", b)
	fmt.Printf("%o\n", c)
	fmt.Printf("%x\n", a) // 十六进制
	fmt.Printf("%x\n", b)
	fmt.Printf("%x\n", c)
}
