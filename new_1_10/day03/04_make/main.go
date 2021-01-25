package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 5)
	fmt.Printf("s1=%v len(s2)=%d cap(s2)=%d\n", s1, len(s2), cap(s2))

	s3 := make([]int, 0, 10)
	fmt.Printf("s3=%v len(s3)=%d cap(s3)=%d\n", s3, len(s3), cap(s3))
	// 切片的容量是第一个切片的首元素开始到最后

	s4 := [...]int{1, 2, 3}
	s5 := s4
	fmt.Println(s4, s5)
	s4[0] = 1000
	fmt.Println(s4, s5)
	s6 := s4[1:]

	// 切片的遍历
	for i := 0; i < len(s6); i++{
			fmt.Println(s6[i])
	}
}
