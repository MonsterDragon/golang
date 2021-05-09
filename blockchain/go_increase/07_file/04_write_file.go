package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("aa.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("open file error!")
		return
	}
	defer f.Close()

	n, err1 := f.WriteString("shu")
	if err1 != nil {
		fmt.Println("write file error")
		return
	}
	fmt.Println("n: ", n)

	off, _ := f.Seek(-5, io.SeekEnd)
	fmt.Println(off)

	off1, _ := f.Seek(5, io.SeekStart)
	fmt.Println(off1)
	fmt.Printf("%T\n", off1)

	var b_list []byte = []byte{'a', 'b', 'c'}
	n1, err2 := f.WriteAt(b_list, off1)
	fmt.Println(n1)
	fmt.Println(err2)
}
