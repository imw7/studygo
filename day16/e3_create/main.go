package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GORM CRUD

// User
// 1. 定义模型
type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:'john'"` // 定义默认值
	Age  sql.NullInt64  `gorm:"default:18"`
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db1?charset=utf8&mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer func() { _ = db.Close() }()
	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3. 创建

	// 张三 18
	u := User{ // 在代码层面创建一个User对象
		Name: sql.NullString{String: "张三", Valid: true}, // Valid 为 true 表示不为空
		Age:  sql.NullInt64{Int64: 0, Valid: false},     // Age为空值，创建默认年龄
	}

	// john 22
	u1 := User{
		Name: sql.NullString{String: "", Valid: false}, // 创建默认的名字
		Age:  sql.NullInt64{Int64: 22, Valid: true},
	}

	// "" 22
	u2 := User{
		Name: sql.NullString{String: "", Valid: true}, // 创建空字符串的名字
		Age:  sql.NullInt64{Int64: 22, Valid: true},
	}

	// john 18
	u3 := User{
		Name: sql.NullString{String: "", Valid: false}, // 默认姓名
		Age:  sql.NullInt64{Int64: 0, Valid: false},    // 默认年龄
	}

	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空 true
	db.Debug().Create(&u)         // 在数据库中创建了一条`张三 18`的记录
	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空 false

	db.Debug().Create(&u1) // 在数据库中创建了一条`john 22`的记录
	db.Debug().Create(&u2) // 在数据库中创建了一条`"" 22`的记录
	db.Debug().Create(&u3) // 在数据库中创建了一条`john 18`的记录
}
