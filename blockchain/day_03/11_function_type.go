package main

import "fmt"

func test6() {
	fmt.Println("aa")
}

func test7(a int, b int)  {
	fmt.Println(a + b)
}

type FUNCTYPE func()
type FUNCTEST func(int, int)

func main() {
	var f FUNCTYPE
	f = test6
	f()
	var g FUNCTEST
	g = test7
	g(10, 20)
	var h FUNCTEST
	h = test7
	h(10, 20)
	fmt.Printf("%p, %p\n", g, h)
}
