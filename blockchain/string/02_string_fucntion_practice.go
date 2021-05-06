package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Please enter a date which's format is 'Y-M-D'")
	fmt.Scan(&str)
	s := strings.Split(str, "-")
	fmt.Println(s[0] + "年" + s[1] + "月" + s[2] + "日")
}
