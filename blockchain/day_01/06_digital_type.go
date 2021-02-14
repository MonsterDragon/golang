package main

import "fmt"

func main() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	b := true
	fmt.Println(b)

	var c float32
	var d float64
	c = 3.1415926467512534345456768
	d = 3.1415926467512534345456768
	fmt.Println(c, d)

	var e byte
	var f byte
	e = 'a'
	f = 'b'
	fmt.Printf("%c, %c\n", e, f)

	var g string = "first"
	var h string = "second"
	fmt.Println(g+h)
}
