package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"golang.org/x/net/context"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/kitex_gen/user/userservice"
	"miniTikTok/pkg/constants"
	"miniTikTok/pkg/errno"
	"time"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		constants.UserServiceName,
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
	userClient = c

}

// register user rpc

// CreateUser return id token error
func CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (int64, string, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return -1, "", err
	}
	if resp.StatusCode != 0 {
		return -1, "", errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return resp.UserId, resp.Token, nil
}

func CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (int64, string, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return -1, "", err
	}
	if resp.StatusCode != 0 {
		return -1, "", errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return resp.UserId, resp.Token, nil
}
func InfoUser(ctx context.Context, req *user.DouyinUserRequest) (*user.User, error) {
	resp, err := userClient.QueryUser(ctx, req)
	if err != nil {
		return &user.User{}, err
	}
	if resp.StatusCode != 0 {
		return &user.User{}, errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}
	return resp.User, nil

}
