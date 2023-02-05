package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/feed"
	"miniTikTok/pkg/errno"
	"time"
)

func convertVideoType2(videoList []*feed.Video) (resp []Video) {
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

func Feed(ctx context.Context, c *app.RequestContext) {
	hlog.Info("feed log")
	var feedVar FeedParam
	var err error
	if err = c.Bind(&feedVar); err != nil {
		SendFeedResponse(c, []Video{}, errno.ParamErr)
		return
	}
	feedTime := time.Now().Unix()
	//hlog.Info(feedVar.LatestTime)
	if len(feedVar.LatestTime) != 0 {
		ptime, err := time.ParseInLocation("2006-01-02 15:04:05", feedVar.LatestTime, time.Local)
		//if err != nil {
		//SendFeedResponse(c, []Video{}, errno.NewErrNo(10088, "time parse err"))
		//return
		//}
		if err == nil {
			feedTime = ptime.Unix()
		}
	}
	//log.Println(feedTime)
	videos, err := rpc.Feed(context.Background(), &feed.DouyinFeedRequest{
		LatestTime: &feedTime,
		Token:      nil,
	})
	if err != nil {
		SendFeedResponse(c, []Video{}, errno.ConvertErr(err))
		return
	}
	SendFeedResponse(c, convertVideoType2(videos), errno.Success)
}
