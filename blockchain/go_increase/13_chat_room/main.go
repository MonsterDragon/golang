package main

import (
	"fmt"
	"net"
)

// 创建结构体类型
type Client struct {
	C    chan string
	Name string
	Addr string
}

// 创建全部map，存储在线用户
var onlineMap map[string]Client

// 创建全局channel传递用户信息
var message = make(chan string)

// 将客户map的channel中的信息发送给客户
func WriteMsgToClient(n_cli Client, conn net.Conn) {

	for msg := range n_cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	// 获取用户网络地址ip+port
	netAddr := conn.RemoteAddr().String()
	// 创建新用户的变量
	n_cli := Client{make(chan string), netAddr, netAddr}
	// 将新连接的用户，添加到在线用户map中 key:ip:port value:client
	onlineMap[netAddr] = n_cli

	go WriteMsgToClient(n_cli, conn)

	// 将用户上线的信息写入到message channel中
	message <- "[" + netAddr + "]" + n_cli.Name + "login"

	for {

	}
}

func manager() {
	// 初始化onlinemap
	onlineMap = make(map[string]Client)

	for {
		// 监听message中是否有新用户的数据，有数据存储到c_message，无数据阻塞
		c_message := <-message

		// 循环发送消息给在线用户
		for _, cli := range onlineMap {
			cli.C <- c_message
		}
	}
}

func main() {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:7770")
	if err != nil {
		fmt.Println("Listen err", err)
		return
	}
	defer listener.Close()

	// 创建管理者go程，管理map和全局channel
	go manager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
			return
		}
		// 启动go程处理客户端请求
		go HandlerConnect(conn)
	}
}
