package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dsn := "root:password@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := db.Init(dsn)
	if err != nil {
		panic(err)
	}
	ginpprof.Wrapper(router)
	// 加载静态文件
	router.Static("/static/", "./static")
	// 加载模板
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandler)
	router.GET("/category/", controller.CategoryListHandler)
	router.GET("/article/new/", controller.NewArticleHandler)
	router.POST("/article/submit/", controller.ArticleSubmitHandler)
	router.GET("/article/detail/", controller.ArticleDetailHandler)
	router.POST("/upload/file/", controller.UploadFileHandler)
	router.GET("/leave/new/", controller.NewLeaveHandler)
	_ = router.Run(":8080")
}
