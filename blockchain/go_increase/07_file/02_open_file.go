package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("aa.txt") // 以只读的方式进行打开
	if err != nil {
		fmt.Println("open file error!")
		return
	}

	_, err1 := f.WriteString("aaaaaaa")
	if err1 != nil {
		fmt.Println("write string error!")
		fmt.Println(err1)
		return
	}
	defer f.Close()
}
