package main

import "fmt"

// 声明变量

var (
	name string
	age int
	isOK bool
)

func foo() (int, string){
	return 10, "Q1mi"
}

func main(){
	name = "lixiang"
	age = 18
	isOK = true
	// Go语言中非全局变量声明必须使用，不适用就编译不过去
	fmt.Print(isOK)
	fmt.Printf("name: %s", name)
	fmt.Println(age)

	var s1 string = "shzuhan"
	fmt.Print(s1)
	s2 := 10 // 只有在函数中才能使用
	fmt.Println(s2)
	x, _ := foo()
	_, y := foo()
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}
