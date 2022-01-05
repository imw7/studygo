package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

// User 用户信息
type User struct {
	Id   int
	Name string `orm:"unique"` // 用户名唯一
	Pwd  string
}

func init() {
	// 设置数据库基本信息
	if err := orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/my_test?charset=utf8"); err != nil {
		logs.Error(err)
	}
	// 映射model数据库
	orm.RegisterModel(&User{})
	// 生成表
	if err := orm.RunSyncdb("default", false, true); err != nil {
		logs.Error(err)
	}
}
