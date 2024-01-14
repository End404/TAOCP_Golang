package main

import "fmt"

// 函数

func sayHello() { //没有参数，没有返回值
	fmt.Println("Hello 小王子")
}

func sayHello2(name string) { //参数
	fmt.Println("Hai", name)
}

func intSun(a int, b int) int { //多参数，一个返回值
	ret := a + b
	return ret
}
func intSun2(a int, b int) (ret int) { //多参数，一个返回值
	ret = a + b
	return
}

func intSun3(a ...int) int { //可变参数
	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	return ret
}
func intSun4(a int, b ...int) int { //固定参数，可变参数
	ret := a
	for _, arg := range b {
		ret = ret + arg
	}
	return ret
}
func calc(a, b int) (sum, sub int) { //多个返回值
	sum = a + b
	sub = a - b
	return
}

func main() {
	// sayHello() //函数调用

	// name := "沙河"
	// sayHello2(name)

	// r := intSun(10, 20)
	// fmt.Println(r)

	// r1 := intSun3()
	// r2 := intSun3(10)
	// r3 := intSun3(10, 20)
	// fmt.Println(r1)
	// fmt.Println(r2)
	// fmt.Println(r3)

	// r1 := intSun4(0)
	// r2 := intSun4(10)
	// r3 := intSun4(10, 20)
	// fmt.Println(r1)
	// fmt.Println(r2)
	// fmt.Println(r3)

	x, y := calc(100, 200)
	fmt.Println(x, y)
}
