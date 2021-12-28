package handler

import (
	"context"
	log "go-micro.dev/v4/logger"
	pb "micro_etcd/hello/proto"
)

type Hello struct{}

func (e *Hello) Call(_ context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Hello.Call request: %v", req)
	rsp.Msg = "Hello, " + req.Name + "!"
	return nil
}
