module client

go 1.16

require (
	google.golang.org/grpc v1.39.0
	imw7.com/token/conf v1.0.0
	imw7.com/token/pb v1.0.0
)

replace (
	imw7.com/token/conf => ../conf
	imw7.com/token/pb => ../pb
)
