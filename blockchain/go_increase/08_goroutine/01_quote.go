package main

import (
	"fmt"
	"time"
)

func sing() {
	for i:=0; i<10; i++{
		fmt.Println("sing a song")
		time.Sleep(100 * time.Millisecond)
	}
}

func dance() {
	for i:=0; i<10; i++ {
		fmt.Println("last dance")
		time.Sleep(100*time.Millisecond)
	}
}

func main() {
	fmt.Println("start")
	go sing()
	go dance()
	//sing()
	//dance()
	time.Sleep(1*time.Second)
	fmt.Println("stop")
}
