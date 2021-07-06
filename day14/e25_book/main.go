package main

import (
	"book/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	// 初始化数据库
	err := db.Init()
	if err != nil {
		fmt.Println("init database failed, err:", err)
		return
	}

	r := gin.Default()
	// 加载页面
	r.LoadHTMLGlob("templates/*")
	// 查询所有图书
	r.GET("/book/list", func(c *gin.Context) {
		books, err := db.ShowAllBooks()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg":  err,
			})
			return
		}
		// 返回数据
		c.HTML(http.StatusOK, "book_list.html", gin.H{
			"code": 0,
			"data": books,
		})
	})
	// 添加图书
	r.GET("/book/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new_book.html", nil)
	})
	r.POST("/book/new", func(c *gin.Context) {
		title := c.PostForm("title")
		price := c.PostForm("price")
		p, _ := strconv.Atoi(price)
		err = db.InsertBook(title, int64(p))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg":  err,
			})
			return
		}
		// c.JSON(http.StatusOK, gin.H{
		// 	 "message": "insert book succeed",
		// })
		c.Redirect(http.StatusMovedPermanently, "/book/list")
	})
	// 删除图书
	r.GET("/book/delete", func(c *gin.Context) {
		id := c.Query("id")
		theID, _ := strconv.Atoi(id)
		err = db.DeleteBook(int64(theID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg":  err,
			})
			return
		}
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "delete book succeed",
		// })
		c.Redirect(http.StatusMovedPermanently, "/book/list")
	})
	_ = r.Run(":8080")
}
