package main

import (
	"fmt"
)

type student struct {
	id int
	age int
	name string
	score float32
}

// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
type zombie struct {
	string
	int
}

func main() {
	var s1 student
	s1.score = 10.002
	s1.id = 1
	s1.age = 20
	s1.name = "shuzhan"
	fmt.Printf("p1=%v\n", s1)
	fmt.Printf("p1=%#v\n", s1)

	var user struct{Name string; Age int} // 匿名结构体
	user.Name = "shuzhan"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	var u1 = new(student) // 实例化结构体 创建指针类型的结构体
	u1.score = 10.002
	u1.id = 1
	u1.age = 20
	u1.name = "shuzhan"
	fmt.Printf("%T\n", u1) // 值的类型
	fmt.Printf("p2=%#v\n", u1)

	s2 := &student{} // 取结构体的地址实例化 创建指针类型的结构体
	fmt.Printf("%T\n", s2)     //*main.person
	fmt.Printf("p3=%#v\n", s2) //p3=&main.person{name:"", city:"", age:0}
	s2.score = 10.002
	s2.id = 1
	s2.age = 20
	s2.name = "shuzhan"

	s3 := student{		// 使用键值对对结构体进行初始化 非指针类型初始化
		score : 10.002,
		id : 1,
		age : 20,
		name : "shuzhan",
	}
	fmt.Printf("%#v\n", s3)

	s4 := &student{
		score : 10.002,
		id : 1,
		age : 20,
		name : "shuzhan",
	}
	fmt.Printf("%#v\n", s4) // 对结构体指针进行键值对初始化

	s5 := &student{			// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
		score : 10.002,
		id : 1,
		age : 20,
	}
	fmt.Printf("%#v\n", s5)

	/*
	 1.必须初始化结构体的所有字段。
	 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	 3.该方式不能和键值初始化方式混用。
	*/
	s6 := &student{
		2,
		18,
		"shuzhan",
		10.002,
	}
	fmt.Printf("%#v\n", s6)

	z1 := zombie{"mickeal", 10}
	fmt.Printf("%#v\n", z1)
	fmt.Println(z1.string, z1.int)
}
