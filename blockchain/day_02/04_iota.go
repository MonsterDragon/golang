package main

import "fmt"

func main_01() {
	const(
		a = iota // 常量自动生成器
		b = iota
		c = iota
		d = iota
	)

	fmt.Println(a, b, c, d)
}

func main02() {
	const(
		a = iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}

func main() {
	const(
		a = 10
		b, c = iota, iota // 1, 1 同一行的iota相同
		d, e // 2, 2
	)

	fmt.Println(a, b, c, d, e)
	const f = iota
	fmt.Println(f) // 1 在一个const声明语句中，在第一个声明的常量所在的行，iota会被重置为0
}
