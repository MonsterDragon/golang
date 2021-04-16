package main

import (
	"fmt"
)

type Human interface {
	sayhi()
}

type Student struct {
	name string
	id int
}

func (tmp *Student) sayhi() {
		fmt.Printf("%s, %d, sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr string
	group string
}

func (tmp *Teacher) sayhi() {
	fmt.Printf("%s, %s, sayhi\n", tmp.addr, tmp.group)
}

func main() {
	var i Human
	s := &Student{"shuzhan", 100}
	i = s
	i.sayhi()
	ha := &Teacher{"harbin", "3glass"}
	i = ha
	i.sayhi()
}
