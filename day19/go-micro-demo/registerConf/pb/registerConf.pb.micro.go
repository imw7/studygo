// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: registerConf.proto

package pb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RegisterConf service

func NewRegisterConfEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RegisterConf service

type RegisterConfService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (RegisterConf_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (RegisterConf_PingPongService, error)
}

type registerConfService struct {
	c    client.Client
	name string
}

func NewRegisterConfService(name string, c client.Client) RegisterConfService {
	return &registerConfService{
		c:    c,
		name: name,
	}
}

func (c *registerConfService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RegisterConf.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerConfService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (RegisterConf_StreamService, error) {
	req := c.c.NewRequest(c.name, "RegisterConf.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &registerConfServiceStream{stream}, nil
}

type RegisterConf_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*StreamingResponse, error)
}

type registerConfServiceStream struct {
	stream client.Stream
}

func (x *registerConfServiceStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *registerConfServiceStream) Close() error {
	return x.stream.Close()
}

func (x *registerConfServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *registerConfServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *registerConfServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *registerConfServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registerConfService) PingPong(ctx context.Context, opts ...client.CallOption) (RegisterConf_PingPongService, error) {
	req := c.c.NewRequest(c.name, "RegisterConf.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &registerConfServicePingPong{stream}, nil
}

type RegisterConf_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type registerConfServicePingPong struct {
	stream client.Stream
}

func (x *registerConfServicePingPong) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *registerConfServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *registerConfServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *registerConfServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *registerConfServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *registerConfServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *registerConfServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RegisterConf service

type RegisterConfHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, RegisterConf_StreamStream) error
	PingPong(context.Context, RegisterConf_PingPongStream) error
}

func RegisterRegisterConfHandler(s server.Server, hdlr RegisterConfHandler, opts ...server.HandlerOption) error {
	type registerConf interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type RegisterConf struct {
		registerConf
	}
	h := &registerConfHandler{hdlr}
	return s.Handle(s.NewHandler(&RegisterConf{h}, opts...))
}

type registerConfHandler struct {
	RegisterConfHandler
}

func (h *registerConfHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.RegisterConfHandler.Call(ctx, in, out)
}

func (h *registerConfHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RegisterConfHandler.Stream(ctx, m, &registerConfStreamStream{stream})
}

type RegisterConf_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type registerConfStreamStream struct {
	stream server.Stream
}

func (x *registerConfStreamStream) Close() error {
	return x.stream.Close()
}

func (x *registerConfStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *registerConfStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *registerConfStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *registerConfStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *registerConfHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.RegisterConfHandler.PingPong(ctx, &registerConfPingPongStream{stream})
}

type RegisterConf_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type registerConfPingPongStream struct {
	stream server.Stream
}

func (x *registerConfPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *registerConfPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *registerConfPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *registerConfPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *registerConfPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *registerConfPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
