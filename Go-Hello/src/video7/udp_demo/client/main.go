package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//udp客户端

func main() {
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	defer c.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		_, err = c.Write([]byte(s))
		if err != nil {
			fmt.Println("发送错误：", err)
			return
		}
		//接受数据
		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			println("读取udp错误：", err)
			return
		}
		fmt.Printf("读取来自 %v, mag:%v\n", addr, string(buf[:n]))
	}
}
