package main

import "fmt"

func main() {
	// append(切片对象， 待追加元素)
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 5, 8)
	s2 = append(s1, 34)
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s2)
}
