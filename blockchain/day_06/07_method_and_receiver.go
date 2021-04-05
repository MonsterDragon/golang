package main

import "fmt"

type Person struct {
	name string
	city string
	age int
}

func NewPerson(name, city string, age int) *Person{
	return &Person{
		name: name,
		city: city,
		age: age,
	}
}

func (p Person) Dream() {
	fmt.Printf("%s wants to learn golang\n", p.name)
} // 实例接收者

func (p *Person) SetAge(newAge int) {
	//fmt.Printf("%p\n", p)
	p.age = newAge
} // 指针类型接收者

func (p Person) SetAge2(newAge int) {
	fmt.Printf("%p\n", &p)
	p.age = newAge
} // 值类型的接收者

func main()  {
	p1 := NewPerson("shuzhan", "harbin", 18)
	p1.Dream()
	p2 := NewPerson("shuchang", "harbin", 22)
	p2.SetAge(31)
	fmt.Println(p2.age)
	p3 := NewPerson("shuchang", "harbin", 22)
	fmt.Println(p3.age)
	fmt.Printf("%p\n", &p3)
	p3.SetAge2(31)
	fmt.Printf("%p\n", &p3)
	fmt.Println(p3.age)
}
