package main

import "fmt"

//使用值接受者实现接口 和 使用指针接受者实现接口 的 区别

type animal interface { //接口的嵌套
	mover
	sayer
}
type mover interface {
	move()
}
type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

// func (p person) move() { //值接受者实现接口
//
//		fmt.Printf("%s在跑...", p.name)
//	}
func (p *person) move() { //指针接受者实现接口
	fmt.Printf("%s在跑...", p.name)
}
func (p *person) say() {
	fmt.Printf("%s在叫。。。", p.name)
}

func main() {
	var m mover //定义一个mover类型的变量
	//p1 := person{ //person类型的值
	//	name: "小王子",
	//	age:  18,
	//}
	p2 := &person{ //person类型的指针
		name: "娜扎",
		age:  18,
	}
	//m = p1
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer //定义一个sayer类型的变量
	s = p2
	s.say()
	fmt.Println(s)

	var a animal //animal类型的变量
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}
