package main

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// MemorySessionMgr设计
// 1.定义MemorySessionMgr对象（字段：存放所有session的map，读写锁）
// 2.构造函数
// 3.Init(): 初始化，加载redis地址
// 4.CreateSession(): 创建一个新的session
// 5.GetSession(): 通过sessionId获取对应的session对象

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

// NewMemorySessionMgr 构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

// CreateSession 创建session
func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	// go get github.com/satori/go.uuid
	id := uuid.NewV4()
	// 转string
	sessionId := id.String()
	// 创建一个session
	session = NewMemorySession(sessionId)
	// 加入到大map
	s.sessionMap[sessionId] = session
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = errors.New("session does not exist")
		return
	}
	return
}
