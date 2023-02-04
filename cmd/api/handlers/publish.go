package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/video"
	"miniTikTok/middleware"
	"miniTikTok/pkg/errno"
)

func Publish(ctx context.Context, c *app.RequestContext) {
	hlog.Info("publish log")
	var publishVar PublishParam
	if err := c.Bind(&publishVar); err != nil {
		SendPublishResponse(c, errno.ParamErr)
		return
	}
	if len(publishVar.Title) == 0 || len(publishVar.Token) == 0 {
		SendPublishResponse(c, errno.ParamErr)
		return
	}
	if publishVar.Data == nil {
		SendPublishResponse(c, errno.ParamErr)
		return
	}
	//TODO:可能需要抽象到service层
	videoData, err := publishVar.Data.Open()
	if err != nil {
		SendPublishResponse(c, errno.PublishErr)
		return
	}
	byteContainer, err := io.ReadAll(videoData)
	if err != nil {
		SendPublishResponse(c, errno.PublishErr)
		return
	}
	//log.Print(byteContainer)
	//JWT
	token, _, err := middleware.ParseToken(publishVar.Token)
	if err != nil || !token.Valid {
		SendPublishResponse(c, errno.AuthorizationFailedErr)
		return
	}
	err = rpc.PublishVideo(context.Background(), &video.DouyinPublishActionRequest{
		Token: publishVar.Token,
		Data:  byteContainer,
		Title: publishVar.Title,
	})
	if err != nil {
		SendPublishResponse(c, err)
		return
	}
	SendPublishResponse(c, errno.Success)
}
