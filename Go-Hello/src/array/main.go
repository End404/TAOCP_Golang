package main

import "fmt"

//第3章 复合数据类型
// 3.1 数组
func main() {
	// var a [3]int
	// var b [4]int
	// fmt.Println(a)
	// fmt.Println(b)

	// 数组的初始化
	//1.定义时使用初始值列表的方式初始化
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// fmt.Println(cityArray)
	// fmt.Println(cityArray[3])

	// // 2.编译器推导数组的长度
	// var boolArray = [...]bool{true, false, true, false, true}
	// fmt.Println(boolArray)

	// // 3.使用索引值方式初始化
	// var langArray = [...]string{1: "Go", 3: "Python", 7: "Java"}
	// fmt.Println(langArray)
	// fmt.Printf("%T\n", langArray)

	//数组的遍历
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// 1.for循环遍历
	// for i := 0; i < len(cityArray); i++ {
	// 	fmt.Println(cityArray[i])
	// }
	// 2. for range遍历
	// for _, value := range cityArray {
	// 	fmt.Println(value)
	// }

	// 二维数组
	// cityArray := [4][2]string{
	// 	{"bj", "xa"},
	// 	{"sh", "cq"},
	// 	{"hz", "cd"},
	// 	{"gz", "sz"},
	// }
	// fmt.Println(cityArray)
	// fmt.Println(cityArray[2][0])
	// // 二维数组的遍历
	// for _, v1 := range cityArray {
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	// 数组是值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
}

func f1(a [3]int) {
	a[0] = 100
}
