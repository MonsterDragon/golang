package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30
	d := 40
	fmt.Println(a<b && c<d)

	fmt.Println(a<b || c>d)

	fmt.Printf("%p\n", &a)
}
