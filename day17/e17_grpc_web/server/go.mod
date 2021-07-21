module server

go 1.16

require (
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/grpc v1.39.0
	imw7.com/grpc-web/pb v1.0.0
)

replace imw7.com/grpc-web/pb => ../pb
