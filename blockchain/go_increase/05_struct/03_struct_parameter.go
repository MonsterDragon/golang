package main

import "fmt"

type Person_2 struct {
	name string
	age int
	flg bool
	interest []string
}

func initFunc(p *Person_2) {
	p.name = "shuzhan"
	p.age = 10
	p.flg = true
	p.interest = []string{"eat", "sports"}
}

func main() {
	person := new(Person_2)
	initFunc(person)
	fmt.Println(person)
}
