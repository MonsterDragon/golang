package main

import (
	"fmt"
)

func main() {
	var p [2]*int
	var i int = 10
	var j int = 20
	p[0] = &i
	p[1] = &j
	for index, value := range(p) {
		fmt.Printf("index=%d, value=%p\n", index, value)
	}
}
