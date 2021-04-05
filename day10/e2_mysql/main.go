package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// Go连接MySQL示例

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

type user struct {
	id   int
	name string
	age  int
}

// 查询单条数据
func queryRow(id int) {
	// 1.写查询单条记录的sql语句
	sqlStr := "select id, name, age from user where id=?"
	var u user
	// 2.执行并拿到结果
	// 从连接池里拿一个连接出来去数据库查询单条记录
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println("scan failed, err:", err)
		return
	}
	// 打印结果
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 查询多条数据
func queryMultiRow(n int) {
	// 1.SQL语句
	sqlStr := "select id, name, age from user where id>?"
	// 2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("query failed, err:", err)
		return
	}
	// 3.非常重要：关闭rows释放持有的数据库链接
	defer func() { _ = rows.Close() }()

	// 4.循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("scan failed, err:", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func insertRow(id int, name string, age int) {
	// 1.写SQL语句
	sqlStr := "insert into user(id, name, age) values (?, ?, ?)"
	// 2.执行插入
	result, err := db.Exec(sqlStr, id, name, age)
	if err != nil {
		fmt.Println("insert failed, err:", err)
		return
	}
	// 如果是插入数据的操作，能够拿到插入数据的id
	theID, err := result.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Println("get LastInsertId failed, err:", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRow(age int, id int) {
	sqlStr := "update user set age=? where id=?"
	result, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Println("update failed, err:", err)
		return
	}
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("get RowsAffected failed, err:", err)
		return
	}
	fmt.Printf("update success, affected rows:%d.\n", n)
}

// 删除数据
func deleteRow(id int) {
	sqlStr := "delete from user where id=?"
	result, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed, err:", err)
		return
	}
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("get RowsAffected failed, err:", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d.\n", n)
}

// 预处理查询
func prepareQuery(n int) {
	sqlStr := "select id, name, age from user where id>?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed, err:", err)
		return
	}
	defer func() { _ = stmt.Close() }()
	rows, err := stmt.Query(n)
	if err != nil {
		fmt.Println("query failed, err:", err)
		return
	}
	defer func() { _ = rows.Close() }()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("scan failed, err:", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入
func prepareInsert() {
	sqlStr := "insert into user(id, name, age) values (?, ?, ?)"
	stmt, err := db.Prepare(sqlStr) // 把SQL语句先发给MySQL预处理一下
	if err != nil {
		fmt.Println("prepare failed, err:", err)
		return
	}
	defer func() { _ = stmt.Close() }()
	// 后续只需要拿到stmt去执行一些操作
	var m = []*user{
		{6, "Tom", 30},
		{7, "Brian", 32},
		{8, "Elizabeth", 21},
		{9, "Bradley", 43},
	}
	for _, u := range m {
		_, err = stmt.Exec(&u.id, &u.name, &u.age) // 后续只需要传值
		if err != nil {
			fmt.Println("insert failed, err:", err)
			return
		}
	}
	fmt.Println("insert success.")
}

func main() {
	err := initDB() // 调用初始化数据库的函数
	if err != nil {
		fmt.Println("init database failed, err:", err)
		return
	}
	fmt.Println("连接数据库成功")

	queryRow(2)
	queryMultiRow(0)
	insertRow(5, "Judy", 33)
	updateRow(32, 5)
	deleteRow(5)
	prepareQuery(2)
	prepareInsert()
}
