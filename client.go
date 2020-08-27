package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"go-micro-demo/proto/greeter"
)

func main() {
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://127.0.0.1:2379"}
	})
	srv := grpc.NewService(
		service.Registry(reg),
	)
	srv.Init()

	client := greeter.NewHelloService("go.micro.srv.greeter", srv.Client())
	rsp, err := client.Hello(context.Background(), &greeter.HelloRequest{
		Name: "kimi",
	})
	fmt.Println(rsp, err)
}
