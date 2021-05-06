package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Please a statement: ")
	fmt.Scan(&str)
	if strings.Contains(str, "evil") {
		str_1 := strings.Replace(str, "evil", "**", -1)
		fmt.Println(str_1)
	}
}
