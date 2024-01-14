package main

import "fmt"

//空接口

type xxx interface { //没有定义如何方法，可以存储任意变量值；空接口一般不需要提前定义

}

func main() {
	var x interface{} //空接口变量
	x = "hello"
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	x = false
	fmt.Println(x)

	var m = make(map[string]interface{}, 16)
	m["name"] = "娜扎"
	m["age"] = 18
	m["hobby"] = []string{"篮球", "bai球", "红旗"}
	fmt.Println(m)

	ret, ok := x.(string) //类型断言
	if !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Println("是字符串类型", ret)
	}
	switch v := x.(type) {
	//类型断言
	case string:
		fmt.Println("字符串", v)
	case bool:
		fmt.Println("布尔类型", v)
	case int:
		fmt.Println("int类型", v)
	default:
		fmt.Println("猜不到", v)
	}
}
