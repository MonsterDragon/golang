package main

import "fmt"

func main() {
	var arr [10]int = [10]int{1, 2, 3, 4}

	arr[0] = 123
	arr[1] = 10

	fmt.Println(arr[0], arr[1], arr[2], arr[3])

	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	for m, n := range arr{
		fmt.Println(m, n)
	}

	arr_1 := [10]int{1, 2, 4, 10}
	for _, v := range arr_1{
		fmt.Println(v)
	}

	arr_2 := [...]int{1, 2, 4, 10}
	for _, v := range arr_2{
		fmt.Println(v)
	}
}