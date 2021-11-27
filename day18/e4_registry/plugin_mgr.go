package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// 插件管理器
//  可以用一个大map管理，key字符串，value是Registry接口对象
//  用户自定义去调用，自定义插件
//  实际注册中心的初始化，供系统使用

// PluginMgr 声明管理者结构体
type PluginMgr struct {
	// map维护所有插件
	plugins map[string]Registry
	lock    sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

// RegisterPlugin 插件注册
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

// 注册插件
func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 先去看里面有没有
	_, ok := p.plugins[plugin.Name()]
	if ok {
		err = fmt.Errorf("registry plugin exist")
		return
	}
	p.plugins[plugin.Name()] = plugin
	return
}

// InitRegistry 进行初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegistry(ctx, name, opts...)
}

func (p *PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 查看服务列表，服务是否存在
	plugin, ok := p.plugins[name]
	// 若不存在，就没得初始化，报错
	if !ok {
		err = fmt.Errorf("plugin %s not exist", name)
		return
	}
	// 存在，返回值赋值
	registry = plugin
	// 进行组件初始化
	if err = plugin.Init(ctx, opts...); err != nil {
		log.Fatal(err)
	}
	return
}
