package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30

	fmt.Println(a+b>=c)
	var year int
	fmt.Println("enter year")
	fmt.Scanf("%d", &year)
	//fmt.Println(year)
	fmt.Println(year % 400)
	fmt.Println(year % 400 == 0 || (year % 4 == 0 && year % 100 !=0))
}
