package main

import (
	"Go-Hello/src/package_demo/calc"
	"fmt"
	// nazha "package_demo/calc"	//启用别名
)

func init() { //初始化的操作
	fmt.Println("main.init()")
}

func main() {
	fmt.Println("hello")
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
