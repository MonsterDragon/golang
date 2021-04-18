package main

import "fmt"

func double(x int) int {
	return x + x
}

func triple(x int) (r int) {
	defer func() {
		r += x
	}()
	return double(x)
}

func main() {
	fmt.Println(triple(3))
}
