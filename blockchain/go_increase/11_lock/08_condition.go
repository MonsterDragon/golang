package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func producer_08(in chan<- int, idx int)  {
	for  {
		cond.L.Lock() // 一把互斥锁，临界资源互斥访问，先上锁
		for len(in) == 5 {
			cond.Wait() // 挂起当前协程，等待条件变量满足，被消费者唤醒 p
		}
		num := rand.Intn(1000)
		in <- num
		fmt.Printf("生产者%dth，生产：%d\n", idx, num)
		cond.L.Unlock() // 生产结束，解锁互斥锁
		cond.Signal() // 唤醒阻塞的消费者 v
		time.Sleep(time.Millisecond*300)
	}
}

func consumer_08(out <-chan int, idx int)  {
	for {
		cond.L.Lock()
		for len(out) == 0 {
			cond.Wait()
		}
		num := <- out
		fmt.Printf("消费者%dth, 消费：%d\n", idx, num)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond*300)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int , 5)

	cond.L = new(sync.Mutex)
	for i:=0; i<5; i++{
		go producer_08(ch, i+1)
	}

	for i:=0; i<5; i++ {
		go consumer_08(ch, i+1)
	}

	for {
		;
	}
}
