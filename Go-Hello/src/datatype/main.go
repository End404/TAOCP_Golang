package main

import (
	"fmt"
)

/* 第2章
2.3 基本数据类型
*/

func main() {
	//十进制打印为二进制
	// n := 10
	// fmt.Printf("%b\n", n)
	// fmt.Printf("%d\n", n)

	// // 八进制
	// m := 075
	// fmt.Printf("%d\n", m)
	// fmt.Printf("%o\n", m)

	// // 十六进制
	// f := 0xff
	// fmt.Println(f)
	// fmt.Printf("%x\n", f)

	// //浮点数
	// fmt.Println(math.MaxFloat32)
	// fmt.Println(math.MaxFloat64)

	// //布尔值
	// var a bool
	// fmt.Println(a)
	// a = true
	// fmt.Println(a)

	//字符串
	s1 := "hello"
	s2 := "你好"
	fmt.Println(s1, s2)

	//打印路径
	fmt.Println("D:\\ProgramData\\Go-Hello")
	fmt.Println("\t制表符\n换行")
	s3 := `
	多好字符串
	反引号里的内容
	会
	原样
	\t
	\n
	`
	fmt.Println(s3)

}
