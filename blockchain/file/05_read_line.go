package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Read_file(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	r := bufio.NewReader(f)

	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break // EOF -1 文件结束标志
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}
}

func main() {
	var path string = "bb.txt"
	Read_file(path)
}