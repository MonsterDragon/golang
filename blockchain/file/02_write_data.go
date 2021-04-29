package main

import (
	"fmt"
	"os"
)

// WriteString方法
//func CreateFile_1(path string) {
//	f, err := os.Create(path)
//	if err != nil {
//		fmt.Println("err = ", err)
//		return
//	}
//	for i := 1; i < 10; i++ {
//		n, err := f.WriteString("hello world\n") // 返回写入数据的长度和错误信息
//		fmt.Println(n)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//	defer f.Close()
//}

// write方法
//func CreateFile_1(path string)  {
//	f, err := os.Create(path)
//	if err != nil {
//		fmt.Println("err = ", err)
//		return
//	}
//	var str string
//	for i := 0; i < 10; i++ {
//		str = fmt.Sprintf("i = %d\n", i)
//		buf := []byte(str)
//		n, err := f.Write(buf)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(n)
//	}
//}

func CreateFile_1(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	var str string
	var a int
	for i := 0; i < 10; i++ {
		str = fmt.Sprintf("i = %d\n", i + 1)
		buf := []byte(str)
		n, _ := f.Seek(0, os.SEEK_END)
		fmt.Printf("n = %d\n", n)
		a, err = f.WriteAt([]byte(buf), n)
		fmt.Println(a)
	}
	defer f.Close()
}

func main() {
	var path string = "bb.txt"
	CreateFile_1(path)
}
