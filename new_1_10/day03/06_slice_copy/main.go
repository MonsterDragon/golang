package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s1[2] = 4
	s2 := s1
	s3 := make([]int, 3, 3)
	copy(s3, s1)
	fmt.Println(s2, s3)
	// 将数组中2删除
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(cap(s1))

	a1 := []int{1, 3, 5}
	fmt.Printf("%p\n", a1)
	a2 := a1[:]
	fmt.Printf("%p\n", a2)
	fmt.Println(a2, cap(a2), len(a2))
	a2 = append(a2[:1], a2[2:]...)
	fmt.Println(a2, cap(a2), len(a2))
	fmt.Println(a1)
	fmt.Printf("%p, %p\n", a1, a2)
}
