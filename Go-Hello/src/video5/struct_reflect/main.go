package main

import (
	"fmt"
	"reflect"
)

// 结构体反射
type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func main() {
	stu1 := student{
		Name:  "xwz",
		Score: 90,
	}

	//通过反射获取结构体字段信息
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	//遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		//根据结构体字段的索引去去字段
		fileObj := t.Field(i)
		fmt.Println(fileObj.Name, fileObj.Tag, fileObj.Type)
		fmt.Println(fileObj.Tag.Get("json"), fileObj.Tag.Get("ini"))
	}
	//根据名字取结构体字段
	filedObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("%v %v %v\n", filedObj2.Name, filedObj2.Type, filedObj2.Tag)
	}
}
