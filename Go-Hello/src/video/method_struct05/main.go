package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性与JSON序列化

type student struct { //首字母大写，对外可见
	ID   int
	Name string
}

func newStudent(id int, name string) student { //student的构造函数
	return student{
		ID:   id,
		Name: name,
	}
}

type class struct {
	Title    string   `json:"title"`
	Students []string `json:"students_list"`
}

func main() {
	c1 := class{ //创建一个班级变量c1
		Title:    "火箭101",
		Students: make([]string, 0, 20),
	}
	for i := 0; i < 10; i++ { //c1中添加学生
		//造十个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
		fmt.Printf("%#v\n", c1)
	}
	// fmt.Printf("%#v\n", c1)
	// JSON序列化
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json序列化出错: ", err)
		return
	}
	// fmt.Println("%T\n", data)
	fmt.Println("%v\n", data)
	//JSON反序列化
	// jsonStr := `{"Title":"火箭101","Students":[{"ID":0,"Name":"stu001"}]}`
	// var c2 class
	// err = json.Unmarshal([]byte(jsonStr), &c2)
	// if err != nil {
	// 	fmt.Println("json反序列化出错：", err)
	// 	return
	// }
	// fmt.Printf("%#v\n", c2)
}
