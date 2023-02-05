package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"miniTikTok/kitex_gen/feed"
	"miniTikTok/kitex_gen/feed/feedservice"
	"miniTikTok/pkg/constants"
	"miniTikTok/pkg/errno"
	"time"
)

var feedclient feedservice.Client

func initFeedRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := feedservice.NewClient(
		constants.FeedServiceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		panic(err)
	}
	feedclient = c
}

func Feed(ctx context.Context, req *feed.DouyinFeedRequest) ([]*feed.Video, error) {
	resp, err := feedclient.Feed(ctx, req)
	if err != nil {
		return []*feed.Video{}, err
	}
	if resp.StatusCode != 0 {
		return []*feed.Video{}, errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return resp.VideoList, nil
}
