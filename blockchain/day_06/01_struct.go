package main

import "fmt"

func main() {
	// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）
	type student struct {
		id int
		name string
		sex byte
		age int
		addr string
	}
	var s1 student = student{1, "mike", 'a', 18, "hrb"} // 顺序必须保持一致
	fmt.Println(s1)
	var s student
	s.age = 11
	s.name = "John"
	s.id = 2
	s.sex = 'w'
	s.addr = "shanghai"
	fmt.Println(s)
	var tmp student
	tmp = s1
	fmt.Println(tmp)
	fmt.Printf("%v", tmp)
}
