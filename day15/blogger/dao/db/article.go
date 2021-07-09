package db

import (
	"blogger/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// InsertArticle 插入文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	// 加个验证
	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}
	sqlStr := `insert into article(content, summary, title, username,
			category_id, view_count, comment_count) values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlStr, article.Content, article.Summary, article.Title,
		article.Username, article.ArticleInfo.CategoryId, article.ViewCount, article.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

// GetArticleList 获取文章列表，做个分页
// pageNo当前是第几页
// pageSize当前页几篇文章
func GetArticleList(pageNo, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNo < 0 || pageSize <= 0 {
		err = fmt.Errorf("invalid parameter, page_no:%d, page_size:%d\n", pageNo, pageSize)
		return
	}
	// 时间降序排序
	sqlStr := `select id, summary, title, view_count, create_time, 
		comment_count, username, category_id from article where status=1 
		order by create_time desc limit ?, ?`
	err = DB.Select(&articleList, sqlStr, pageNo, pageSize)
	return
}

// GetArticleDetail 根据文章id，查询单个文章
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter, article_id:%d\n", articleId)
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlStr := `select id, summary, title, view_count, content, create_time, 
       comment_count, username, category_id from article where id=? and status=1`
	err = DB.Get(articleDetail, sqlStr, articleId)
	return
}

// GetArticleListByCategoryId 根据分类id，查询这一类的文章
func GetArticleListByCategoryId(categoryId, pageNo, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNo < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d page_size:%d\n", pageNo, pageSize)
		return
	}
	sqlStr := `select id, summary, title, view_count, create_time, 
       comment_count, username, category_id from article where 
       status=1 and category_id=? order by create_time desc limit ?, ?`
	err = DB.Select(&articleList, sqlStr, categoryId, pageNo, pageSize)
	return
}

// GetRelativeArticle 获取相关文章
func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	var categoryId int64
	sqlStr := "select category_id from article where id=?"
	err = DB.Get(&categoryId, sqlStr, articleId)
	if err != nil {
		return
	}
	sqlStr = "select id, title from article where category_id=? and id!=? limit 10"
	err = DB.Select(&articleList, sqlStr, categoryId, articleId)
	return
}

// GetPrevArticleById 获取上一篇文章
func GetPrevArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlStr := "select id, title from article where id<? order by id desc limit 1"
	err = DB.Get(info, sqlStr, articleId)
	if err != nil {
		return
	}
	return
}

// GetNextArticleById 获取下一篇文章
func GetNextArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlStr := "select id, title from article where id<? order by id asc limit 1"
	err = DB.Get(info, sqlStr, articleId)
	if err != nil {
		return
	}
	return
}

// IsArticleExist 判断文章是否存在
func IsArticleExist(articleId int64) (exists bool, err error) {
	var id int64
	sqlStr := "select id from article where id=?"
	err = DB.Get(&id, sqlStr, articleId)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	exists = true
	return
}
