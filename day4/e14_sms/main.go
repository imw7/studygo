package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看/新增/删除学生
*/
var allStudents map[int64]*student // 变量声明

type student struct {
	id   int64
	name string
}

// newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudents() {
	// 把所有的学生都打印出来
	for k, v := range allStudents {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	// 向allStudents中添加一个新学生
	// 1.创建一个新学生
	// 1.1获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号：")
	_, _ = fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名：")
	_, _ = fmt.Scanln(&name)
	// 1.2 造学生（调用student的构造函数）
	newStu := newStudent(id, name)

	// 2.追加到allStudents这个map中
	allStudents[id] = newStu
}

func deleteStudent() {
	// 1.请用户输入要删除的学生的学号
	var deleteID int64
	fmt.Print("请输入学生的学号：")
	_, _ = fmt.Scanln(&deleteID)
	// 2.去allStudents这个map中根据学号删除对应的键值对
	delete(allStudents, deleteID)
}

func main() {
	allStudents = make(map[int64]*student, 50) // 初始化（开辟内存空间）
	for {
		// 1.打印菜单
		fmt.Println(`--------------欢迎使用学生管理系统--------------
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出`)
		fmt.Print("请选择操作: ")
		// 2.等待用户选择要做什么
		var choice int
		_, _ = fmt.Scanln(&choice)
		fmt.Printf("选择了%d操作\n", choice)
		// 3.执行对应的函数
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出
		default:
			fmt.Println("选择错误，请重新选择")
		}
	}
}
