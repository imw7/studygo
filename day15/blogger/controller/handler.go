package controller

import (
	"blogger/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// IndexHandler 访问主页的控制器
func IndexHandler(c *gin.Context) {
	// 从service取数据
	// 1.加载文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// gin.H本质上是一个map
	// var data map[string]interface{} = make(map[string]interface{}, 16)
	// data["article_list"] = articleRecordList
	// data["category_list"] = categoryList
	// c.HTML(http.StatusOK, "views/index.html", data)

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

// CategoryListHandler 点击分类云，展示具体的分类的内容
func CategoryListHandler(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 根据分类id，获取文章列表
	articleRecordList, err := service.GetArticleRecordListById(int(categoryId), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 再次加载所有分类数据，用于分类云显示
	categoryList, err := service.GetAllCategoryList()
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}
