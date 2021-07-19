package main

import (
	"context"
	"github.com/moby/moby/pkg/pubsub" // docker项目实现发布订阅模式的包
	"google.golang.org/grpc"
	"imw7.com/pubsub/pb"
	"log"
	"net"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

// Publish 发布方法
func (p *PubsubService) Publish(ctx context.Context, arg *pb.String) (*pb.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pb.String{}, nil
}

// Subscribe 订阅方法
func (p *PubsubService) Subscribe(arg *pb.String, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&pb.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// 构建一个gRPC服务
	grpcServer := grpc.NewServer()
	// 通过gRPC插件生成的Register函数注册自己实现的PubsubService服务
	pb.RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	// 在一个监听端口上提供gRPC服务
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
