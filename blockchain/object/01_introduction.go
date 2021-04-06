package main

import "fmt"

func main() {
	type Person struct {
		id int
		name string
		age int
	}

	type Student struct {
		*Person
		score float64
	}

	// 严格意义上说，GO语言中没有类(class)的概念,但是我们可以将结构体比作为类，因为在结构体中可以添加属性（成员），方法（函数）。
	var s1 Student = Student{&Person{10, "shuzhan", 18}, 10}
	fmt.Println("s1 = ", s1)
	fmt.Printf("s1 = %+v\n", s1) // s1 = {Person:{id:10 name:shuzhan age:18} score:10}
	s2 := Student{score: 97}
	fmt.Printf("%+v\n", s2) // {Person:{id:0 name: age:0} score:97}
	s1.Person.id = 100
	fmt.Printf("s1 = %+v\n", s1)
	s1.Person = &Person{100, "shuchang", 18}
	fmt.Printf("s1 = %+v\n", s1)
	fmt.Println(s1.id, s1.name, s1.age)
	fmt.Println(s1.Person.name)
	// 继承的结构体如果有重复字段

	var s Student
	s.Person = new(Person)
	s.id = 10
	s.name = "shuzhan"
	s.age = 100
	s.score = 91
	fmt.Println(s.id, s.score)
}
