package calc

import (
	"Go-Hello/src/package_demo/snow"
	"fmt"
)

var Name = "shxwz" //定义一个测试的全局变量

func Add(x, y int) int { //计算两个int类型数据和的函数
	snow.Snow()
	// Sub(x, y) //同一个包中的不同文件可以直接调用
	return x + y
}

func init() { //init函数导入时自动执行 无参数无返回值
	fmt.Println("calc.init()")
	fmt.Println(Name)
}
