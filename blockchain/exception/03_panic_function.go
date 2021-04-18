package main

import "fmt"

/*
error返回的是一般性的错误，但是panic函数返回的是让程序崩溃的错误。
也就是当遇到不可恢复的错误状态的时候，如数组访问越界、空指针引用等，这些运行时错误会引起painc异常，在一般情况下，我们不应通过调用panic函数来报告普通的错误，而应该只把它作为报告致命错误的一种方式。当某些不应该发生的场景发生时，我们就应该调用panic。
*/

func TestA() {
	fmt.Println("func TestA()")
}

func TestB() {
	panic("func TestB(): panic")
}

func TestC() {
	fmt.Println("func TestC()")
}

func main() {
	TestA()
	TestB()
	TestC()
}
