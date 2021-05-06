package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "{“诸葛亮”,”鸟叔”,”卡卡西”,”卡哇伊”}"
	fmt.Println(str)
	str = strings.Trim(str, "{}")
	fmt.Println(str)
	aa := strings.Split(str, ",")
	fmt.Println(aa)
	bb := make([]string, 4, 8)
	for i:=0; i<len(aa); i++ {
		bb[i] = strings.Trim(aa[i], "“”")
	}
	fmt.Println(bb)
	var str_1 string
	str_1 = strings.Join(bb, "|")
	fmt.Println(str_1)
	//for i:=0; i<len(bb); i++ {
	//	str_1 = strings.Join(str_1, bb[i])
	//}
	str_1 = strings.Replace(str_1, "|", "", -1)
	fmt.Println(str_1)
}
