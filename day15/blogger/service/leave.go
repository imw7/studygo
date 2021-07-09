package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
)

func InsertLeave(username, email, content string) (err error) {
	var leave model.Leave
	leave.Content = content
	leave.Email = email
	leave.Username = username
	err = db.InsertLeave(&leave)
	if err != nil {
		fmt.Println("insert leave failed, err:", err)
		return
	}
	return
}

func GetLeaveList() (leaveList []*model.Leave, err error) {
	leaveList, err = db.GetLeaveList()
	if err != nil {
		fmt.Println("get leave list failed, err:", err)
		return
	}
	return
}
