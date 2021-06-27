package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// YAML渲染

func main() {
	r := gin.Default()
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})
	_ = r.Run(":8080")
}
