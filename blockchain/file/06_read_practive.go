package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fp, err := os.Open("bb.txt")
	if err != nil {
		fmt.Println("fail to open fp")
		return
	}
	defer fp.Close()

	//b := make([]byte, 200)
	//fp.Read(b)
	//fmt.Println(string(b)) // 只读取200个字节大小

	//for {
	//	n, err := fp.Read(b)
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//	}
	//	fmt.Print(string(b[:n]))
	//} // 按行读取

	r := bufio.NewReader(fp)
	for {
		c, err := r.ReadBytes('\n') // 返回出现第一次分隔符前的所有数据组成的字节切片
		fmt.Print(string(c))
		if err == io.EOF {
			break
		}
	}
}
