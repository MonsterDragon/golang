package main

import "fmt"

type Student struct {
	id int
	name string
	age int

	score float64
}

func (s *Student) Printshow()  {
	s.id = 100
	fmt.Println(*s)
}

func main() {
	s := &Student{101, "zhangsan", 19, 90}
	s.Printshow()
	s1 := Student{100, "shuzhan", 20, 80}
	s1.Printshow()
	/*
	第一：只要接收者类型不一样，这个方法就算同名，也是不同方法，不会出现重复定义函数的错误
	第二：接收者可以是指针
	第三：接收者为普通变量，非指针，值传递 接收者为指针变量，引用传递
	第四：无论我们生成的是指针变量还是结构体变量都可以调用相关的方法
	*/
}
