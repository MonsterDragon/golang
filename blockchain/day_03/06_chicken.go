package main

import "fmt"

func main() {
	var count int = 0
	for a:=0;a<=20;a++{
		for b:=0;b<=33;b++{
			if a*5+b*3+(100-a-b)/3==100 && (100-a-b)%3==0 {
				fmt.Println(a, b, (100-a-b))
			}
			count++
		}
	}
	fmt.Println(count)
}
