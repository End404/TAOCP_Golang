package main

import (
	"bufio"
	"fmt"
	"net"
)

//网络编程；TCP服务端 demo

func process(conn net.Conn) {
	defer conn.Close() //处理完关闭连接
	for {              //数据的发送接受操作
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("读取错误：", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("接受到数据：", recv)
		conn.Write([]byte("ok")) //收到的数据返回客户端
	}
}

func main() {
	//1.开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	for { //等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			println("错误：", err)
			continue
		}
		//3.启动一个单独的goroutine去处理
		go process(conn)
	}
}
