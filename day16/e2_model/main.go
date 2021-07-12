package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// User 定义模型
type User struct {
	gorm.Model   // 内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// Animal 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64  `gorm:"primaryKey" gorm:"column:beast_id"`
	Name     string `gorm:"column:name_of_the_beast"`
	Age      int64  `gorm:"column:age_of_the_beast"` // 使用tag指定列名
}

// TableName 将 Animal 的表名设置为 `beasts`
// 唯一指定表名
func (Animal) TableName() string {
	return "beasts"
}

func main() {
	// 连接MySQL数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "prefix_", // 更改命名规则
			SingularTable: true,      // 禁用复数
		},
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&User{})
	_ = db.AutoMigrate(&Animal{})

	// 使用User结构体创建名为`deleted_users`的表
	_ = db.Table("deleted_users").AutoMigrate(&User{})
}
