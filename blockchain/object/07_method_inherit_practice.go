package main

import "fmt"

type Reporter struct {
	Person
	hobby string
}

type Programmer struct {
	Person
	WorkYear int
}

func (a *Reporter) ReporterSayHello(h string) {
	a.hobby = h
	fmt.Printf("%s, %s, %c, %d\n", a.name, a.hobby, a.sex, a.age)
}

func (a *Programmer) ProgrammerSayHello(work int) {
	a.WorkYear = work
	fmt.Printf("%s, %c, ")
}

func main()  {

}
