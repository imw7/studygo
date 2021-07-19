module server

go 1.16

require (
	github.com/moby/moby v20.10.7+incompatible
	google.golang.org/grpc v1.39.0
	imw7.com/pubsub/pb v0.0.0-00010101000000-000000000000
)

replace imw7.com/pubsub/pb => ../pb
