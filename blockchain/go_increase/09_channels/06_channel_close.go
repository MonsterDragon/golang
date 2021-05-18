package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan int) // 无缓冲channel
	//ch := make(chan string)

	go func() {
		for i:=0; i<8; i++{
			ch <- i
		}
		close(ch)
	}()

	for  {
		if num, ok := <- ch; ok == true {
			fmt.Println("主go程读", num)
		} else {
			fmt.Println("channel 关闭")
			n := <- ch
			fmt.Println(n) // 无缓冲，或有缓冲channel，在channel关闭后仍能从中读到0（channel 声明时是int类型，若为非int则不能继续读取），但是不能写入channel
			os.Exit(0)
		}
	}
}
