package main

import "fmt"

func main(){
	// 基本格式
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	// 变种1
	j := 5
	for ; j < 10; j++{
		fmt.Println(j)
	}

	// 变种2
	m := 5
	for ; m < 10; {
		fmt.Println(m)
		m++
	}

	// while TRUE{}；
	//for true {
	//	fmt.Println("a")
	//}

	// for range循环
	s := "HelloHello"
	for i, v := range s{
		fmt.Println(i, v)
	}

	for a:=1; a<10; a++{
		for b:=1; b<=a; b++{
			fmt.Printf("%d * %d = %d ", a, b, a * b)
		}
		fmt.Println()
	}
}
