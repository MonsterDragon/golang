package main

import "fmt"

type Person_3 struct {
	name string
	age int
	gender byte
}

func (p *Person_3) PrintInfo() {
	fmt.Printf("%s, %c, %d\n", p.name, p.gender, p.age)
}

type Student_3 struct {
	Person_3
	id int
	addr string
}

func (s *Student_3) PrintInfo() {
	fmt.Println("s = ", s)
}

func main() {
	s := Student_3{Person_3{"shuzhan", 18, 'm'}, 100, "harbin"}
	s.PrintInfo()
}
