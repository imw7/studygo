package main

// 定义管理者，管理所有session
// SessionMgr接口设计
// 1.Init(): 初始化，加载redis地址
// 2.CreateSession(): 创建一个新的session
// 3.GetSession(): 通过sessionId获取对应的session对象

type SessionMgr interface {
	// Init 初始化
	Init(addr string, options ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
}
