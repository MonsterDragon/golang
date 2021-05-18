package main

import "fmt"

func main() {
	ch := make(chan int) // 内建函数make(T, args)与new(T)的用途不一样。它只用来创建slice，map和channel，并且返回一个初始化的(而不是置零)，类型为T的值（而不是*T）

	go func() {
		var sendCH chan<- int = ch
		sendCH <- 1
	}()

	var recvCH <- chan int = ch
	n := <- recvCH
	fmt.Println(n)
}
