package main

import "time"

// 选项设计模式，实现参数初始化

type Options struct {
	// 地址
	Addrs []string
	// 超时时间
	Timeout time.Duration
	// 心跳时间
	Heartbeat int64
	// 注册地址
	//  /a/b/c/xx/10.xxx
	RegistryPath string
}

// Option 定义函数类型的变量
type Option func(opts *Options)

func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithHeartbeat(heartbeat int64) Option {
	return func(opts *Options) {
		opts.Heartbeat = heartbeat
	}
}

func WithRegistryPath(registryPath string) Option {
	return func(opts *Options) {
		opts.RegistryPath = registryPath
	}
}
