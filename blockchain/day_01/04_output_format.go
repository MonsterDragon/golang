package main

import "fmt"

func main() {
	fmt.Println() // 打印时自带换行
	fmt.Print() // 没有换行操作
	a := 10
	b := 3.1415

	fmt.Printf("%3d\n", a)
	fmt.Printf("%f", b)
	fmt.Printf("%05d\n", a) // 其余位置用0补齐
	fmt.Printf("%.3f\n", b) // %f 默认保留6位小数 %.3f保留3位小数
	fmt.Printf("%p\n", &a)

	var c bool
	fmt.Printf("%t\n", c)
	var d string = "hello"
	fmt.Printf("%s\n", d)
	var e byte = 'a'

	fmt.Printf("%c\n", e)
	fmt.Println(e)
}
