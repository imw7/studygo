package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM

// Product --> 数据表
type Product struct {
	Code  string
	Price uint
}

func main() {
	// 连接MySQL数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建表 自动迁移（把结构体和数据表进行对应）
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to auto migrate")
	}

	// 创建数据行
	p := Product{Code: "C91", Price: 300}
	db.Create(&p)

	// 查询
	var product Product
	db.First(&product, 1) // 根据整形主键查找
	fmt.Printf("product:%v\n", product)
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Printf("product:%v\n", product)

	// 更新 - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	fmt.Printf("product:%v\n", product)
	// 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	fmt.Printf("product:%v\n", product)
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Printf("product:%v\n", product)

	// 删除 - 删除 product
	db.Delete(&product, 1)
	fmt.Printf("product:%v\n", product)
}
