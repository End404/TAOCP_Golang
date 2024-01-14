package main

import "fmt"

//自定义类型 类型别名

// 1.自定义类型
type MyInt int //基于int类型的自定义类型

// 2.类型别名
type NewInt = int //int类型别名

func main() {
	var i MyInt
	fmt.Printf("type:%T value:%v\n", i, i)

	var x NewInt
	fmt.Printf("type:%T value:%v\n", x, x)
}
