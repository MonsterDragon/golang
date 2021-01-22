package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 切片定义
	var s1 []int // 定义一个存放int类型的元素切片
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	s1 = []int{1, 2, 3}
	s2 = []string{"hh", "aa"}
	fmt.Println(s1, s2)
	fmt.Println(reflect.TypeOf(s1), reflect.TypeOf(s2))

	var c1 [3]int
	var c2 [4]string
	fmt.Println(reflect.TypeOf(c1), reflect.TypeOf(c2))

	// len()函数求长度 cap()函数求切片的容量
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))

	d1 := [...]int{1, 2, 3, 4}
	s3 := d1[1:3] // 左包含右不包含
	fmt.Println(s3)
	s4 := d1[1:] // => [1:len(d1)]
	fmt.Println(s4)
	s5 := d1[:4] // => [0:4]
	fmt.Println(s5)
	s6 := d1[:] // => [0:len(d1)]
	fmt.Println(s6)
	// 切片的容量是指从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s3)=%d, cap(s3)=%d\n", len(s3), cap(s3))
	fmt.Printf("len(s4)=%d, cap(s4)=%d\n", len(s4), cap(s4))
	fmt.Printf("len(s5)=%d, cap(s5)=%d\n", len(s5), cap(s5))

	// 切片再切片

}
