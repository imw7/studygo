package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"imw7.com/tls/pb"
	"io/ioutil"
	"log"
)

func main() {
	certificate, err := tls.LoadX509KeyPair("../conf/client/client.pem", "../conf/client/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../conf/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "e14_tls", // NOTE: this is required!
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":8972", grpc.WithTransportCredentials(creds))
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
