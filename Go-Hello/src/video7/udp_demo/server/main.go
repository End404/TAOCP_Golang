package main

import (
	"fmt"
	"net"
)

//udp服务端

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		println("连接错误：", err)
		return
	}
	defer listen.Close()
	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("读取udp错误：", err)
			return
		}
		fmt.Println("接受到数据：", string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Println("写入udp错误：", addr, err)
			return
		}
	}
}
