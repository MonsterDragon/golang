package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("dial error: ", err)
		return
	}
	defer conn.Close()

	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("read keyboard: ", err)
				continue
			}
			conn.Write(str[:n])
		}
	}()

	buf := make([]byte, 4096)
	for  {
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("read err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}