module server

go 1.16

require (
	google.golang.org/grpc v1.39.0
	imw7.com/grpc_stream/pb v0.0.0
)

replace imw7.com/grpc_stream/pb => ../pb
