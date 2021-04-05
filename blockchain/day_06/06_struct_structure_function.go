package main

import "fmt"

type person struct {
	name string
	city string
	age int
}

func newPerson(name, city string, age int) (*person) {
	return &person{
		name: name,
		city: city,
		age: age,
	}
}

func main() {
	p1 := newPerson("shuzhan", "harbin", 27)
	fmt.Println(p1)
}
