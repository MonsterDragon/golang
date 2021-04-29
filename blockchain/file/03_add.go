package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	f, err := os.OpenFile(path, os.O_APPEND, 6) // 路径、追加、权限
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	n, err1 := f.WriteString("Hello\n") // 写入数据的长度 错误信息
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(n)
	defer f.Close()
}

func main() {
	var path string = "bb.txt"
	WriteFile(path)
}