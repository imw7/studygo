package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// Go连接MySQL示例

func main() {
	// DSN:Data Source Name
	// dsn := "用户名:密码@tcp(ip:端口)/数据库的名字"
	// 数据库信息
	dsn := "root:password@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	// 不会校验用户名和密码是否正确
	db, err := sql.Open("mysql", dsn)
	if err != nil { // dsn格式不正确的时候会报错
		fmt.Printf("dsn:%s invalid, err:%v\n", dsn, err)
		return
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}
	defer func() { _ = db.Close() }() // 注意这行代码要写在上面err判断的下面
	fmt.Println("connect database succeed!")
}
