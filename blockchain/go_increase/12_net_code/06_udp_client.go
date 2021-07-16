package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:4555")

	if err != nil {
		fmt.Println("dial err: ", err)
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
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
