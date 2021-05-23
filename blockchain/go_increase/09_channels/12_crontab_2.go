package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.After(time.Second * 2 )
	fmt.Println("now time : ", <- nowTime)
}
