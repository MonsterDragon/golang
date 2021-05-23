package main

import (
	"fmt"
	"time"
)

func main() {
	myTimer := time.NewTimer(time.Second * 3)
	myTimer.Reset(time.Second * 2)
	go func() {
		fmt.Println("now time: ", time.Now())
		timer := <- myTimer.C
		fmt.Println("子go程，定时完毕: ", timer)
	}()
	//myTimer.Stop() // 设置定时器停止，将定时器归零
	for  {
		;
	}
}
