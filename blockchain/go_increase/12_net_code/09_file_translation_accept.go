package main

import (
	"fmt"
	"net"
	"os"
)

func recvFile(conn net.Conn, fileName string)  {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n_r, _ := conn.Read(buf)
		if n_r == 0 {
			fmt.Println("accept file finish!")
			return
		}
		// 写入本地文件，读多少，写多少
		f.Write(buf[:n_r])
	}
}

func HandlerConnect(conn net.Conn)  {
	defer conn.Close()
	buf := make([]byte, 4096)

	n, err := conn.Read(buf)

	if n == 0 {
		fmt.Println("translation has been finished!")
		return
	}

	if err != nil {
		fmt.Println("read err:", err)
		return
	}

	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("write err:", err)
		return
	}

	// 获取文件内容
	recvFile(conn, string(buf[:n]))
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()

	// 阻塞监听
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			return
		}
		defer conn.Close()
		go HandlerConnect(conn)
	}
}
