package controller

import (
	"blogger/service"
	"blogger/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"path"
	"path/filepath"
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

// NewArticleHandler 新建文章
func NewArticleHandler(c *gin.Context) {
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Println("get article failed, err:", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/post_article.html", categoryList)
}

// ArticleSubmitHandler 发表文章
func ArticleSubmitHandler(c *gin.Context) {
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryIdStr := c.PostForm("category_id")
	title := c.PostForm("title")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	err = service.InsertArticle(content, author, title, categoryId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

// ArticleDetailHandler 文章详细页
func ArticleDetailHandler(c *gin.Context) {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleDetail, err := service.GetArticleDetail(articleId)
	if err != nil {
		fmt.Printf("get article detail failed, article_id:%d err:%v\n", articleId, err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	relativeArticle, err := service.GetRelativeArticleList(articleId)
	if err != nil {
		fmt.Println("get relative article failed, err:", err)
	}
	prevArticle, nextArticle, err := service.GetPrevAndNextArticleInfo(articleId)
	if err != nil {
		fmt.Println("get prev or next article failed, err:\n", err)
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Println("get all categories failed, err:", err)
	}
	// 获取评论列表
	commentList, err := service.GetCommentList(articleId)
	if err != nil {
		fmt.Println("get comment list failed, err:", err)
	}
	var m = make(map[string]interface{}, 10)
	m["detail"] = articleDetail
	m["relative_article"] = relativeArticle
	m["prev"] = prevArticle
	m["next"] = nextArticle
	m["category"] = allCategoryList
	m["article_id"] = articleId
	m["comment_list"] = commentList
	c.HTML(http.StatusOK, "views/detail.html", m)
}

// UploadFileHandler 上传文件
func UploadFileHandler(c *gin.Context) {
	// 单个文件
	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	rootPath := utils.GetRootDir()
	id := uuid.NewV4()
	ext := path.Ext(file.Filename)
	url := fmt.Sprintf("/static/upload/%s%s", id, ext)
	dst := filepath.Join(rootPath, url)
	// 上传文件到指定路径
	_ = c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"uploaded": true,
		"usl":      url,
	})
}

// NewLeaveHandler 新留言
func NewLeaveHandler(c *gin.Context) {
	leaveList, err := service.GetLeaveList()
	if err != nil {
		fmt.Println("get leave failed, err:", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/gbook.html", leaveList)
}

// AboutMeHandler 关于我页面
func AboutMeHandler(c *gin.Context) {
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryIdStr := c.PostForm("category_id")
	title := c.PostForm("title")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	err = service.InsertArticle(content, author, title, categoryId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
