package main

import "fmt"

func main() {
	var total int = 0
	// for循环是函数格式，i是内部定义的不能在外部使用
	for i:=1;i<=100;i++{
		total += i
	}
	fmt.Println(total)

	var sum int = 0
	for i:=1;i<=100;i++{
		if i%2 == 0{
			sum += i
		}
	}
	fmt.Println(sum)

	//i:=1
	//for ;i<=10;i++{
	//	fmt.Println(i)
	//}
	//fmt.Println(i)
	i := 1
	for ;;i++{
		if i>10 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println(i)
}
