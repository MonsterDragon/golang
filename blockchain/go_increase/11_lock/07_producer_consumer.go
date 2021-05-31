package main

import (
	"fmt"
	"os"
	"time"
)

func producer(ch chan <- int) {
	for i:=0; i<10; i++ {
		ch <- i
		fmt.Println("producer write ", i)
	}
	close(ch)
}

func consumer(ch <- chan int)  {
	for  {
		if info, ok := <- ch; ok {
			fmt.Println("consumer read ", info)
		} else if ok == false {
			time.Sleep(time.Second * 1)
			fmt.Println("channel close")
			os.Exit(1)
		}
	}
}

func main()  {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)
	go consumer(ch)
	go consumer(ch)

	for  {
		;
	}
}
