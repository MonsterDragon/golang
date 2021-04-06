package main

import (
	"fmt"
)

type Student_1 struct {
	name string
	gender string
	age int
	chin float64
	math float64
	english float64
}

func (a *Student_1) SayHello(name string, gender string, age int) {
	a.name = name
	a.gender = gender
	a.age = age
	if a.age < 0 || a.age > 100 {
		a.age = 0
	}
	if a.gender != "male" || a.gender != "female" {
		a.gender = "male"
	}
	fmt.Printf("%s, %s, %d\n", name, gender, age)
}

func (a *Student_1) ShowScore(chin float64, math float64, english float64) {
	a.chin = chin
	a.math = math
	a.english = english
	var sum float64
	sum = a.chin + a.math + a.english
	fmt.Printf("%f, %f\n", sum, sum/3)
}

func main() {
	var s Student_1
	s.SayHello("shuzhan", "male", 22)
}
