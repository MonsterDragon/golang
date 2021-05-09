package main

import "fmt"

type Person_1 struct {
	name string
	sex byte
	age int
}

func main() {
	var man *Person_1 = &Person_1{"shu", 'm', 20}
	// 结构体指针使用“.”来索引成员变量
	man.name = "pp"
	fmt.Println(man)

	p3 := new(Person_1)
	p3.name = "n3"
	p3.age = 11
	p3.sex = 'f'
	fmt.Println(p3)
}
