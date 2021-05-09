package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("aa.txt")
	if err != nil {
		fmt.Println("open file err!")
		return
	}
	defer f.Close()

	w, err1 := os.Create("bb.txt")
	if err1 != nil {
		fmt.Println("create fil err!")
		return
	}
	defer w.Close()

	d := make([]byte, 4*1024)
	for  {
		n, err2 := f.Read(d)
		if err2 != nil && n == 0 {
			fmt.Println("read finish")
			return
		} else if err2 != nil {
			fmt.Println("read fail")
			return
		}
		w.Write(d[:n])
	}
}