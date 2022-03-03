package main

import "fmt"

type Student1 struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func Test(stu Student1) {
	stu.id = 666
	fmt.Printf("%#v\n", stu)
}

func main() {
	s := Student1{1, "Mike", 'm', 11, "hb"}
	Test(s)
	fmt.Printf("%#v\n", s)
}
