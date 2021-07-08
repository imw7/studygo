package db

import (
	"blogger/model"
	"testing"
	"time"
)

func init() {
	// parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	dsn := "root:password@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dsn)
	if err != nil {
		panic(err)
	}
}

// 测试插入文章
func TestInsertArticle(t *testing.T) {
	// 构建对象
	article := &model.ArticleDetail{
		ArticleInfo: model.ArticleInfo{
			CategoryId:   1,
			CommentCount: 0,
			CreateTime:   time.Now(),
			Summary:      "唐代诗人李白的作品",
			Title:        "静夜思",
			Username:     "李白",
			ViewCount:    1,
		},
		Content: `窗前明月光，
		 疑似地上霜。
         举头望明月，
         低头思故乡。`,
		Category: model.Category{CategoryId: 1},
	}
	articleId, err := InsertArticle(article)
	if err != nil {
		t.Errorf("insert article failed, err:%v\n", err)
		return
	}
	t.Logf("insert article succ, articleId:%d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		t.Errorf("get article list failed, err:%v\n", err)
		return
	}
	t.Logf("get article list succ, len:%d\n", len(articleList))
}

func TestGetRelativeArticle(t *testing.T) {
	articleList, err := GetRelativeArticle(3)
	if err != nil {
		t.Errorf("get relative article failed, err:%v\n", err)
		return
	}
	for _, v := range articleList {
		t.Logf("id:%d title:%s\n", v.ArticleId, v.Title)
	}
}

func TestGetPrevArticleById(t *testing.T) {
	articleInfo, err := GetPrevArticleById(5)
	if err != nil {
		t.Errorf("get prev article failed, err:%v\n", err)
		return
	}
	t.Logf("article info:%#v\n", articleInfo)
}

func TestGetNextArticleById(t *testing.T) {
	articleInfo, err := GetNextArticleById(5)
	if err != nil {
		t.Errorf("get prev article failed, err:%v\n", err)
		return
	}
	t.Logf("article info:%#v\n", articleInfo)
}
