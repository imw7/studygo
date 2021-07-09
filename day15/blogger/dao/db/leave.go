package db

import (
	"blogger/model"
	"fmt"
)

// InsertLeave 插入留言
func InsertLeave(leave *model.Leave) (err error) {
	sqlStr := "insert into leave(username, email, content) values (?, ?, ?)"
	_, err = DB.Exec(sqlStr, leave.Username, leave.Email, leave.Content)
	if err != nil {
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlStr, err)
		return
	}
	return
}

// GetLeaveList 获取留言
func GetLeaveList() (leaveList []*model.Leave, err error) {
	sqlStr := "select id, username, email, content, create_time from leave order by id desc"
	err = DB.Select(&leaveList, sqlStr)
	if err != nil {
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlStr, err)
		return
	}
	return
}
