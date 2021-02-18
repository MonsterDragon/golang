package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	func (a int, b int){
		fmt.Println(a + b)
	} (a, b)

	f := func (a int, b int) int{
		return a+b
	}
	fmt.Printf("%T\n", f)
	v := f(1, 2)
	fmt.Println(v)
}
