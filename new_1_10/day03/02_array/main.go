package main

import (
	"fmt"
)

// 存放元素的容器
// 必须指定存放元素的类型和容量（长度）
// 数组长度是数组类型的一部分

var a1 [3]bool
var a2 [4]bool

func search(num int){
	var d1 [5]int
	d1 = [5]int{1, 3, 5, 7, 9}

	for i:=0; i<5; i++{
		for j:=i+1; j<5; j++{
			if d1[i] + d1[j] == num{
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}

func main() {
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//数组的初始化:默认元素都是0
	fmt.Println(a1, a2)
	// 1、初始化方式
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	a100 := [100]int{1, 2, 43}
	// 2、自动判断数组长度
	fmt.Println(a100)
	a10 := [...]int{3, 7, 10}
	fmt.Println(a10)
	// 3、根据索引初始化
	a3 := [5]int{0:100, 3:200}
	fmt.Println(a3)
	citys := [...]string{"beijing", "shanghai", "shenzhen"}
	for i:=0; i<len(citys); i++{
		fmt.Println(citys[i])
	}
	// 多维数组
	var a11 [3][4]int
	a11 = [3][4]int{
		[4]int{10, 13},
		[4]int{0, 1, 2},
		[4]int{100, 101, 102},
	}
	fmt.Println(a11)
	for num, v := range a11{
		fmt.Println(num, v)
		for num, v1 := range v{
			fmt.Println(num, v1)
		}
	}

	b1 := [...]int{1, 2, 3}
	b2 := b1
	fmt.Printf("%p, %p\n", &b1, &b2)

	c1 :=  1
	fmt.Printf("%v\n", &c1)
	// 数组支持 == !=
	// [n]*T表示指针数组 *[n]T表示数组指针

	var d1 [5]int
	var count int
	d1 = [5]int{1, 3, 5, 7, 9}
	for _, v := range d1{
		count += v
	}
	fmt.Println(count)

	search(8)
}
