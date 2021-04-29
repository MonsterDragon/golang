package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*2)

	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Println("buf = ", string(buf[:n]))
}

func main() {
	var path string = "bb.txt"
	ReadFile(path)
}