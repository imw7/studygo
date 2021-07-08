package db

import "testing"

func init() {
	// parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	dsn := "root:password@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dsn)
	if err != nil {
		panic(err)
	}
}

// 测试获取单个分类信息
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v\n", category)
}

// 测试获取多个分类信息
func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		t.Errorf("get category list failed, err:%#v\n", err)
		return
	}
	if len(categoryList) != len(categoryIds) {
		t.Errorf("get category list failed, len(categoryList):%d, len(ids):%d\n",
			len(categoryList), len(categoryIds))
	}
	for _, v := range categoryList {
		t.Logf("id:%d category:%#v\n", v.CategoryId, v)
	}
}

// 测试获取所有分类信息
func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		t.Errorf("get all categories failed, err:%v\n", err)
		return
	}
	for _, v := range categoryList {
		t.Logf("category:%#v\n", v)
	}
}
