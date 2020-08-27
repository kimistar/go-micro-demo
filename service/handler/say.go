package handler

import (
	"context"
	"go-micro-demo/proto/greeter"
	"time"
)

type Say struct {
}

var _ greeter.HelloHandler = (*Say)(nil)

func (*Say) Hello(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	time.Sleep(2 * time.Second)
	rsp.Msg = "hello " + req.Name
	return nil
}

func (*Say) Hi(ctx context.Context, req *greeter.HiRequest, rsp *greeter.HiResponse) error {
	time.Sleep(2 * time.Second)
	rsp.Reply = "hello " + req.Name
	return nil
}
