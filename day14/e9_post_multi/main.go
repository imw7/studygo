package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 上传多个文件
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			_ = c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	_ = r.Run(":8000")
}
