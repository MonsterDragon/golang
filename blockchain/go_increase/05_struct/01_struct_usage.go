package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	sex byte
	age int
}

func test(a Person)  {
	a.name = "name1"
	a.age = 10
	fmt.Println(a)
}

func main() {
	// 顺序初始化
	var man Person = Person{"shuzhan", 'm', 27}
	fmt.Println(man)

	// 部分初始化
	man_1 := Person{name: "ss"}
	fmt.Println(man_1)

	// 索引成员变量
	var man_2 Person
	man_2.name = "aa"
	man_2.age = 99
	fmt.Println(man_2)

	// 结构体比较 只能使用 == 和 !=
	var p1 Person = Person{"shuzhan", 'm', 27}
	var p2 Person = Person{"shuzhan", 'm', 271}
	var p3 Person = Person{"shuzhan", 'm', 272}
	fmt.Println(p1 == p2, p1 == p3)
	var tmp Person
	fmt.Println(unsafe.Sizeof(tmp))
	// 结构体传参是值传递
	test(tmp)
	fmt.Println(tmp)
	// 结构体地址和首个元素的地址是一样的
	fmt.Printf("%p\n", &tmp)
	fmt.Println(&tmp.name)
}
