package rpc

import (
	"context"
	"miniTikTok/kitex_gen/favorite"
	"miniTikTok/kitex_gen/favorite/favoriteservice"
	"miniTikTok/pkg/constants"
	"miniTikTok/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
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
	favoriteClient = c

}

func ActionFavorite(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) error {
	resp, err := favoriteClient.ActionFavorite(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return nil
}

func ListFavorite(ctx context.Context, req *favorite.DouyinFavoriteListRequest) ([]*favorite.Video, error) {
	resp, err := favoriteClient.ListFavorite(ctx, req)
	if err != nil {
		return []*favorite.Video{}, err
	}
	if resp.StatusCode != 0 {
		return []*favorite.Video{}, errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return resp.VideoList, nil
}
