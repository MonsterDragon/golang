package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	quit := make(chan bool)
	myTicker := time.NewTicker(time.Second)
	i := 0

	go func() {
		for  {
			nowTimer := <- myTicker.C
			i++
			fmt.Println("nowTime: ", nowTimer)
			if i == 8 {
				quit <- true
			}
		}
	}()
	if <- quit {
		runtime.Goexit()
	}
}
