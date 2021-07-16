package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string)  {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("translate file finish!")
			} else {
				fmt.Println("read file err:", err)
			}
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("socket write err:", err)
			return
		}
	}
}

func tcpTranslate(filename string, filepath string) {
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()

	// 发送文件名给服务器
	conn.Write([]byte(filename))

	// 读取服务器发回的ok
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn read err:", err)
		return
	}

	if "ok" == string(buf[:n]) {
		sendFile(conn, filepath)
	}
}

func GetAllFile(filepath string) ([]string, error) {
	result := []string{}

	fs, err := ioutil.ReadDir(filepath)
	if err != nil {
		fmt.Printf("读取文件目录失败，filepath=%v，err=%v\n", filepath, err)
		return result, err
	}

	for _, f := range fs {
		fullname := filepath + "/" + f.Name()
		if f.IsDir() {
			temp, err := GetAllFile(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败，fullname=%v，err=%v\n", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			result = append(result, fullname)
		}
	}
	return result, nil
}

func main() {
	list := os.Args

	if len(list) != 2 {
		fmt.Println("err args")
		return
	}

	filepath := list[1]
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	if fileInfo.IsDir() {
		filelist, err := GetAllFile(filepath)
		fmt.Println(filelist)
		if err != nil {
			fmt.Println("terminal the execution!")
			return
		}
		for _, f := range filelist {
			fileInfo, err := os.Stat(f)
			if err != nil {
				fmt.Println("os.Stat err:", err)
				return
			}
			tcpTranslate(fileInfo.Name(), f)
		}
	} else {
		filename := fileInfo.Name()
		tcpTranslate(filename, filepath)
	}
}
