package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
// 创建数据库
CREATE DATABASE students;

// 进入数据库
use students;

// 创建数据表
CREATE TABLE `student` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
*/

type student struct {
	id   int
	name string
	age  int
}

type studentMgr struct {
}

var db *sql.DB

// initDB 初始化数据库连接
func (s *studentMgr) initDB() (err error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/students"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

// queryRow 根据id查找student表单
func (s *studentMgr) queryRow(id int) *student {
	sqlStr := "select id, name, age from student where id=?"
	var stu student
	err := db.QueryRow(sqlStr, id).Scan(&stu.id, &stu.name, &stu.age)
	if err != nil {
		return nil
	}
	fmt.Printf("学号:%d 姓名:%s 年龄:%d\n", stu.id, stu.name, stu.age)
	return &stu
}

// showAllStudents 显示所有学生
func (s *studentMgr) showAllStudents() {
	sqlStr := "select id, name, age from student where id>0"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("query failed, err:", err)
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var stu student
		err := rows.Scan(&stu.id, &stu.name, &stu.age)
		if err != nil {
			fmt.Println("scan failed, err:", err)
			return
		}
		fmt.Printf("学号:%d 姓名:%s 年龄:%d\n", stu.id, stu.name, stu.age)
	}
}

// addStudent 添加学生
func (s *studentMgr) addStudent() {
	var (
		id   int
		name string
		age  int
	)
	fmt.Print("输入学号:")
	_, _ = fmt.Scanln(&id)
	fmt.Print("输入姓名:")
	_, _ = fmt.Scanln(&name)
	fmt.Print("输入年龄:")
	_, _ = fmt.Scanln(&age)
	sqlStr := "insert into student(id, name, age) values (?, ?, ?)"
	_, err := db.Exec(sqlStr, id, name, age)
	if err != nil {
		fmt.Println("insert failed, err:", err)
		return
	}
	fmt.Println("添加成功")
}

// deleteStudent 删除学生
func (s *studentMgr) deleteStudent() {
	var id int
	fmt.Print("输入要删除的学生学号:")
	_, _ = fmt.Scanln(&id)
	if s.queryRow(id) == nil {
		fmt.Println("查无此人")
		return
	}
	sqlStr := "delete from student where id=?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed, err:", err)
		return
	}
	fmt.Println("删除成功")
}

// updateStudent 更新学生
func (s *studentMgr) updateStudent() {
	var id int
	fmt.Print("输入要修改的学生学号:")
	_, _ = fmt.Scanln(&id)
	if s.queryRow(id) == nil {
		fmt.Println("查无此人")
		return
	}
	var (
		name string
		age  int
	)
	fmt.Print("输入新姓名:")
	_, _ = fmt.Scanln(&name)
	fmt.Print("输入新年龄:")
	_, _ = fmt.Scanln(&age)
	sqlStr := "update student set name=?, age=? where id=?"
	_, err := db.Exec(sqlStr, name, age, id)
	if err != nil {
		fmt.Println("update failed, err:", err)
		return
	}
	fmt.Println("更新成功")
}

// queryStudent 查找学生
func (s *studentMgr) queryStudent() {
	var id int
	fmt.Print("输入要查找的学生学号:")
	_, _ = fmt.Scanln(&id)
	if s.queryRow(id) == nil {
		fmt.Println("查无此人")
		return
	}
}
