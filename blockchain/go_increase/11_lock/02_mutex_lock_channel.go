package main

import (
	"fmt"
	"time"
)

var cha = make(chan int)

func printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
}

func p1()  {
	printer("hello")
	cha <- 10
}

func p2() {
	<- cha
	printer("world")
}

func main() {
	go p1()
	go p2()
	for  {
		;
	}
}
