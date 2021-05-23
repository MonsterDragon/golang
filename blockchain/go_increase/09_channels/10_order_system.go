package main

import "fmt"

type OrderInfo struct {
	id int
}

func producer2(in chan <- OrderInfo)  {
	for i:=0; i<10; i++ {
		fmt.Println("producer write ", i+1)
		order := OrderInfo{id: i+1}
		in <- order
	}
	close(in)
}

func consumer2(out <- chan OrderInfo)  {
	for order := range out {
		fmt.Println("order id is ", order)
	}
}

func main() {
	ch := make(chan OrderInfo)

	go producer2(ch)
	go consumer2(ch)

	for  {
		;
	}
}
