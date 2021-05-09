package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	//var a int = 100
	//var p *int = &a
	//
	//*p = 250
	//
	//fmt.Println(a, *p)
	var p *int
	p = new(int)
	fmt.Println(*p)

	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b)
}
