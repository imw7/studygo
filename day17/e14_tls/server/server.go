package main

import (
	"context"
	"google.golang.org/grpc"
	"imw7.com/tls/pb"
	"log"
	"net"
)

// 证书认证

/* 服务端 生成私钥和证书
$ openssl genrsa -out server.key 2048
$ openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" \
    -key server.key -out server.crt
*/

type HelloService struct{}

func (p *HelloService) Hello(ctx context.Context, req *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "Hello, " + req.GetValue() + "!"}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))

	listener, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
