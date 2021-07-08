package main

import "fmt"

// 中间件让用户选择使用哪个版本

var sessionMgr SessionMgr

func Init(provider, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		_ = fmt.Errorf("不支持")
		return
	}
	err = sessionMgr.Init(addr, options...)
	if err != nil {
		return
	}
	return
}
