package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int) // 用来进行数据通信的 channel
	quit := make(chan bool) // 用来判断是否退出的 channel
	
	go func() {
		for i:=0; i<5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true // 通知主go程退出
		runtime.Goexit()
	}()
	
	for {
		select {
		case num := <- ch:
			fmt.Println("read ", num)
		case <- quit:
			return
		}
		fmt.Println("+++++++++++++")
	}
}
