package db

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 通过RPC构造一个简单的内存KV数据库

type KVStoreService struct {
	m      map[string]string           // 用于存储KV数据
	filter map[string]func(key string) // 对应每个Watch调用定义的过滤器函数列表
	mu     sync.Mutex                  // 互斥锁，用于在多个Goroutine访问或修改时对其它成员提供保护
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

// Set 输入参数是key和value组成的数组，用一个匿名的空结构体表示忽略了输出参数
// 当修改某个key对应的值时会调用每一个过滤器参数
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

// Watch 提供过滤器列表
// 输入参数是超时的秒数，当有key变化时将key作为返回值返回
// 如果超过时间后依然没有key被修改，则返回超时错误
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int()) // 用唯一的id表示每个Watch调用
	ch := make(chan string, 10)                                // buffered

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key } // 根据id将对应的过滤器函数注册到p.filter列表
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}
