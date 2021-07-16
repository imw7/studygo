package model

import "net/rpc"

// HelloServiceName 服务的名字
const HelloServiceName = "path/to/pkg.HelloService" // 避免名字冲突

// HelloServiceInterface 服务要实现的详细方法列表
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// HelloServiceClient 该类型必须满足HelloServiceInterface接口
// 这样客户端用户就可以直接通过接口对应的方法调用RPC函数
type HelloServiceClient struct {
	*rpc.Client
}

// _ 用作类型断言，判断是否实现了接口定义的所有方法，如果没有会报错
// 确保HelloServiceClient结构体类型实现HelloServiceInterface接口的所有方法
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

// DialHelloService 直接拨号HelloService服务
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

// RegisterHelloService 注册该类型服务的函数
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
