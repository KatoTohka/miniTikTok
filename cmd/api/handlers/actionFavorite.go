package handlers

import (
	"context"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/favorite"
	"miniTikTok/middleware"
	"miniTikTok/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	hlog.Info("favorite action log")
	var favoriteActionVar FavoriteActionParam
	if err := c.Bind(&favoriteActionVar); err != nil {
		SendFavoriteActionResponse(c, errno.ParamErr)
		return
	}
	if len(favoriteActionVar.Token) == 0 || len(favoriteActionVar.ActionType) == 0 || len(favoriteActionVar.VideoID) == 0 {
		SendFavoriteActionResponse(c, errno.ParamErr)
		return
	}
	videoId, _ := strconv.ParseInt(favoriteActionVar.VideoID, 10, 64)
	actionType, _ := strconv.ParseInt(favoriteActionVar.ActionType, 10, 64)
	if actionType != 1 && actionType != 2 {
		SendFavoriteActionResponse(c, errno.ParamErr)
		return
	}
	//JWT
	token, _, err := middleware.ParseToken(favoriteActionVar.Token)
	if err != nil || !token.Valid {
		SendFavoriteActionResponse(c, errno.AuthorizationFailedErr)
		return
	}
	err = rpc.ActionFavorite(context.Background(), &favorite.DouyinFavoriteActionRequest{
		Token:      favoriteActionVar.Token,
		VideoId:    videoId,
		ActionType: int32(actionType),
	})
	if err != nil {
		SendFavoriteActionResponse(c, err)
		return
	}
	SendFavoriteActionResponse(c, errno.Success)
}
