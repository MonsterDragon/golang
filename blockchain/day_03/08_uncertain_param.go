package main

import "fmt"

func test(args ...int) {
	fmt.Println(args)
	for i:=0;i<len(args);i++{
		fmt.Println(args[i])
	}
	for _, i:= range args{
		fmt.Println(i)
	}
}

func main() {
	test(1,2)
	test(1,2,3,4)
}
