package main

import "fmt"

func test1(a int, b int)  {
	sum := a + b
	fmt.Println(sum)
}

func test2(a int, b int)  {
	test1(a, b)
}

func test3(a ...int){
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}
}

func test4(a ...int){
	test3(a[0:]...)
}

func main()  {
	a := 10
	b := 10
	test2(a, b)
	test4(1, 2, 3, 4)
}