package main

import (
	"fmt"
)

type student_s struct {
	name string
	age int
}

func main() {
	m := make(map[string]*student_s)
	stus := []student_s{
		{
			name: "shuzhan",
			age: 18,
		},
		{
			name: "haha",
			age: 22,
		},
		{
			name: "mm",
			age: 23,
		},
	}
 	fmt.Printf("%p\n", &stus)
	fmt.Printf("%p\n", &stus[0])
	fmt.Println(stus[0])
	fmt.Println(stus[1])
	fmt.Println(stus[2])
	fmt.Printf("%p\n", &stus[1])
	fmt.Printf("%p\n", &stus[2])
	for index, stu := range stus{
		fmt.Println(index, stu)
		fmt.Printf("%p\n", &stu)
		m[stu.name] = &stu
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	ma_list := []int{1, 2, 3}
	for _, i := range ma_list{
		fmt.Println(i)
		fmt.Println(&i)
	}
}
