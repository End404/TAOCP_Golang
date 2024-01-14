package main

import "fmt"

// 方法的定义

type Person struct { //结构体
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person { //Person类型的构造函数
	return &Person{
		name: name,
		age:  age,
	}
}

func (p Person) Dream() { //为Person类型定义方法
	fmt.Printf("%s 学好Go语言\n", p.name)
}

func (p *Person) SetAge(newAge int8) { //修改年龄的方法 指针接受者
	p.age = newAge
}
func (p Person) SetAge2(newAge int8) { //修改年龄的方法 使用值接受者
	p.age = newAge
}

func main() {
	p1 := NewPerson("娜扎", int8(22))
	// (*p1).Dream()
	// p1.Dream()

	//调用修改年龄的方法
	fmt.Println(p1.age)
	p1.SetAge(int8(20))
	fmt.Println(p1.age)
	p1.SetAge2(int8(30))
	fmt.Println(p1.age)
}
