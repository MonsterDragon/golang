package main

import "fmt"

type Student struct {
	id int
	name string
	age int
}

// 结构体数组
func main() {
	students := []Student{
		{
			1,
			"shuzhan",
			18,
		},
		{
			2,
			"shuchang",
			22,
		},
	}
	fmt.Printf("%#v\n", students)

	age := students[0].age

	for i := 0; i < len(students); i++ {
		if students[i].age > age {
			age = students[i].age
		}
	}
	fmt.Printf("%d\n", age)
}
