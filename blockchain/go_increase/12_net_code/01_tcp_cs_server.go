package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	fmt.Println("addr listener: ", &listener)
	defer listener.Close()

	fmt.Println("server wait client to connect")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	defer conn.Close()

	fmt.Println("addr conn: ", &conn)
	fmt.Println("connection has been established!")
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("server read data: ", string(buf[:n]))
}
