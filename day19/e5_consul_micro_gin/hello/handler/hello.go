package handler

import (
	pb "consul_micro_gin/hello/proto"
	"context"
	log "go-micro.dev/v4/logger"
)

type Hello struct{}

func (e *Hello) Call(_ context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Hello.Call request: %v", req)
	rsp.Msg = "Hello, " + req.Name + "!"
	return nil
}
