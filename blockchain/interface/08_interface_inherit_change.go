package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Perspner interface {
	Humaner
	sing(lrc string)
}

type Student_2 struct {
	name string
	id int
}

func (tmp *Student_2) sayhi() {
	fmt.Printf("%s, %d\n", tmp.name, tmp.id)
}

func (tmp *Student_2) sing(lrc string) {
	fmt.Printf("%s singing %s\n", tmp.name, lrc)
}

func main() {
	var i Perspner
	s := Student_2{"shuzhan", 100}
	i = &s
	i.sayhi()
	i.sing("dangnianqing")
}
