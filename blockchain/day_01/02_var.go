package main

import "fmt"

func main() {
	var a int= 10
	fmt.Println(a)

	var PI float64 = 3.1415926
	var r float64 = 2.5
	var s, l float64

	s = PI * r * r
	l = PI * r * 2
	fmt.Println(s, l)

	w := 2.0
	var p float64 = 5.0
	// 不同类型不能进行计算 需要类型转换
	fmt.Println(w * p)
}
