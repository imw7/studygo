module server

go 1.16

require (
	google.golang.org/grpc v1.39.0
	imw7.com/interceptor/pb v1.0.0
)

replace imw7.com/interceptor/pb => ../pb
