package main

import "fmt"

var (
	a = 5
	b = 2
)

func main() {
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// 位运算:针对二进制
	// &按位与
	fmt.Println(5 & 2) // 都为1才是1
	// |按位或（两位有一个为1就是1）
	fmt.Println(5 | 2)
	// ^按位异或（两位不一样为1）
	fmt.Println(5 ^ 2)
	fmt.Println(5 << 1)
	fmt.Println(5 << 2)
	fmt.Println(5 >> 2)
	fmt.Println(5 >> 1) // 向右移位除2
}
