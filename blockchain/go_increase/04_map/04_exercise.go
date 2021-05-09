package main

import (
	"fmt"
	"strings"
)

func wordCountFunc(a string) map[string]int {
	string_map := make(map[string]int, 10)
	string_list := strings.Fields(a)
	for _, i := range string_list {
		if v, ok := string_map[i]; ok == true {
			string_map[i] = v + 1
		} else {
			string_map[i] = 1
		}
	}
	return string_map
}

func main() {
	var s string = "I love my work and I love my family too"

	terminal := wordCountFunc(s)
	fmt.Println(terminal)
}
