package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Println("Please enter a string!")
	fmt.Scan(&str)
	n := strings.Index(str, "e")
	fmt.Println(n)
}
