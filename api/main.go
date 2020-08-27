package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro-demo/api/handler"
	"go-micro-demo/proto/greeter"
)

func main() {
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://127.0.0.1:2379"}
	})

	service := micro.NewService(
		micro.Name("go.micro.api.user"),
		micro.Version("latest"),
		micro.Registry(reg),
		//micro.Selector(selector.NewSelector(func(options *selector.Options) {
		//	options.Registry = reg
		//})),
	)
	service.Init()

	service.Server().Handle( //nolint
		service.Server().NewHandler(
			&handler.Say{Client: greeter.NewHelloService("go.micro.srv.greeter", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
