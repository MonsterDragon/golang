package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func printer(s string)  {
	for _, ch := range s {
		fmt.Printf("%c\n", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

func Person1() {
	//ch <- 1 // 若channel未被读出则阻塞，不会打印hello
	printer("hello")
	ch <- 1
}

func Person2() {
	out := <- ch
	fmt.Println("channel's injury is ", out)
	printer("world") // 这个可以读出
}

func main() {
	go Person1()
	go Person2()
	for  {
		;
	}
}
