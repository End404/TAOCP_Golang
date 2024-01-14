package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件操作

func readFrile() {
	//打开文件
	fileObj, err := os.Open("D:\\ProgramData\\Go-Hello\\src\\video4\\xx.txt") //路径
	if err != nil {
		fmt.Printf("打开文件出错, err:%v\n", err)
		return
	}

	defer fileObj.Close()
	//读取文件内容
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取文件字节：", n)
	fmt.Println(string(tmp))
}

func readAll() {
	//打开文件
	fileObj, err := os.Open("D:\\ProgramData\\Go-Hello\\src\\video4\\xx.txt") //路径
	if err != nil {
		fmt.Printf("打开文件出错, err:%v\n", err)
		return
	}

	defer fileObj.Close()

	//读取文件内容
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { //End Of File
			//当前虚弱字节的数据打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Println("读取文件失败：", err)
			return
		}
		fmt.Println("读取文件字节：", n)
		fmt.Println(string(tmp[:n]))
	}
}

// read by bufio
func readByfio() {
	//打开文件
	fileObj, err := os.Open("D:\\ProgramData\\Go-Hello\\src\\video4\\xx.txt") //路径
	if err != nil {
		fmt.Printf("打开文件出错, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	line, err := reader.ReadSlice('\n')
	if err != nil {
		fmt.Println("读取文件 by bufio 出错：", err)
		return
	}
	fmt.Println(line)
}

// 写入操作
func write() {
	fileObj, err := os.OpenFile("D:\\ProgramData\\Go-Hello\\src\\video4\\xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("打开放假出错：%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "小王子"
	fileObj.Write([]byte(str))
	fileObj.WriteString("hello 沙河")
}

func main() {
	//readAll()
	//readByfio()
	write()
}
