package main

import (
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	opentrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/go-plugins/wrapper/validator/v2"
	"github.com/opentracing/opentracing-go"
	"go-micro-demo/proto/greeter"
	"go-micro-demo/service/handler"
	"go-micro-demo/wrapper/tracer"
	"time"
)

func main() {
	t, closer, err := tracer.NewTracer("go.micro.srv.greeter", "127.0.0.1:6831")
	if err != nil {
		logger.Fatal(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)

	reg := etcd.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{"http://127.0.0.1:2379"}
		},
		//etcd.Auth("username", "password"),
	)
	srv := grpc.NewService(
		service.Name("go.micro.srv.greeter"),
		service.Version("latest"),
		service.Registry(reg),
		service.RegisterTTL(30*time.Second),
		service.RegisterInterval(15*time.Second),
		service.WrapHandler(validator.NewHandlerWrapper()),
		service.WrapHandler(opentrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	srv.Init()

	if err := greeter.RegisterHelloHandler(srv.Server(), new(handler.Say)); err != nil {
		panic(err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
