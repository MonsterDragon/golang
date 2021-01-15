package main

import "fmt"

/*
Go 语言的字符有以下两种：

uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
rune类型，代表一个 UTF-8字符。
*/
func main(){
	s := "hello你好"
	for i := 0; i < len(s); i++{ // len()获取字符串字节长度
		fmt.Printf("%v(%c)", s[i], s[i])
	}

	fmt.Println(len(s))

	for _, r := range s{
		fmt.Printf("%c\n", r)
	}

	// 字符串修改
	s2 := "你好"
	s3 := []rune(s2)
	s3[0] = '歪'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T, c2:%T", c1, c2) // c1:string, c2:int32

	// 类型转转
	n := 10
	var f float64
	f = float64(n)
	fmt.Println()
	fmt.Printf("%T\n%f", f, f)
}
