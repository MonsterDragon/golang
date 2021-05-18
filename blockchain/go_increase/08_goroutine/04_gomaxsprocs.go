package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(24)
	fmt.Println("n = ", n)
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}