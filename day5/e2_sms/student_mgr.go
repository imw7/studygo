package main

import "fmt"

type student struct {
	id    int64
	name  string
	age   int8
	score int16
}

// 造一个学生的管理者
type studentMgr struct {
	allStudents map[int64]*student
}

func (s studentMgr) newStudent(id int64, name string, age int8, score int16) *student {
	return &student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}

// 显示所有学生
func (s studentMgr) showAllStudents() {
	for _, stu := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%d\n", stu.id, stu.name, stu.age, stu.score)
	}
}

// 添加学生
func (s studentMgr) addStudent() {
	// 向allStudents中添加一个学生
	var (
		stuID    int64
		stuName  string
		stuAge   int8
		stuScore int16
	)
	fmt.Print("请输入学生的学号：")
	_, _ = fmt.Scanln(&stuID)
	fmt.Print("请输入学生的姓名：")
	_, _ = fmt.Scanln(&stuName)
	fmt.Print("请输入学生的年龄：")
	_, _ = fmt.Scanln(&stuAge)
	fmt.Print("请输入学生的分数：")
	_, _ = fmt.Scanln(&stuScore)

	newStu := s.newStudent(stuID, stuName, stuAge, stuScore)
	s.allStudents[stuID] = newStu
}

// 删除学生
func (s studentMgr) deleteStudent() {
	var deleteID int64

	if len(s.allStudents) == 0 {
		fmt.Println("请先添加学生")
		return
	}
	s.showAllStudents()
	fmt.Print("输入要删除的学生的学号：")
	// 获取用户输入
	_, _ = fmt.Scanln(&deleteID)
	stu, ok := s.allStudents[deleteID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%d\n", stu.id, stu.name, stu.age, stu.score)
	var choice string
	fmt.Print("确定要删除该学生吗？（确定：Y/取消：其他任何键）")
	_, _ = fmt.Scanln(&choice)
	switch choice {
	case "Y":
		// 去allStudents这个map中删除对应id的学生
		delete(s.allStudents, deleteID)
		fmt.Println("删除成功！")
	case "y":
		delete(s.allStudents, deleteID)
		fmt.Println("删除成功！")
	default:
		fmt.Println("取消删除成功！")
	}
}

// 修改学生
func (s studentMgr) updateStudent() {
	// 获取用户输入的学号
	var stuID int64
	if len(s.allStudents) == 0 {
		fmt.Println("请先添加学生")
		return
	}
	s.showAllStudents()
	fmt.Print("请输入要修改的学生的学号：")
	_, _ = fmt.Scanln(&stuID)
	// 展示该学号对应的学生信息，如果没有提示查无此人
	stuObj, ok := s.allStudents[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Println("你要修改的学生信息如下：")
	fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%d\n", stuObj.id, stuObj.name, stuObj.age, stuObj.score)
	// 请输入修改后的学生信息
	var (
		updateName  string
		updateAge   int8
		updateScore int16
	)
	fmt.Print("请输入学生的新名字：")
	_, _ = fmt.Scanln(&updateName)
	fmt.Print("请输入学生的新年龄：")
	_, _ = fmt.Scanln(&updateAge)
	fmt.Print("请输入学生的新分数：")
	_, _ = fmt.Scanln(&updateScore)

	// 更新学生的信息
	stuObj.name = updateName
	stuObj.age = updateAge
	stuObj.score = updateScore
	s.allStudents[stuID] = stuObj // 更新map中的学生
	fmt.Println("修改成功")
}
