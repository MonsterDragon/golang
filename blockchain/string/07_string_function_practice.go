package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fp, err := os.Open("book.txt")
	if err != nil {
		fmt.Println("Failed to open file!")
		return
	}
	defer fp.Close()

	r := bufio.NewReader(fp)

	for {
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		str := string(b[:len(b)-1])
		//c := strings.Trim(str, "\n")
		//fmt.Println(str)
		d := strings.Fields(str)
		fmt.Println(d)
		//for i:=0; i<len(d); i++ {
		//	fmt.Println(d[i])
		//}
		for index, str_1 := range d {
			fmt.Println(index, str_1)
		}
	}
}
