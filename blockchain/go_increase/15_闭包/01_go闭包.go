package main

import "fmt"

// func a(x int) func(y int) int {
// 	b := func(y int) int {
// 		z := x + y
// 		fmt.Println(z)
// 		return z
// 	}
// 	return b
// }

// func main() {
// 	c := a(1)

// 	fmt.Printf("%#v\n", c)
// 	c(2)
// 	c(3)
// 	c(4)
// }

// func test() func() {
// 	x := 100
// 	fmt.Printf("x (%p) = %d\n", &x, x)

// 	return func() {
// 		fmt.Printf("x (%p) = %d\n", &x, x)
// 	}
// }

// func main() {
// 	f := test()
// 	f()
// }

// func add(base int) func(int) int {
// 	return func(i int) int {
// 		base += i
// 		return base
// 	}
// }

// func main() {
// 	tmp := add(10)
// 	fmt.Println(tmp(1), tmp(2)) // 11 13
// }

// 返回两个闭包

// func test01(base int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		base += i
// 		return base
// 	}

// 	sub := func(i int) int {
// 		base -= i
// 		return base
// 	}

// 	return add, sub
// }

// func main() {
// 	f1, f2 := test01(10)
// 	fmt.Println(f1(1), f2(2))
// }
