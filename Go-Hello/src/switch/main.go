package main

import "fmt"

// 2.5 流程控制
// switch

func main() {
	// finger := 3
	// switch finger {
	// case 1:
	// 	fmt.Println("da")
	// case 2:
	// 	fmt.Println("shi")
	// case 3:
	// 	fmt.Println("zhong")
	// default:
	// 	fmt.Println("wu")
	// }

	// case 一次判断多个值
	// num := 5
	// switch num {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("奇数")
	// case 2, 4, 6, 8:
	// 	fmt.Println("偶数")
	// }

	// 3.case中使用表达式
	age := 30
	switch {
	case age > 18:
		fmt.Println("开业")
	case age < 18:
		fmt.Println("关闭")
	default:
		fmt.Println("终止")
	}

}
