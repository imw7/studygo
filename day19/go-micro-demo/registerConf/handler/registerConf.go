package handler

import (
	"context"
	"go-micro-demo/registerConf/pb"
	"go-micro.dev/v4/util/log"
)

type RegisterConf struct{}

// Call is a single request handler called via client.Call or the generated client code
func (r *RegisterConf) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Info("Received RegisterConf.Call request")
	rsp.Msg = "Hello, " + req.Name + "!"
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (r *RegisterConf) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.RegisterConf_StreamStream) error {
	log.Info("Received RegisterConf.Stream request")

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d\n", i)
		if err := stream.Send(&pb.StreamingResponse{Count: int64(i)}); err != nil {
			return err
		}
	}
	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (r *RegisterConf) PingPong(ctx context.Context, stream pb.RegisterConf_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v\n", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
