package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sql注入

var db *sqlx.DB

// 初始化数据库连接
func initDB() (err error) {
	// DSN:Data Source Name
	// dsn := "用户名:密码@tcp(ip:端口)/数据库的名字"
	// 数据库信息
	dsn := "root:password@tcp(127.0.0.1:3306)/test"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect DB failed, err:", err)
		return
	}
	db.SetMaxOpenConns(20) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10) // 设置最大空闲连接数
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

func sqlInject(name string) {
	// 自己拼接sql语句的字符串
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Println("SQL:", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Println("exec failed, err:", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init DB failed, err:", err)
		return
	}
	fmt.Println("连接数据库成功")

	// sql注入的几种示例
	sqlInject("Eric")
	sqlInject("xxx' or 1=1#")
	sqlInject("xxx' union select * from user #")
	sqlInject("xxx' and (select count(*) from user) <10 #")
}
