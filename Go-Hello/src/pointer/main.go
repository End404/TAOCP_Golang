package main

import "fmt"

// 指针

/*
func main() {
	a := 10 //int类型
	b := &a //*int类型(int指针)
	// fmt.Println(a, &a)
	// fmt.Println(b)
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", b, b)
	c := *b
	fmt.Println(c)
}
*/

func modify1(x int) {
	x = 100
}
func modify2(y *int) {
	*y = 100
}
func main() {
	a := 1
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}
