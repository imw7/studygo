module web-server

go 1.16

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	google.golang.org/grpc v1.39.0
	imw7.com/grpc_rest/pb v0.0.0
)

replace imw7.com/grpc_rest/pb => ../pb
