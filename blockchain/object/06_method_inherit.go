package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (a *Person) PrintInfo() {
	fmt.Printf("name=%s, sec=%c, age=%d\n", a.name, a.sex, a.age)
}

type Student_2 struct {
	Person
	id int
	addr string
}

func main() {
	s := Student_2{Person{"shuzhan", 'm', 18}, 100, "harbin"}
	s.PrintInfo()
}
