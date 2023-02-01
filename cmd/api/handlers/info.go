package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/middleware"
	"miniTikTok/pkg/errno"
	"strconv"
)

func Info(ctx context.Context, c *app.RequestContext) {
	var infoVar UserInfoParam
	if err := c.Bind(&infoVar); err != nil {
		SendInfoResponse(c, errno.ConvertErr(err), User{})
		return
	}
	if len(infoVar.ID) == 0 || len(infoVar.Token) == 0 {
		SendInfoResponse(c, errno.AuthorizationFailedErr, User{})
		return
	}
	//JWT
	token, _, err := middleware.ParseToken(infoVar.Token)
	if err != nil || !token.Valid {
		SendInfoResponse(c, errno.AuthorizationFailedErr, User{})
		return
	}

	uId, _ := strconv.ParseInt(infoVar.ID, 10, 64)
	infoUser, err := rpc.InfoUser(context.Background(), &user.DouyinUserRequest{
		UserId: uId,
		Token:  infoVar.Token,
	})
	if err != nil {
		SendInfoResponse(c, errno.ConvertErr(err), User{})
		return
	}
	SendInfoResponse(c, errno.Success, User{
		ID:            infoUser.Id,
		UserName:      infoUser.Name,
		FollowCount:   *infoUser.FollowCount,
		FollowerCount: *infoUser.FollowerCount,
		IsFollow:      infoUser.IsFollow,
	})
}
