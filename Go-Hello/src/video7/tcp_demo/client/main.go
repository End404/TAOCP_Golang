package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//TCP客户端

func main() {
	//1.与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	//2.利用该连接进行数据的发送和接受
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.TrimSpace(s) == "Q" {
			return
		}
		//给服务端发消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Println("发送错误：", err)
			return
		}
		//从服务端接受回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("读取错误：", err)
			return
		}
		fmt.Println("收到服务端回复：", string(buf[:n]))

	}
}
