package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
	"time"
)

// InsertComment 插入评论
func InsertComment(comment, author, email string, articleId int64) (err error) {
	// 判断文章是否存在
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Println("query database failed, err:", err)
		return
	}
	if exist == false {
		err = fmt.Errorf("article id:%d not found\n", articleId)
		return
	}
	var c model.Comment
	c.ArticleId = articleId
	c.Content = comment
	c.Username = author
	c.CreateTime = time.Now()
	c.Status = 1
	err = db.InsertComment(&c)
	return
}

// GetCommentList 获取评论
func GetCommentList(articleId int64) (commentList []*model.Comment, err error) {
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Println("query database failed, err:", err)
		return
	}
	if exist == false {
		err = fmt.Errorf("article id:%d not found\n", articleId)
		return
	}
	commentList, err = db.GetCommentList(articleId, 0, 100)
	return
}
