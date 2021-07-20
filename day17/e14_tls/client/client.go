package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"imw7.com/tls/pb"
	"log"
)

/* 客户端 生成私钥和证书
$ openssl genrsa -out client.key 2048
$ openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" \
    -key client.key -out client.crt
*/

func main() {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	client := pb.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &pb.String{Value: "Edward"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
