package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love my work and I love my family too"
	str_1 := strings.Split(str, " ") // 按照空格进行拆分，返回切片
	fmt.Println(str_1)
	for _, i := range str_1 {
		fmt.Println(i)
	}

	// 判断字符串时候前后缀
	flg := strings.HasSuffix("test.abc", ".abc")
	fmt.Println(flg)

	flg_1 := strings.HasPrefix("test.abc", "test")
	fmt.Println(flg_1)
}
