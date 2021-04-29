package main

import (
	"fmt"
	"os"
)

func CreateFile(path string) {
	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()
}

func main() {
	var file_path string = "a.txt"
	CreateFile(file_path)
}
