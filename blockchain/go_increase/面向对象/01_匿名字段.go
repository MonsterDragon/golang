package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s1 := Student{Person{"shuzhan", "male", 22}, 1, "sz"}
	fmt.Println(s1)

	s2 := Student{Person: Person{"shuzhan", "male", 22}}
	fmt.Println(s2)

	s3 := Student{Person: Person{name: "shuzhan"}}
	fmt.Println(s3)

	currentTime := time.Now()
	fmt.Println(currentTime)
}
