package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) SayHello() {
	fmt.Println("My name is "+p.name)
}

type Chinese struct {
	Person
}

func (c *Chinese) SayHello() {
	fmt.Println("I am Chinese, My name is "+c.name)
}

type English struct {
	Person
}

func (e *English) SayHello() {
	fmt.Println("I am English, My name is "+e.name)
}

type Personer interface {
	SayHello()
}

func main() {
	c := &Chinese{Person{"shuzhan"}}
	e := &English{Person{"Mike"}}
	x := []Personer{c, e}
	//x := make([]Personer, 2, 4)
	//x[0] = c
	//x[1] = e
	for i:=0; i<len(x); i++ {
		x[i].SayHello()
	}
}
