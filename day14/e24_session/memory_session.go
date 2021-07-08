package main

import (
	"errors"
	"sync"
)

// MemorySession设计
// 1.定义MemorySession对象（字段：sessionId、存k-v的map、读写锁）
// 2.构造函数，为了获取对象
// 3.Set()
// 4.Get()
// 5.Del()
// 6.Save()

type MemorySession struct {
	sessionId string
	// 存k-v
	data   map[string]interface{}
	rwLock sync.RWMutex
}

// NewMemorySession 构造函数
func NewMemorySession(id string) *MemorySession {
	return &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	// 加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	// 设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key does not exist in session")
		return
	}
	return
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}
