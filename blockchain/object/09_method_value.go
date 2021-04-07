package main

import "fmt"

type Person struct {
	name string
	gender byte
	age int
}

func (p Person) SetInfoVlaue() {
	fmt.Printf("%p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("%p, %v\n", &p, p)
}

func main() {
	p := Person{"shuzhan", 'm', 18}
	fmt.Printf("%p, %v\n", &p, p)

	f := (*Person).SetInfoPointer
	f(&p)

	f1 := (Person).SetInfoVlaue
	f1(p)
}
