package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"imw7.com/grpc_rest/pb"
	"log"
	"net/http"
)

// REST接口

func main() {
	// 启动Web服务
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux() // 创建路由处理器

	// 将RestService服务相关的REST接口中转到后面的gRPC服务
	err := pb.RegisterRestServiceHandlerFromEndpoint(ctx, mux, "localhost:5000", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}

	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
