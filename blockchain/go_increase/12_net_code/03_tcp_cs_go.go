package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if "exit\n" == string(buf[:n]) {
			fmt.Println("client request to disconnect")
			return
		}
		if n == 0 {
			fmt.Println("client has been closed! disconnected")
			return
		}
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}
		fmt.Printf("read message is: %s\n", string(buf[:n]))
		// 小写转大写，回发客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n-1]))))
	}
}

func main() {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}
	defer listener.Close()

	// 监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.accept err: ", err)
			return
		}

		// 具体完成通信过程
		go HandlerConnect(conn)
	}
}
