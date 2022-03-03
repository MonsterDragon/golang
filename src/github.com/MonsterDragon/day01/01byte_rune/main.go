package main

import "fmt"

func main() {
	s1 := "哈哈"
	s2 := []rune(s1)
	s2[0] = '嘿'
	fmt.Println(string(s2))
	c1 := "红"
	c2 := '红'
	fmt.Printf("%T, %T\n", c1, c2)
	c3 := byte(c2)
	fmt.Printf("%T\n", c3)
}
