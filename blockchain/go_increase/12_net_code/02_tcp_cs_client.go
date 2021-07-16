package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	defer conn.Close()

	//buf := make([]byte, 4096)
	//fmt.Println("enter message: ")
	//fmt.Scan(&buf)
	//conn.Write(buf)
	conn.Write([]byte("exit"))
}
