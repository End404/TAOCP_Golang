package main

import "fmt"

//延迟执行
func main() {
	fmt.Println("开始...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println("结束...")
}
