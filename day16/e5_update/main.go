package main

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 更新

// User 1.定义模型
type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	// 2.连接MySQL数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8&mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 3.把模型与数据库中的表对应起来
	_ = db.AutoMigrate(&User{})

	// 4.创建
	// u1 := User{Name: "jessie", Age: 18, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "edward", Age: 21, Active: false}
	// db.Create(&u2)

	// 5.查询
	var user User
	db.First(&user)

	// 6.更新
	user.Name = "婕西"
	user.Age = 99
	db.Debug().Save(&user) // 默认会修改所有字段

	// db.Debug().Model(&user).Update("name", "张三")

	// m1 := map[string]interface{}{
	// 	"name":   "jack",
	// 	"age":    25,
	// 	"active": true,
	// }
	// db.Debug().Model(&user).Updates(m1) // m1列出的所有字段都会更新
	// db.Debug().Model(&user).Select("age").Updates(m1) // 只更新age字段
	// db.Debug().Model(&user).Omit("active").Updates(m1) // 排除m1中的active更新其他的字段

	// db.Debug().Model(&user).UpdateColumn("age", 30)
	// rowsNum := db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected
	// fmt.Println(rowsNum)

	// 让users表中所有的用户的年龄在原来的基础上+2
	db.Model(&user).Update("age", gorm.Expr("age+?", 2))
}
