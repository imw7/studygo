package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"imw7.com/grpc-web/pb"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 证书认证
/*
1.根证书 /conf/
生成key: openssl genrsa -out ca.key 2048
生成密钥: openssl req -new -x509 -days 3650 -key ca.key -out ca.pem
填写信息:
	Country Name (2 letter code) []:
	State or Province Name (full name) []:
	Locality Name (eg, city) []:
	Organization Name (eg, company) []:
	Organizational Unit Name (eg, section) []:
	Common Name (eg, fully qualified host name) []:localhost
	Email Address []:

2.Server /conf/server/
生成key: openssl ecparam -genkey -name secp385r1 -out server.key
生成CSR: openssl req -new -key server.key -out server.csr
填写信息: ...
基于CA签发: openssl x509 -req -sha256 -CA ../ca.pem -CAkey ../ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem

3.Client /conf/client/
生成key: openssl ecparam -genkey -name secp385r1 -out client.key
生成CSR: openssl req -new -key client.key -out client.csr
填写信息: ...
基于CA签发: openssl x509 -req -sha256 -CA ../ca.pem -CAkey ../ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem

目录结构:
$ tree conf
conf
├── ca.key
├── ca.pem
├── ca.srl
├── client
│   ├── client.csr
│   ├── client.key
│   └── client.pem
└── server
    ├── server.csr
    ├── server.key
    └── server.pem
*/

type HelloService struct{}

func (p *HelloService) Hello(ctx context.Context, req *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "Hello, " + req.GetValue() + "!"}
	return reply, nil
}

func main() {
	certificate, err := tls.LoadX509KeyPair("../conf/server/server.pem", "../conf/server/server.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../conf/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile("./index.html")
		if err != nil {
			_, _ = w.Write([]byte(fmt.Sprintf("%v\n", err)))
		}
		_, _ = w.Write(b)
	})
	h2Handler := h2c.NewHandler(mux, &http2.Server{})
	server := &http.Server{Addr: ":8972", Handler: h2Handler}
	_ = server.ListenAndServe()

	err = http.ListenAndServeTLS(":8972", "../conf/server/server.pem", "../conf/server/server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
				return
			}
			mux.ServeHTTP(w, r)
			return
		}))

	if err != nil {
		log.Fatal(err)
	}
}
