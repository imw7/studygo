package main

import (
	"context"
	"go-micro-demo/pubsub/pb"
	"go-micro.dev/v4"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/util/log"
)

// Sub All methods of Sub will be executed when a message is received
type Sub struct{}

// Process Method can be of any name
func (s *Sub) Process(ctx context.Context, event *pb.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1] received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// Alternatively a function can be used
func subEv(ctx context.Context, event *pb.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.2] received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {
	// create a service
	service := micro.NewService(
		micro.Name("go.micro.srv.pubsub"),
		micro.Version("latest"),
	)
	// initialise flags
	service.Init()

	// register subscriber
	if err := micro.RegisterSubscriber("example.topic.pubsub.1", service.Server(), new(Sub)); err != nil {
		log.Fatal(err)
	}

	// register subscriber with queue, each msg is delivered to a unique subscriber
	if err := micro.RegisterSubscriber("example.topic.pubsub.2", service.Server(), subEv, server.SubscriberQueue("queue.pubsub")); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
