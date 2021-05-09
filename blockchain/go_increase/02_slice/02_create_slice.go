package main

import "fmt"

func main() {
	// way1
	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)
	// way2
	s2 := make([]int, 5, 10)
	fmt.Println(len(s2), cap(s2))
	// way3
	s3 := make([]int, 5)
	fmt.Println(len(s3), cap(s3))
	// way4
	var s4 []int = []int{1, 2, 3, 4}
	fmt.Println(s4)
	// make只能创建slice map channel
}
