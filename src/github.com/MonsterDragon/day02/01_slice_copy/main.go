package main

import (
	"fmt"
	"sort"
)

func main() {
	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Printf("%T, %T\n", x1, s1)

	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(len(s1), cap(s1))

	fmt.Println(x1)

	a := make([]string, 5, 10)
	var b = make([]string, 5, 10)
	fmt.Println(a == nil)
	fmt.Println(b)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	for _, m := range a {
		fmt.Println(m)
	}

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])
	fmt.Println(a1)
}
