package main

import "fmt"

func main() {
	var arr[10] int = [10] int{1, 2, 3, 4, 5, 6}
	var max int = 0
	for _, m := range arr{
		if m > max{
			max = m
		}
	}
	fmt.Println(max)
}
