package main

import "fmt"

func main() {
	//var s [] int
	//fmt.Println(s)
	//自动推导类型创建切片 make([]数据类型，长度)
	s := make([]int, 5)
	s[0] = 123
	s[1] = 234
	s[2] = 213
	fmt.Println(s)

	s = append(s, 465)
	fmt.Println(s)

	var a[] int = []int{1, 2, 3, 4, 5}
	fmt.Println(len(a), cap(a))
	a = append(a, 1, 2, 3)
	fmt.Println(len(a), cap(a))
}