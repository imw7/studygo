module server

go 1.21

require (
	github.com/moby/moby v20.10.7+incompatible
	google.golang.org/grpc v1.59.0
	imw7.com/pubsub/pb v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231127180814-3a041ad873d4 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace imw7.com/pubsub/pb => ../pb
