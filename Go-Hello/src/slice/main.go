package main

import "fmt"

// 3.2 切片
func main() {
	// var a []string
	// var b []int

	// var c = []bool{false, true}
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// 基于数组得到切片
	// a := [5]int{55, 56, 57, 58, 59}
	// b := a[1:4]
	// fmt.Println(b)
	// fmt.Printf("%T\n", b)
	// // 切片再次切片
	// c := b[0:len(b)]
	// fmt.Println(c)
	// // make函数构造切片
	// d := make([]int, 5, 10)
	// fmt.Println(d)
	// fmt.Printf("%T\n", d)
	// //获取切片的长度
	// fmt.Println(len(d))
	// //获取切片的容量
	// fmt.Println(cap(d))

	// nil
	// var a []int     //声明int类型切片
	// var b = []int{} //声明并且初始化
	// if a == nil {
	// 	fmt.Println("a是一个nil")
	// }
	// fmt.Println(a, len(a), cap(a))
	// fmt.Println(b, len(b), cap(b))

	//切片的赋值拷贝
	// a := make([]int, 3)
	// b := a
	// b[0] = 101
	// fmt.Println(a)
	// fmt.Println(b)

	//切片的遍历
	// a := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(i, a[i])
	// }
	// fmt.Println()
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }

	//切片的扩容
	// var a []int //此时它并没有申请内存
	// // a[0] = 100
	// // fmt.Println(a)
	// a = append(a, 10,2,3,4)
	// fmt.Println(a)
	// b := []int{12,13,14,15}
	// a = append(a, b...)
	// fmt.Println(a)

	//copy
	// a := []int{1, 2, 3, 4, 5}
	// b := make([]int, 5, 5)
	// copy(b, a)
	// b[0] = 100
	// fmt.Println(a)
	// fmt.Println(b)

	//删除元素
	a := []string{"bj", "sh", "gz", "sz"}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)

}
