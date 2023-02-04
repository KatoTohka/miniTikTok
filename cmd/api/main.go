package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	"miniTikTok/cmd/api/handlers"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/pkg/tracer"
	"os"
)

func Init() {
	rpc.InitRPC()
	tracer.InitJaeger("api")

}
func main() {
	Init()
	logger := hertzZerolog.New(
		hertzZerolog.WithOutput(os.Stdout),     // allows to specify output
		hertzZerolog.WithLevel(hlog.LevelInfo), // option with log level
		hertzZerolog.WithTimestamp(),           // option with timestamp
		hertzZerolog.WithCaller(),              // option with caller
	)
	hlog.SetLogger(logger)

	r := server.New(
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithHandleMethodNotAllowed(true),
	)

	douyin := r.Group("/douyin")
	userGroup := douyin.Group("/user")
	userGroup.POST("/login/", handlers.Login)
	userGroup.POST("/register/", handlers.Register)
	userGroup.GET("/", handlers.Info)
	publishGroup := douyin.Group("/publish")
	publishGroup.POST("/action/", handlers.Publish)
	publishGroup.GET("/list/", handlers.VideoList)
	feedGroup := douyin.Group("/feed/")
	feedGroup.GET("/", handlers.Feed)
	r.Spin()
}
