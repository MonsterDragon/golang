package main

import "fmt"

func main() {
	// new
	a := new(int) // *int
	b := new(bool) // *bool
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)

	var c *int
	c = new(int)
	*c = 100
	fmt.Println(c)

	// make
	s1 := make([]int, 5)
	fmt.Printf("%p\n", s1)
	s2 := make(map[string]int, 10)
	fmt.Printf("%p\n", &s2)

	//test
	var d int
	fmt.Printf("%p\n", &d)
	var e *int
	e = &d
	*e = 10
	fmt.Println(d)
}
