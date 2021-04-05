package main

import (
	"fmt"
	"os"
)

var mgr *studentMgr

func showMenu() {
	fmt.Println(`---------------欢迎使用学生管理系统---------------
		1.显示所有学生
		2.添加学生
		3.删除学生
		4.修改学生
		5.查找学生
		6.退出`)
	var choice int
	fmt.Print("请选择:")
	_, _ = fmt.Scanln(&choice)
	fmt.Printf("选择了%d操作\n", choice)
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
		mgr.queryStudent()
	case 6:
		os.Exit(1)
	default:
		fmt.Println("输入错误，请重新选择")
	}
}

func main() {
	err := mgr.initDB()
	if err != nil {
		fmt.Println("init database failed, err:", err)
		return
	}
	fmt.Println("connect database success!")
	for {
		showMenu()
	}
}
