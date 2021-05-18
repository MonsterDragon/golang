package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i:=0; i<5; i++ {
			fmt.Println("子go程写")
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i:=0; i<5; i++ {
		integar := <- ch
		fmt.Println("主go程读")
		fmt.Println(integar)
	}
}
