package main

// session接口设计
// 1.Set()
// 2.Get()
// 3.Del()
// 4.Save(): session存储，redis实现延迟加载

type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}
