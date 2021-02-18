package main

import "fmt"

func main() {
	a := 10
	{
		var a int = 20
		fmt.Println(a)
	}
	fmt.Println(a)

	i := 0
	for ;i<=10;i++{

	}
	fmt.Println(i)
}