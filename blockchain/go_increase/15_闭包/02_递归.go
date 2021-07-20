package main

import "fmt"

// func factorial(i int) int {
// 	if i <= 1 {
// 		return 1
// 	}
// 	return i * factorial(i-1)
// }

// func main() {
// 	fmt.Println(factorial(10))
// }

func fabonacci(i int) int {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return fabonacci(i-1) + fabonacci(i-2)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(fabonacci(i))
	}
}
