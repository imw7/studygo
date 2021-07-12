package model

import (
	"bubble/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/

// CreateATodo 创建todo
func CreateATodo(todo *Todo) error {
	return dao.DB.Create(&todo).Error
}

// GetAllTodos 获取所有的todo
func GetAllTodos() (todos []*Todo, err error) {
	if err = dao.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return
}

// GetATodo 获取一个todo
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateATodo 更新一个todo
func UpdateATodo(todo *Todo) error {
	return dao.DB.Save(todo).Error
}

// DeleteATodo 删除一个Todo
func DeleteATodo(id string) error {
	return dao.DB.Where("id=?", id).Delete(&Todo{}).Error
}
