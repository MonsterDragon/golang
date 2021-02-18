package main

import "fmt"

func main() {
	var (
		b int
		c int
		d int
	)
	for i:=100; i<=999; i++{
		b = i / 100
		c = i % 100 / 10
		d = i % 10
		if b*b*b + c*c*c + d*d*d == i{
			fmt.Println(i)
		}
	}
}
