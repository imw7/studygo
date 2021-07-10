package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 处理multipart forms提交文件时默认的内存限制是 32 MiB
	// 可以通过下面的方式修改
	// r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// r.POST("/upload", func(c *gin.Context) {
	// 	// 单个文件
	// 	f, err := c.FormFile("file") // 从请求中获取携带的参数
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	// 将读取到的文件保存在本地（服务端本地）
	// 	// dst := fmt.Sprintf("./%s", f.Filename)
	// 	dst := path.Join("./", f.Filename)
	// 	err = c.SaveUploadedFile(f, dst)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "ok",
	// 	})
	// })

	r.POST("/upload", func(c *gin.Context) {
		// 多个文件
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%d_%s\n", index, file.Filename)
			// 上传文件到指定的目录
			err = c.SaveUploadedFile(file, dst)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	_ = r.Run(":8080")
}
