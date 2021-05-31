package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex_3 sync.RWMutex

func writeGo_2(out chan<- int, idx int)  {
	//for {
		num := rand.Intn(1000)
		rwMutex_3.Lock()
		out <- num
		fmt.Printf("%d write go, write: %d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
		rwMutex_3.Unlock()
	//}
}

func readGo_2(in <-chan int, idx int) {
	//for {
		rwMutex_3.RLock()
		num := <-in
		fmt.Printf("%d read go, read : %d\n", idx, num)
		rwMutex_3.RUnlock()
	//}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	for i:=0; i<5; i++ {
		go readGo_2(ch, i+1)
	}

	for i:=0; i<5; i++ {
		go writeGo_2(ch, i+1)
	}

	for {
		;
	}
}
