package db

import (
	"blogger/model"
	"fmt"
)

// InsertComment 插入评论
func InsertComment(comment *model.Comment) (err error) {
	if comment == nil {
		err = fmt.Errorf("invalid comment parameter")
		return
	}
	tx, err := DB.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
	}()
	sqlStr := `insert into comment(content, username, article_id) values (?, ?, ?)`
	_, err = tx.Exec(sqlStr, comment.Content, comment.Username, comment.ArticleId)
	if err != nil {
		return
	}
	sqlStr = `update article set comment_count=comment_count+1 where id=?`
	_, err = tx.Exec(sqlStr, comment.ArticleId)
	if err != nil {
		return
	}
	err = tx.Commit()
	return
}

// UpdateViewCount 更新评论次数
func UpdateViewCount(articleId int64) (err error) {
	sqlStr := `update article set comment_count=comment_count+1 where id=?`
	_, err = DB.Exec(sqlStr, articleId)
	if err != nil {
		return
	}
	return
}

// GetCommentList 获取评论列表
func GetCommentList(articleId int64, pageNo, pageSize int) (commentList []*model.Comment, err error) {
	if pageNo < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, pageNo:%d, pageSize:%d\n", pageNo, pageSize)
		return
	}
	sqlStr := `select id, content, username, create_time from comment 
		where article_id=? and status=1 order by create_time desc limit ?, ?`
	err = DB.Select(&commentList, sqlStr, articleId, pageNo, pageSize)
	return
}
