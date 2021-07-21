module client

go 1.16

require (
	google.golang.org/grpc v1.39.0
	imw7.com/grpc-web/pb v1.0.0
)

replace imw7.com/grpc-web/pb => ../pb
