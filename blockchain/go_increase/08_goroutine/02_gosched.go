package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		for {
			fmt.Println("this is goroutine test")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		runtime.Gosched() // 出让时间片
		fmt.Println("this is main test")
	}
}
