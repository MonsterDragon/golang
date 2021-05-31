package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func printer_2(str string)  {
	mutex.Lock() // 访问共享数据之前加锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()
}

func pers_1() {
	printer_2("hello")
}

func pers_2()  {
	printer_2("world")
}

func main() {
	go pers_1()
	go pers_2()
	for {
		;
	}
}
