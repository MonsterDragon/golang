package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello"
	value := strings.Contains(str, "llo")
	fmt.Println(value)

	str_1 := []string{"abcd", "efgh"}
	str_2 := strings.Join(str_1, "-")
	fmt.Println(str_2)

	str_3 := "llo"
	n := strings.Index(str, str_3)
	fmt.Println(n)

	s := strings.Repeat(str, 2) // 字符串重复
	fmt.Println(s)

	m := strings.Replace(str, "l", "w", 2) // 小于0全部替换
	fmt.Println(m)

	str_4 := "www.baidu.com"
	a := strings.Split(str_4, ".") // 返回切片
	fmt.Println(a)
	for i:=0; i<len(a); i++ {
		fmt.Println(a[i])
	}

	str_5 := "     hello     world      "
	b := strings.Trim(str_5, " ") // 去除字符串前后指定的内容
	fmt.Println(b)

	str_6 := "   are you ok    "
	c := strings.Fields(str_6) //["are", "you", "ok"]
	fmt.Println(c)
	for i:=0; i<len(c); i++ {
		fmt.Println(c[i])
	}
}
