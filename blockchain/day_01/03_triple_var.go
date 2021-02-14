package main

import "fmt"

func main() {
	a, b, c := 10, 20, 30
	fmt.Println(a, b, c)
	a, b, c = c, a, b
	fmt.Println(a, b, c)
}
