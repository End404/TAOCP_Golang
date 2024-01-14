package main

import "fmt"

// 结构体构造函数

type person struct {
	name, city string
	age        int8
}

func newPerson(name, city string, age int8) *person { //构造函数
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func main() {
	p1 := newPerson("洗完澡", "沙河", int8(18))
	fmt.Printf("type:%T value:%#v\n", p1, p1)
}
