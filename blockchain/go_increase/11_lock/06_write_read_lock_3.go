package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex_4 sync.RWMutex
var value int

func writeGo_3(idx int)  {
	//for {
	rwMutex_4.Lock()
	num := rand.Intn(1000)

	value = num
	fmt.Printf("%d write go, write: %d\n", idx, num)
	time.Sleep(time.Millisecond * 300)
	rwMutex_4.Unlock()
	//}
}

func readGo_3(idx int) {
	for {
	rwMutex_4.RLock()
	num := value
	fmt.Printf("----%d read go, read : %d\n", idx, num)
	rwMutex_4.RUnlock()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//ch := make(chan int)

	for i:=0; i<5; i++ {
		go readGo_3(i+1)
	}

	for i:=0; i<5; i++ {
		go writeGo_3(i+1)
	}

	for {
		;
	}
}

