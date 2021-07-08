package db

import (
	"blogger/model"
	"github.com/jmoiron/sqlx"
)

// InsertCategory 添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "insert into category(category_name, category_no) values (?, ?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

// GetCategoryById 获取单个分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	sqlStr := "select id, category_name, category_no from category where id=?"
	category = &model.Category{}
	// 查询单行数据
	err = DB.Get(category, sqlStr, id)
	return
}

// GetCategoryList 获取多个分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlStr, args, err := sqlx.In( // 使用sqlx.In拼接语句和参数
		"select id, category_name, category_no from category where id in (?)",
		categoryIds,
	)
	if err != nil {
		return
	}
	// 查询多行数据
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

// GetAllCategoryList 获取全部分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlStr)
	return
}
