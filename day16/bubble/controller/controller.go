package controller

import (
	"bubble/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
url       --> controller --> logic    -->   model
请求来了   -->  控制器     --> 业务逻辑  -->  模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo model.Todo
	_ = c.BindJSON(&todo)
	// 2. 存入数据库
	if err := model.CreateATodo(&todo); err != nil {
		// 3. 返回响应
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		// c.JSON(http.StatusOK, gin.H{
		// 	"code": 2000,
		// 	"msg":  "success",
		// 	"data": todo,
		// })
	}
}

func GetTodos(c *gin.Context) {
	// 查看todo这个表里的所有数据
	if todos, err := model.GetAllTodos(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "invalid id"})
		return
	}
	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	_ = c.BindJSON(&todo)
	if err = model.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "invalid id"})
		return
	}
	if err := model.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
