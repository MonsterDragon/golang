package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("aa.txt", os.O_RDWR, 6) // 按行去读
	if err != nil {
		fmt.Println("open file error!")
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b_list, err1 := r.ReadBytes('\n')
		if err1 != nil && err1 == io.EOF {
			fmt.Println("read finish")
			return
		} else if err1 != nil {
			fmt.Println("read error")
		}
		fmt.Printf("%s\n", b_list)
	}
}
