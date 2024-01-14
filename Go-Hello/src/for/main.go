package main

import "fmt"

// 2.5 流程控制
// for循环
func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 2. 省略初识语句
	// var i = 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 3.省略初始语句和结束语句
	// var i = 10
	// for i > 0 {
	// 	fmt.Println(i)
	// 	i--
	// }

	// 4.无限循环
	// for{
	// 	fmt.Println("hello")
	// }

	// 5.break 跳出for循环
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// 6. continue 继续下一个循环
	for i := 0; i < 5; i++ {
		// fmt.Println(i)
		if i == 3 {
			continue //跳过本次循环
		}
		fmt.Println(i)
	}

}
