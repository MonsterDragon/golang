package main

import "fmt"

func Test_A() {
	fmt.Println("func TestA()")
}

func Test_B(x int)  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 111
}

func Test_C() {
	fmt.Println("func TestC()")
}

func main() {
	Test_A()
	Test_B(9)
	Test_C()
}
