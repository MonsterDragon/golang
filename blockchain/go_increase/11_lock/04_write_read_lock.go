package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex
var rwMutes_2 sync.RWMutex
//var mutex_2 sync.Mutex

func writeGo(ch chan<- int, num int) {
	// 生成随机数
	//mutex_2.Lock()
	//rwMutex.Lock()
	rwMutes_2.Lock()
	id := rand.Intn(1000)
	ch <- id
	fmt.Printf("%dth write go, write: %d\n", num, id)
	time.Sleep(time.Millisecond*300)
	rwMutes_2.Unlock()
	//rwMutex.Unlock()
	//mutex_2.Unlock()
}

func readGo(ch <-chan int, num int, quit chan <- int)  {
	rwMutex.RLock()
	if num <= 4 {
		id := <-ch
		fmt.Printf("%dth read go, read: %d\n", num, id)
	} else if num == 5 {
		time.Sleep(time.Second * 2)
		id := <-ch
		fmt.Printf("%dth read go, read: %d\n", num, id)
		quit <- 1
	}
	rwMutex.RUnlock()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan int) // 用于 关闭主go程的channel

	ch := make(chan int) // 用于 传递数据的channel

	for i:=0; i<5; i++ {
		go writeGo(ch, i+1)
	}

	for i:=0; i<5; i++ {
		go readGo(ch, i+1, quit)
	}
	<- quit
}
