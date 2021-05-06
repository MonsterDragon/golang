package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "abc@123.com"
	s := strings.Split(str, "@")
	fmt.Println("username: ", s[0])
	fmt.Println("place: ", s[1])
}
