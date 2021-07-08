package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
	"math"
)

// GetArticleRecordList 获取文章和对应的分类
func GetArticleRecordList(pageNo, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.GetArticleList(pageNo, pageSize)
	if err != nil {
		fmt.Println("get articleInfo list failed, err:", err)
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2.获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Println("get category list failed, err:", err)
		return
	}
	// 返回页面，做聚合
	// 遍历所有文章
	for _, articleInfo := range articleInfoList {
		// 根据当前文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
		}
		// 文章取出分类id
		categoryId := articleInfo.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// 根据多个文章的id，获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	// 遍历文章，得到每篇文章
	for _, articleInfo := range articleInfoList {
		// 从当前文章取出分类id
		categoryId := articleInfo.CategoryId
		// 去重，防止重复
		for _, id := range ids {
			// 看当前id是否存在
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

// GetArticleRecordListById 根据分类id，获取该类文章和他们对应的分类信息
func GetArticleRecordListById(categoryId, pageNo, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNo, pageSize)
	if err != nil {
		fmt.Println("get article list failed, err:", err)
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2.获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Println("get category list failed, err:", err)
		return
	}
	// 返回页面，做聚合
	// 遍历所有文章
	for _, articleInfo := range articleInfoList {
		// 根据当前文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
		}
		// 文章取出分类id
		categoryId := articleInfo.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// GetArticleDetail 获取文章详情
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(articleId)
	if err != nil {
		return
	}
	category, err := db.GetCategoryById(articleDetail.ArticleInfo.CategoryId)
	if err != nil {
		return
	}
	articleDetail.Category = *category
	return
}

// GetRelativeArticleList 获取相关文章
func GetRelativeArticleList(articleId int64) (articleList []*model.RelativeArticle, err error) {
	articleList, err = db.GetRelativeArticle(articleId)
	return
}

// GetPrevAndNextArticleInfo 获取上下篇文章
func GetPrevAndNextArticleInfo(articleId int64) (prevArticle, nextArticle *model.RelativeArticle, err error) {
	prevArticle, _ = db.GetPrevArticleById(articleId)
	nextArticle, _ = db.GetNextArticleById(articleId)
	return
}

// InsertArticle 插入文章
func InsertArticle(content, author, title string, categoryId int64) (err error) {
	articleDetail := &model.ArticleDetail{}
	articleDetail.Content = content
	articleDetail.Username = author
	articleDetail.Title = title
	articleDetail.ArticleInfo.CategoryId = categoryId
	contentUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	articleDetail.Summary = string([]rune(content)[:minLength])
	id, err := db.InsertArticle(articleDetail)
	fmt.Printf("insert article succ, id:%d, err:%v\n", id, err)
	return
}
