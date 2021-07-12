package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM 创建

// User
// 1. 定义模型
type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:'john'"` // 定义默认值
	Age  sql.NullInt64  `gorm:"default:18"`
}

func main() {
	// 连接MySQL数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8&mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 2. 把模型与数据库中的表对应起来
	_ = db.AutoMigrate(&User{})

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

	// 批量插入
	users := []User{u, u1, u2, u3}
	db.Create(&users)
	for _, user := range users {
		fmt.Println(user.ID)
	}

	// 根据Map创建
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "jessie", "Age": 21,
	})
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "eric", "Age": 19},
		{"Name": "sarah", "Age": 32},
	})
}
