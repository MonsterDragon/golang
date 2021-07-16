package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args

	if len(list) != 2 {
		fmt.Println("wrong arguments")
		return
	}

	path := list[1]

	fileinfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	fmt.Println("file name: ", fileinfo.Name())
}
