package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
)

// GetAllCategoryList 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Println("get category list failed, err:", categoryList)
		return
	}
	return
}
