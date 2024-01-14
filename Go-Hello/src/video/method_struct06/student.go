package main

import "fmt"

type student struct {
	id    int //学号
	name  string
	class string
}

func newStudent(id int, name, class string) *student { //newStudent 是student类型的构造函数
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct { //学员管理的类型
	allStudents []*student
}

func newStudnetMgr() *studentMgr { //构造函数
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 1.添加学员信息
func (s *studentMgr) addStudnet(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2.编辑学员信息
func (s *studentMgr) modifyStudnet(newStu *student) {
	// s.allStudents = append(s.allStudents, newStu)
	for i, v := range s.allStudents {
		if newStu.id == v.id { //学号相同时，表示找到要修改的学生
			s.allStudents[i] = newStu //根据切片的索引直接把新学生赋值进来
			return
		}
	}
	// 这里说明输入的学生没有找到
	fmt.Println("输入学生信息有误，没有学号是：%d 的学生\n", newStu.id)
}

// 3.展示所有学员信息
func (s *studentMgr) showStudnet() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)

	}
}
