package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <- chan int, quit <- chan bool)  {
	for  {
		select {
		case num := <- ch:
			fmt.Print(num, " ")
		case <- quit:
			runtime.Goexit()
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go fibonacci(ch, quit)

	x, y := 1, 1
	for i:=0; i<20; i++ {
		x, y = y, x+y
		ch <- x
	}
	quit <- true
}
