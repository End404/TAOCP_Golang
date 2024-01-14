package main

import "fmt"

//为什么需要接口

type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

type sayer interface {
	say()
}

func da(arg sayer) { //打的函数
	arg.say() //打它 就会叫
}
func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
}
