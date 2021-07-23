package main

import (
	"context"
	"fmt"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"grpc-web/pb"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// 使用cmux让grpc和http监听同一个端口
// 注意cmux不支持TLS

type HelloService struct{}

func (p *HelloService) Hello(_ context.Context, req *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "Hello, " + req.GetValue() + "!"}
	return reply, nil
}

func main() {
	// 创建监听
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)

	// 创建grpc监听器
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	// 其余的都是http监听
	httpListener := m.Match(cmux.Any())

	// 创建服务
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./index.html")
		if err != nil {
			_, err := w.Write([]byte(fmt.Sprintf("%v\n", err)))
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = w.Write(file)
		if err != nil {
			log.Fatal(err)
		}
	})
	httpServer := &http.Server{}

	// 使用error group启动所有服务
	g := errgroup.Group{}
	g.Go(func() error {
		return grpcServer.Serve(grpcListener)
	})
	g.Go(func() error {
		return httpServer.Serve(httpListener)
	})
	g.Go(func() error {
		return m.Serve()
	})

	// 等待执行，检查错误
	if err = g.Wait(); err != nil {
		log.Fatal(err)
	}
}
