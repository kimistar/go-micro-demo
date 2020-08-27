package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"go-micro-demo/gin/core"
	"go-micro-demo/proto/greeter"
	"go-micro-demo/wrapper/hystrix"
)

var greeterService greeter.HelloService

func init() {
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://127.0.0.1:2379"}
	})
	srv := grpc.NewService(
		service.Registry(reg),
	)
	srv.Init()

	cl := srv.Client()
	if err := cl.Init(client.Retries(3)); err != nil {
		panic(err)
	}

	// hystrix.NewClientWrapper 熔断器
	greeterService = greeter.NewHelloService("go.micro.srv.greeter", hystrix.NewClientWrapper()(cl))
}

type SayHelloRequest struct {
	Name string `json:"name" binding:"required"`
}

func SayHello(c *core.Context) {
	var r SayHelloRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.Fail(40001, "缺少参数", nil)
		return
	}

	helloRsp, err := greeterService.Hello(c.Request.Context(), &greeter.HelloRequest{
		Name: r.Name,
	})
	if err != nil {
		fmt.Println(err)
		c.Fail(40000, "网络异常", nil)
		return
	}

	hiRsp, _ := greeterService.Hi(c.Request.Context(), &greeter.HiRequest{
		Name: r.Name,
	})

	c.Success(gin.H{
		"msg":   helloRsp.Msg,
		"reply": hiRsp.Reply,
	})
}
