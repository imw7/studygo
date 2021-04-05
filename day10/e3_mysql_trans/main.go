package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db，是一个连接池对象
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	// dsn := "用户名:密码@tcp(ip:端口)/数据库的名字"
	// 数据库信息
	dsn := "root:password@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	// 不会校验用户名和密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil { // dsn格式不正确的时候会报错
		return
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

// 事务操作
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			_ = tx.Rollback() // 回滚
		}
		fmt.Println("begin trans failed, err:", err)
		return
	}
	sqlStr1 := "update user set age=age+2 where id=?"
	_, err = tx.Exec(sqlStr1, 2)
	if err != nil {
		_ = tx.Rollback() // 回滚
		fmt.Println("exec sql1 failed, err:", err)
		return
	}
	sqlStr2 := "update user set age=age-2 where id=?"
	_, err = tx.Exec(sqlStr2, 4)
	if err != nil {
		_ = tx.Rollback() // 回滚
		fmt.Println("exec sql2 failed, err:", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		_ = tx.Rollback() // 回滚
		fmt.Println("commit failed, err:", err)
		return
	}
	fmt.Println("exec trans success!")
}

func main() {
	err := initDB() // 调用初始化数据库的函数
	if err != nil {
		fmt.Println("init database failed, err:", err)
		return
	}
	fmt.Println("连接数据库成功")

	transactionDemo()
}
