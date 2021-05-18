package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("stop test")
	runtime.Goexit()
	fmt.Println("start test")
}

func main() {
	go func() {
		defer fmt.Println("goroutine exit")
		test()
		defer fmt.Println("goroutine start")
	}()

	for  {
		;
	}
}
