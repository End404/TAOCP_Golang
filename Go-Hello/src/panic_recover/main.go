package main

import "fmt"

// panic recover
func a() {
	fmt.Println("函数 a")
}
func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("函数 b 错误")
		}
	}()
	panic("panic in b")
}
func c() {
	fmt.Println("函数 c")
}

func main() {
	a()
	b()
	c()
}
