package main

import "fmt"

type car struct {
	kind string
	money float64
}

func (p *car) Move() {
	fmt.Println(p.kind + "car moving")
}

func (p *car) Stop() {
	fmt.Println(p.kind + "car stopping")
}

type BMW struct {
	car
}

type AuDi struct {
	car
}

type CarTyper interface {
	GetCar()
}

func (p *BMW) GetCar() {
	if p.money >= 60000 {
		p.Move()
		p.Stop()
	} else {
		fmt.Println("not enough money!")
	}
}

func (p *AuDi) GetCar() {
	if p.money >= 70000 {
		p.Move()
		p.Stop()
	} else {
		fmt.Println("not enough money!")
	}
}

type CarStore struct {

}

func CarWho(i CarTyper) {
	i.GetCar()
}

func (c *CarStore) Order(money float64, st string) {
	switch st {
	case "BMW":
		CarWho(&BMW{car{st, money}})
	case "AuDi":
		CarWho(&AuDi{car{st, money}})
	}
}

func main() {
	var carstore CarStore
	carstore.Order(60000, "BMW")
}