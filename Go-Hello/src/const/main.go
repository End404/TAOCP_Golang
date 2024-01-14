package main

import "fmt"

/* 第2章
2.2 变量和常量
*/

// 常量
// const pi = 3.1415
// const e = 2.7

// const (
// 	pi = 3.1415
// 	e = 2.7
// )

const (
	n1 = 10
	n2
	n3
)

// const (
// 	nn1 = iota
// 	nn2
// 	nn3
// 	nn4
// )

const (
	nn1 = iota // 0
	nn2 = iota // 1
	nn3 = 100  // 100
	nn4 = iota // 3
)

const nn5 = iota // 0

const (
	_  = iota
	KB = 1 << (10 * iota) // 1<<10
	MB = 1 << (10 * iota) // 1<<20
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 //iota=0, 1,2
	c, d
	e, f
)

func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(nn1, nn2, nn3, nn4, nn5)
	fmt.Println(a, b, c, d, e, f)
}
