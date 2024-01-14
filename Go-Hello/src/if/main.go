package main

import "fmt"

// 2.5 流程控制
// if判断
func main() {
	// 1.基本写法
	// var score = 65
	// if score >= 90 {
	// 	fmt.Println("A")
	// } else if score > 75 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("C")
	// }

	//2.if判断的特殊写法
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	if 10 > 5 {
		fmt.Println("10>5")
	}
}
