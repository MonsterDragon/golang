package main

import "fmt"

func Send(out chan <- int) {
	out <- 12
	close(out)
}

func Recv(in <- chan int)  {
	n := <- in
	fmt.Println(n)
}

func main() {
	ch := make(chan int)

	go func() {
		Send(ch)
	}()

	Recv(ch)
}
