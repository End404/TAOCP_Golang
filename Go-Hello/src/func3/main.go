package main

import "fmt"

//匿名函数和闭包

func a() func() { //定义一个函数的返回值是一个函数
	name := "娜扎"
	return func() {
		fmt.Println("hello", name)
	}
}

func main() {
	// sayHell := func() {
	// 	fmt.Println("匿名函数")
	// }
	// sayHell()

	// func() {
	// 	fmt.Println("匿名函数")
	// }()

	r := a() //闭包 = 函数+外层变量的引用
	r()      //执行a函数内部的匿名函数
}
