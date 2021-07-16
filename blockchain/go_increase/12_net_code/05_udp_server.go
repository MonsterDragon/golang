package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	srvAddr, err1 := net.ResolveUDPAddr("udp", "127.0.0.1:4555")
	if err1 != nil {
		fmt.Println("resolveudpaddr err:", err1)
		return
	}

	uConn, err2 := net.ListenUDP("udp", srvAddr)
	if err2 != nil {
		fmt.Println("listenudp err:", err2)
		return
	}
	defer uConn.Close()

	buf := make([]byte, 4096)
	for  {
		n, cltAddr, err3 := uConn.ReadFromUDP(buf)
		if err3 != nil {
			fmt.Println("Readfromudp err:", err3)
			return
		}

		fmt.Printf("read %v: %s\n", cltAddr, buf[:n])

		current_time := time.Now().String()
		_, err := uConn.WriteToUDP([]byte(current_time), cltAddr)
		if err != nil {
			fmt.Println("writeToUdp err:", err)
			return
		}
	}
}
