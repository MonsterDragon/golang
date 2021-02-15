package main

import "fmt"

func main() {
	var a byte = 'a'
	fmt.Println(a)
	var b string = "a"
	fmt.Println(b)

	var str string = "hello world"
	num := len(str)
	fmt.Println(num)
}
