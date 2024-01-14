package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统

// 1.添加学员信息
// 2.编辑学员信息
// 3.展示所有学员信息

func showMenu() {
	fmt.Println("欢迎")
	fmt.Println("1.添加")
	fmt.Println("2.编辑")
	fmt.Println("3.展示")
	fmt.Println("4.退出")
}

// 获取用户输入的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("按要求输入学员信息")
	fmt.Println("输入学员的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("输入学员的姓名：")
	fmt.Scanf("%d\n", &name)
	fmt.Println("输入学员的班级：")
	fmt.Scanf("%d\n", &class)
	// 拿到用户输入的学员的所有信息
	stu := newStudent(id, name, class) //调用student的构造函数 造一个学生
	return stu
}

func main() {

	sm := newStudnetMgr()
	for { // 1.打印系统菜单
		showMenu()
		// 2.等待用户选择要执行的选项
		var input int
		fmt.Println("输入要操作的序号：")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入：", input)
		// 3.执行用户选择的动作
		switch input {
		case 1:
			// 添加学员信息
			stu := getInput()
			sm.addStudnet(stu)
		case 2:
			// 编辑学员信息
			stu := getInput()
			sm.modifyStudnet(stu)
		case 3:
			// 展示所有学员信息
			sm.showStudnet()
		case 4:
			// 退出
			os.Exit(0)
		}
	}
}
