package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s is moving\n", a.name)
}

type Dog struct {
	Feet int
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s is yelling\n", d.name)
}

func main()  {
	d1 := Dog{4, &Animal{name: "huanhuan"}}
	d1.wang()
	d1.move()
}
