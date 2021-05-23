package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Now: ", time.Now())
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <- myTimer.C
	fmt.Println("当前时间: ", nowTime)
}
