package handlers

import (
	"context"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/favorite"
	"miniTikTok/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func convertFavoriteVideoType(videoList []*favorite.Video) (resp []Video) {
	for _, v := range videoList {
		author := v.Author
		cAuthor := User{
			ID:            author.Id,
			UserName:      author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      false,
		}
		tv := Video{
			ID:            v.Id,
			Author:        cAuthor,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    true,
			Title:         v.Title,
		}
		resp = append(resp, tv)
	}
	return resp
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {
	hlog.Info("favoriteList log")
	var favoriteListVar FavoriteListParam
	if err := c.Bind(&favoriteListVar); err != nil {
		SendFavoriteListResponse(c, []Video{}, errno.ParamErr)
		return
	}
	if len(favoriteListVar.ID) == 0 || len(favoriteListVar.Token) == 0 {
		SendFavoriteListResponse(c, []Video{}, errno.ParamErr)
		return
	}
	//JWT ?
	uId, _ := strconv.ParseInt(favoriteListVar.ID, 10, 64)
	favoriteList, err := rpc.ListFavorite(context.Background(), &favorite.DouyinFavoriteListRequest{
		UserId: uId,
		Token:  favoriteListVar.Token,
	})
	if err != nil {
		SendFavoriteListResponse(c, []Video{}, errno.ConvertErr(err))
		return
	}
	SendFavoriteListResponse(c, convertFavoriteVideoType(favoriteList), errno.Success)
}
