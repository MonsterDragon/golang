package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		for {
			select {
			case v := <- c:
				fmt.Println(v)
			case <- time.After(time.Second * 5):
				fmt.Println("timeout")
				o <- true
				goto lable
			}
		}
		lable:
			return
	}()

	for i:=0; i<2; i++ {
		c <- i
		time.Sleep(time.Second * 2)
	}
	<-o
}
