package main

import "fmt"

func test03(a int) {
	if a==0{
		return
	}
	a--
	test03(a)
	fmt.Println(a)
}

var sum int = 1

func test04(a int) {
	if a == 1{
		return
	}
	test04(a - 1)
	sum *= a
}

func main()  {
	test03(10)
	test04(10)
	fmt.Println(sum)
}