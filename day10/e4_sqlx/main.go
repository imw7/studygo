package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// 初始化数据库连接
func initDB() (err error) {
	// DSN:Data Source Name
	// dsn := "用户名:密码@tcp(ip:端口)/数据库的名字"
	// 数据库信息
	dsn := "root:password@tcp(127.0.0.1:3306)/sql_test"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect database failed, err:", err)
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

// 查询单条数据
func queryRow(id int) {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据
func queryMultiRow(n int) {
	sqlStr := "select id, name, age from user where id>?"
	var users []user
	err := db.Select(&users, sqlStr, n)
	if err != nil {
		fmt.Println("query failed, err:", err)
		return
	}
	for _, user := range users {
		fmt.Printf("id:%d name:%s age:%d\n", user.ID, user.Name, user.Age)
	}
	// fmt.Printf("users:%#v\n", users)
}

// 插入数据
func insertRow(name string, age int) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Println("insert failed, err:", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Println("get LastInsertId failed, err:", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRow(age int, id int) {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Println("update failed, err:", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("get RowsAffected failed, err:", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRow(id int) {
	sqlStr := "delete from user where id=?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed, err:", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("get RowsAffected failed, err:", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init database failed, err:", err)
		return
	}
	fmt.Println("connect database succeed!")

	queryRow(1)
	queryMultiRow(0)
	insertRow("Jack", 42)
	updateRow(41, 3)
	deleteRow(3)
}
