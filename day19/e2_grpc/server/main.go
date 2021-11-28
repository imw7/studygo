package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-user/pb"
	"log"
	"net"
)

type UserInfoService struct{}

func (u *UserInfoService) GetUserInfo(_ context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	// 通过用户名查询用户信息
	name := request.Name
	if name == "Eric" {
		return &pb.UserResponse{
			Id:      10086,
			Name:    name,
			Age:     21,
			Hobbies: []string{"movie", "book", "sing"},
		}, nil
	}
	return nil, nil
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 实例化grpc
	grpcServer := grpc.NewServer()
	// 在grpc上注册微服务
	pb.RegisterUserInfoServiceServer(grpcServer, new(UserInfoService))

	// 启动服务端
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
