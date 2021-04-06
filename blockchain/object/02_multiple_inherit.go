package main

import "fmt"

func main() {
	type Object struct {
		id int
		flag bool
	}

	type Person struct {
		Object
		name string
		age int
	}

	type Student struct {
		Person
		name string
		score float64
	}

	var s1 Student = Student{Person{Object{15, true}, "laozhang", 18}, "zhangsan", 19}
	fmt.Printf("s1 = %+v\n", s1)
}
