package main

import "fmt"

type test struct {
	a int
	b int
	c int
	d int
}

func main() {
	n := test{
		1,
		2,
		3,
		4,
	}
	fmt.Printf("n %p\n", &n)
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.b %p\n", &n.c)
	fmt.Printf("n.b %p\n", &n.d)
}
