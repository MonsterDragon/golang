package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Println("please enter dir name!")
	fmt.Scan(&s)

	d, err := os.OpenFile(s, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open dir error")
		return
	}
	defer d.Close()

	f_list, err1 := d.Readdir(-1) // -1表示全部进行
	if err1 != nil {
		fmt.Println("readdir fail")
		return
	}
	for _, info := range f_list {
		fmt.Println(info.Name())
		if info.IsDir() {
			fmt.Println("is a dir")
		} else {
			fmt.Println("is a file")
		}
	}
}
