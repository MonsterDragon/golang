package main

import "fmt"

func main() {
	ch := make(chan string) // 无缓存channel
	fmt.Println("len(cp)=", len(ch), "cap(cp)=", cap(ch))

	go func() {
		for i:=0; i<2; i++{
			fmt.Println("i=", i)
		}
		ch <- "子go打印完毕"
	}()
	info := <- ch
	fmt.Println(info)
}