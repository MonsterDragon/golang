package main

import "fmt"

func producer(in chan <- int)  {
	for i:=0; i<10; i++{
		in <- i
		fmt.Println("producer write ", i)
	}
	close(in)
}

func consumer(out <- chan int)  {
	for {
		if info, ok := <- out; ok == true {
			fmt.Println("consumer read ", info)
		} else if ok == false {
			fmt.Println("channel close")
			break
		}
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)

	for  {
		;
	}
}
