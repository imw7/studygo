package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"imw7.com/interceptor/pb"
	"log"
	"net"
)

type HelloService struct{}

func (h *HelloService) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	reply := &pb.HelloReply{Message: "Hello, " + request.GetName() + "!"}
	return reply, nil
}

// filter 要实现普通方法的截取器，需要实现的函数
// 函数的ctx和req参数就是每个普通的RPC方法的前两个参数
// 第三个info参数表示当前是对应的那个gRPC方法
// 第四个handler参数对应当前的gRPC方法函数
func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter:", info)
	defer func() { // 添加对gRPC方法异常的捕获
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func main() {
	// 将filter函数作为参数传入
	// 服务器在收到每个gRPC方法调用之前，会首先输出一行日志，然后再调用对方的方法
	// 2020/07/20 23:33:54 filter: &{0xc12cd8 /pb.HelloService/Hello}
	// 注意：gRPC框架中只能为每个服务设置一个截取器，因此所有的截取工作只能在一个函数中完成
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))
	pb.RegisterHelloServiceServer(server, new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
