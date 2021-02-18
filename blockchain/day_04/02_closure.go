package main

import "fmt"

func test2() func() int {
	var a int

	return func() int{
		a++
		return a
	}
}

func main() {
	f := test2()
	fmt.Printf("%T\n", f)

	for i:=0;i<10;i++{
		fmt.Printf("%d\n", f())
	}
}
