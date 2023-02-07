package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"miniTikTok/cmd/comment/dal"
	comment "miniTikTok/kitex_gen/comment/commentservice"
	"miniTikTok/pkg/constants"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		log.Fatal(err, "etcd registry error")
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.CommentServerHost)
	if err != nil {
		log.Fatal(err, "net resolve error")
	}
	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CommentServiceName}), // server name
		//server.WithMiddleware(middleware.CommonMiddleware),                                             // middleware
		//server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		// server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		//server.WithBoundHandler(bound.NewCpuLimitHandler()), // BoundHandler
		server.WithRegistry(r), // registry
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
