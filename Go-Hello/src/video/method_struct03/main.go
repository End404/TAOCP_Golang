package main

import "fmt"

//嵌套结构体的字段冲突

type Address struct {
	Province   string
	City       string
	UpdateTime string
}
type Email struct {
	Addr       string
	UpdateTime string
}
type Person struct {
	Name    string
	Gender  string
	Age     int
	Address //	嵌套另外一个结构体
	Email   //嵌套另一个结构体
}

func main() {
	p1 := Person{
		Name:   "xwz",
		Gender: "nan",
		Age:    18,
		Address: Address{
			Province:   "sd",
			City:       "wh",
			UpdateTime: "2023-08-21",
		},
		Email: Email{
			Addr:       "xiaowangzi@xiaowangzi.com",
			UpdateTime: "2022-08-20",
		},
	}
	fmt.Printf("%#v\n", p1)
	// fmt.Println(p1.UpdateTime)	//多个同名的字段
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.UpdateTime)
}
