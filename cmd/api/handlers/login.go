package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/pkg/errno"
)

func Login(ctx context.Context, c *app.RequestContext) {
	hlog.Info("login log")
	var loginVar UserParam
	if err := c.Bind(&loginVar); err != nil {
		SendLoginResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
		SendLoginResponse(c, errno.ParamErr, -1, "")
		return
	}
	userID, token, err := rpc.CheckUser(context.Background(), &user.DouyinUserLoginRequest{
		Username: loginVar.UserName,
		Password: loginVar.PassWord,
	})
	if err != nil {
		SendLoginResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	SendLoginResponse(c, errno.Success, userID, token)
}
