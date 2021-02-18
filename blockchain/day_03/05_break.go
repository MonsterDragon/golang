package main

import "fmt"

func main() {
	// break continue goto
	for i:=0; i<=100; i++ {
		if i % 2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}
