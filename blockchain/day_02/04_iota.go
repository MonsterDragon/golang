package main

import "fmt"

func main_01() {
	const(
		a = iota
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
		b, c = iota, iota
		d, e
	)

	fmt.Println(a, b, c, d, e)
}
