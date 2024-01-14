package main

import "fmt"

//函数进阶 变量作用域

var num = 10 //全局变量

func testGlobal() {
	// num := 100
	name := "娜扎"
	fmt.Println("变量", num)
	fmt.Println(name)
}

func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func sub(x, y int) int {
	return x - y
}

func main() {
	// testGlobal()
	// abc := testGlobal //函数可作为变量
	// fmt.Printf("%T\n", abc)
	// abc()

	r1 := calc(100, 200, add)
	fmt.Println(r1)
	r2 := calc(100, 200, sub)
	fmt.Println(r2)

}
