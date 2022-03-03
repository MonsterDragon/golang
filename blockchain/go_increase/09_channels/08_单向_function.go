package main

import "fmt"

func Send(out chan <- int) {
	out <- 12
	fmt.Println("aa")
	close(out)
}

func Recv(in <- chan int)  {
	n := <- in
	fmt.Println(n)
}

func main() {
	ch := make(chan int) //内建函数make(T, args)与new(T)的用途不一样。它只用来创建slice，map和channel，并且返回一个初始化的(而不是置零)，类型为T的值（而不是*T）

	go func() {
		Send(ch)
	}()

	Recv(ch)
	for  {
		;
	}
}
