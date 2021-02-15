package main

import "fmt"

func main() {
	var score int
	fmt.Scanf("%d", &score)
	if score > 170 {
		fmt.Println("hello")
	}else {
		fmt.Println("world")
	}

	if score > 170 {
		fmt.Println("a")
	} else if score > 100 {
		fmt.Println("b")
	} else if score > 50 {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}
}