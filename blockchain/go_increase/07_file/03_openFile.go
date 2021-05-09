package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("aa.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("open file error!")
		return
	}

	_, err1 := f.WriteString("bbbbbb")
	if err1 != nil {
		fmt.Println("write file err!")
		return
	}
	defer f.Close()
}
