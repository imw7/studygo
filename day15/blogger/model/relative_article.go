package model

// RelativeArticle 相关文章
type RelativeArticle struct {
	ArticleId int64  `db:"id"`
	Title     string `db:"title"`
}
