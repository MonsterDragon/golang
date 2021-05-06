package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 一、将其他类型转为字符串
	// 1、将bool转换为字符串
	s := strconv.FormatBool(true)
	fmt.Println(s)

	// 2、将int转换为字符串
	a := strconv.FormatInt(12, 2) // 将12转为2进制
	fmt.Println(a)

	b := strconv.Itoa(123) // 只有10进制
	fmt.Println(b)

	// 3、将浮点型转换为字符串（会进行四舍五入）
	// f指打印格式，表示小数方式，3指小数点位数，64以float64处理
	c := strconv.FormatFloat(3.14, 'f', 3, 64)
	fmt.Println(c)

	// 二、将字符串转换为其他类型
	// 1、将bool字符串转换为bool
	d, _ := strconv.ParseBool("false")
	fmt.Println(d)

	// 2、将int类型转换为int
	e, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(e)
	f, _ := strconv.Atoi("1234")
	fmt.Println(f)

	// 3、将浮点字符串转换成浮点型
	g, _ := strconv.ParseFloat("12.45", 64)
	fmt.Println(g)

	// 三、将整数等转换为字符串之后，添加到现有的字节数组中
	// append系列函数将整数等转换为字符串后，添加到现有的字节数组中

	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgd")
	slice = strconv.AppendFloat(slice, 1.234, 'f', 5, 64)
	fmt.Println(slice)
	fmt.Println("slice = ", string(slice))
}
