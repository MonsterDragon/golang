package main

import "fmt"

func sub(a int, b int) int {
	sum := a - b
	return sum
}

func sum(a int, b int) (sum int){
	sum = a + b
	return
}

func test_5()(a int, b int, c int){
	//a = 1
	//b = 2
	//c = 2
	a, b, c = 1, 2, 3
	return
}

func main() {
	value := sub(10, 20)
	fmt.Println(value)
	value = sum(10, 10)
	fmt.Println(value)
	a, b, c := test_5()
	fmt.Println(a, b, c)
	f := test_5
	k, l, m := f()
	fmt.Println(k, l, m)
}
