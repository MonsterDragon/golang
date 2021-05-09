package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("aa.txt")
	if err != nil {
		fmt.Println("create file error!")
		return
	}
	defer f.Close()
}
