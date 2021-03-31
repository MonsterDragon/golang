package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1, 2, 3, 4 ,5)
	fmt.Printf("%p\n", s)
	s = append(s, 1, 2, 3, 4 ,5)
	fmt.Printf("%p\n", s)
	s2 := make([]int, 10)
	copy(s2, s)
	fmt.Println(s2)
	s3 := make([]int, 10)
	copy(s3, s[2:5])
	fmt.Println(s3)
}
