package main

import "fmt"

// 常量
// 定义了常量之后不会改变的量

const p1 = 3.1415926

const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// 类似枚举
const (
	a1 = iota //0
	a2 // 1
	a3 // 2
)
const (
	c1 = iota // 0
	c2 = 100
	c3 = iota
	c4
)

func main(){
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a3", a3)

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)
}