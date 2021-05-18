package main

import (
	"fmt"
	"time"
)

func Count(ch chan int)  {
	fmt.Println("Counting")
	ch <- 1
}

func main() {
	chs := make([]chan int, 10) // go语言中声明类型是空指针，为nil，不能直接使用
	for i:=0; i<10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i]) // 并发执行count
	}
	for _, ch := range chs {
		value := <-ch
		fmt.Println(value)
	}
	time.Sleep(100 * time.Millisecond)
}