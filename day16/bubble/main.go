package main

import (
	"bubble/dao"
	"bubble/model"
	"bubble/router"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	_ = dao.DB.AutoMigrate(&model.Todo{})

	r := router.SetupRouter()

	_ = r.Run(":8080")
}
