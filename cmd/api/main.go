package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"miniTikTok/cmd/api/handlers"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/pkg/tracer"
)

func Init() {
	rpc.InitRPC()
	tracer.InitJaeger("api")

}
func main() {
	Init()
	r := server.New(
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithHandleMethodNotAllowed(true),
	)
	douyin := r.Group("/douyin")
	userGroup := douyin.Group("/user")
	userGroup.POST("/login/", handlers.Login)
	userGroup.POST("/register/", handlers.Register)
	userGroup.GET("/", handlers.Info)
	r.Spin()
}
