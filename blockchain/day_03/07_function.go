package main

import "fmt"

func add(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {
	a := 10
	b := 20
	add(a, b)
}
