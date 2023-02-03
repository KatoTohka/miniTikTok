package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"miniTikTok/kitex_gen/video"
	"miniTikTok/kitex_gen/video/videoservice"
	"miniTikTok/pkg/constants"
	"miniTikTok/pkg/errno"
	"time"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(
		constants.VideoServiceName,
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
	videoClient = c

}

func PublishVideo(ctx context.Context, req *video.DouyinPublishActionRequest) error {
	resp, err := videoClient.PublishVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return nil
}

func ListVideo(ctx context.Context, req *video.DouyinPublishListRequest) ([]*video.Video, error) {
	resp, err := videoClient.ListVideo(ctx, req)
	if err != nil {
		return []*video.Video{}, err
	}
	if resp.StatusCode != 0 {
		return []*video.Video{}, errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return resp.VideoList, nil
}
