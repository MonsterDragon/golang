package main

import "fmt"

type Flyabler interface {
	fly()
}

type Bird struct {

}

func (b *Bird) EatAndDrink() {
	fmt.Println("bird eatanddrink")
}

type MaQue struct {
	Bird
}

func (m *MaQue) fly() {
	fmt.Println("mauqe is flying")
}

type Yingwu struct {
	Bird
}

func (y *Yingwu) fly() {
	fmt.Println("yingwu is flying")
}

type Planes struct {

}

func (p *Planes) fly() {
	fmt.Println("planes is flying")
}

func WhoFly(i Flyabler)  {
	i.fly()
}

func main() {
	m := &MaQue{}
	m.EatAndDrink()
	WhoFly(m)
	plane := &Planes{}
	WhoFly(plane)
}
