package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/pkg/errno"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam
	if err := c.Bind(&registerVar); err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendRegisterResponse(c, errno.ParamErr, -1, "")
		return
	}

	userID, token, err := rpc.CreateUser(context.Background(), &user.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})

	if err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	SendRegisterResponse(c, errno.Success, userID, token)
}
