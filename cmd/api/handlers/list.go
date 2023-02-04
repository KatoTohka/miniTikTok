package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/video"
	"miniTikTok/pkg/errno"
	"strconv"
)

func convertVideoType(videoList []*video.Video) (resp []Video) {
	for _, v := range videoList {
		author := (*v).Author
		cAuthor := User{
			ID:            author.Id,
			UserName:      author.Name,
			FollowCount:   *author.FollowCount,
			FollowerCount: *author.FollowerCount,
			IsFollow:      false,
		}
		tv := Video{
			ID:            (*v).Id,
			Author:        cAuthor,
			PlayUrl:       (*v).PlayUrl,
			CoverUrl:      (*v).CoverUrl,
			FavoriteCount: (*v).FavoriteCount,
			CommentCount:  (*v).CommentCount,
			IsFavorite:    false,
			Title:         (*v).Title,
		}
		resp = append(resp, tv)
	}
	return resp
}

func VideoList(ctx context.Context, c *app.RequestContext) {
	hlog.Info("videoList log")
	var videoVar ListParam
	if err := c.Bind(&videoVar); err != nil {
		SendListResponse(c, []Video{}, errno.ParamErr)
		return
	}
	if len(videoVar.Token) == 0 || len(videoVar.ID) == 0 {
		SendListResponse(c, []Video{}, errno.ParamErr)
		return
	}
	uId, _ := strconv.ParseInt(videoVar.ID, 10, 64)
	videoList, err := rpc.ListVideo(context.Background(), &video.DouyinPublishListRequest{
		UserId: uId,
		Token:  videoVar.Token,
	})
	if err != nil {
		SendListResponse(c, []Video{}, errno.ConvertErr(err))
		return
	}
	SendListResponse(c, convertVideoType(videoList), errno.Success)
}
