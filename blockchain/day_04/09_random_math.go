package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var start time.Time = time.Now()
	start := time.Now().Nanosecond()
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<=100; i++{
		fmt.Println(rand.Intn(100))
	}
	fmt.Println(time.Now())
	//var end time.Time = time.Now()
	end := time.Now().Nanosecond()
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(end-start)
}