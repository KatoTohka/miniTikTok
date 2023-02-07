package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"miniTikTok/kitex_gen/comment"
	"miniTikTok/kitex_gen/comment/commentservice"
	"miniTikTok/pkg/constants"
	"miniTikTok/pkg/errno"
	"time"
)

var commentClient commentservice.Client

func initCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := commentservice.NewClient(
		constants.CommentServiceName,
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
	commentClient = c
}
func Comment(ctx context.Context, req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	resp, err := commentClient.PublishComment(ctx, req)
	if err != nil {
		return &comment.Comment{}, err
	}
	if resp.StatusCode != 0 {
		return &comment.Comment{}, errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return resp.Comment, nil
}
func CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) ([]*comment.Comment, error) {
	resp, err := commentClient.ListComment(ctx, req)
	if err != nil {
		return []*comment.Comment{}, err
	}
	if resp.StatusCode != 0 {
		return []*comment.Comment{}, errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return resp.CommentList, nil

}
