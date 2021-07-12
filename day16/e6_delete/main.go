package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 1.定义模型
type User struct {
	// gorm.Model
	ID     int64
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
	// u1 := User{Name: "phil", Age: 43, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "victoria", Age: 30, Active: false}
	// db.Create(&u2)

	// 5.删除
	// var u = User{}
	// // u.ID = 1
	// u.Name = "phil"
	// db.Debug().Delete(&u)

	db.Debug().Where("name=?", "phil").Delete(&User{})
	db.Debug().Delete(User{}, "age=?", 18)

	// var u []User
	// db.Debug().Unscoped().Where("name = 'phil'").Find(&u)
	// fmt.Println(u)

	// 物理删除
	// db.Debug().Unscoped().Where("name = 'victoria'").Delete(&User{})
}
