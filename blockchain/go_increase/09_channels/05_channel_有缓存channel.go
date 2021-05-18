package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) // 缓存长度为3，缓存channel可以不用写完就会被读走
	fmt.Println("first len(ch):", len(ch), "cap(ch)", cap(ch))
	go func() {
		for i:=0; i<8; i++ {
			fmt.Println("子go程：i=", i)
			ch <- i
			fmt.Println("len(ch):", len(ch), "cap(ch)", cap(ch))
			time.Sleep(time.Second*2)
		}

	}()

	for i:=0; i<8; i++ {

		info := <- ch
		fmt.Println("主go程读：", info)
	}
}
