package main

import "fmt"

type Student struct {
	id int
	name string
	score float64
}

func main() {
	var p *Student = &Student{1, "shuzhan", 90}
	fmt.Println(*p)
}
