package router

import (
	"github.com/gin-gonic/gin"
	"go-micro-demo/gin/core"
	"go-micro-demo/gin/handler"
	"go-micro-demo/gin/middleware"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(core.Handle(middleware.Tracer))
	r.POST("/say/hello", core.Handle(handler.SayHello))
	return r
}
