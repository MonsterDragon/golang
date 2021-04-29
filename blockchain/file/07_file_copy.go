package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	fp, err := os.Open("bb.txt")
	if err != nil {
		fmt.Println("failed to open file")
		return
	}
	defer fp.Close()


	fp_2, err_2 := os.Create("cc.txt")
	if err_2 != nil {
		fmt.Println("create file failed!")
		return
	}
	defer fp_2.Close()



	r := bufio.NewReader(fp)
	for {
		c, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		n, err_3 := fp_2.Write(c)
		if err_3 != nil {
			fmt.Println("write file failed!")
			return
		}
		fmt.Println(n)
	}
}
