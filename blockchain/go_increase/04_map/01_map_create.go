package main

import "fmt"

func main() {
	// 方法一：
	var m1 map[int]string // 只是声明一个map，没有初始化，为空(nil)map
	//m1[] = "shuzhan"
	fmt.Println(m1)
	// 方法二：
	m2 := map[int]string{}
	// 方法三：
	m3 := make(map[int]string)
	fmt.Println(m2, m3)
	// 方法四：
	m4 := make(map[int]string, 5) // 不能在map中使用cap
	fmt.Println(len(m4))

	var s1 map[int]string = map[int]string{1:"aa", 2:"bb"}
	fmt.Println(s1)
	s2 := map[int]string{1:"cc", 2:"dd"}
	fmt.Println(s2)

	m5 := make(map[int]string, 7)
	m5[10] = "haha"
	m5[2] = "worle"
	fmt.Println(m5)
}
