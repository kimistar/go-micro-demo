package main

import (
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"github.com/opentracing/opentracing-go"
	"go-micro-demo/gin/router"
	"go-micro-demo/wrapper/tracer"
	"time"
)

func main() {
	t, closer, err := tracer.NewTracer("api.gateway", "127.0.0.1:6831")
	if err != nil {
		logger.Fatal(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)

	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://127.0.0.1:2379"}
	})

	service := web.NewService(
		web.Name("api.gateway"),
		web.Version("latest"),
		web.Address(":8888"),
		web.Handler(router.Register()),
		web.Registry(reg),
		web.RegisterTTL(30*time.Second),
		web.RegisterInterval(15*time.Second),
	)

	if err := service.Run(); err != nil {
		panic(nil)
	}
}
