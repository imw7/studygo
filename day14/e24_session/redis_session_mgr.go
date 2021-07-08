package main

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

// RedisSessionMgr设计
// 1.定义RedisSessionMgr对象（字段：存放所有session的map，读写锁）
// 2.构造函数
// 3.Init(): 初始化，加载redis地址
// 4.CreateSession(): 创建一个新的session
// 5.GetSession(): 通过sessionId获取对应的session对象

type RedisSessionMgr struct {
	// redis地址
	addr string
	// 密码
	passwd string
	// 连接池
	pool *redis.Pool
	// 锁
	rwLock sync.RWMutex
	// 大map
	sessionMap map[string]Session
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	// 若有其他参数
	if len(options) > 0 {
		r.passwd = options[0]
	}
	// 创建连接池
	r.pool = myPool(addr, r.passwd)
	r.addr = addr
	return
}

func myPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			// 若有密码，判断
			if _, err := conn.Do("AUTH", password); err != nil {
				_ = conn.Close()
				return nil, err
			}
			return conn, err
		},
		// 连接测试，开发时写
		// 上线注释掉
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// CreateSession 创建session
func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	// go get github.com/satori/go.uuid
	id := uuid.NewV4()
	// 转string
	sessionId := id.String()
	// 创建一个session
	session = NewRedisSession(sessionId, r.pool)
	// 加入到大map
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session does not exist")
		return
	}
	return
}

// NewRedisSessionMgr 构造
func NewRedisSessionMgr() SessionMgr {
	return &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
}
