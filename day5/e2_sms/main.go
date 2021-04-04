package main

import (
	"fmt"
	"os"
)

var mgr studentMgr

func showMenu() {
	fmt.Println("--------------欢迎使用学生管理系统--------------")
	fmt.Println(`
		1.显示所有学生
		2.添加学生
		3.删除学生
		4.修改学生
		5.退出
	`)
	fmt.Print("请选择:")
	var choice int
	_, _ = fmt.Scanln(&choice)
	fmt.Printf("选择了%d选项\n", choice)
	switch choice {
	case 1:
		mgr.showAllStudents()
	case 2:
		mgr.addStudent()
	case 3:
		mgr.deleteStudent()
	case 4:
		mgr.updateStudent()
	case 5:
		os.Exit(0)
	default:
		fmt.Println("输入错误，请重新选择")
	}
}

func main() {
	mgr = studentMgr{
		allStudents: make(map[int64]*student, 60),
	}
	for {
		showMenu()
	}
}
