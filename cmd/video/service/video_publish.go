package service

import (
	"context"
	"miniTikTok/cmd/video/dal/db"
	"miniTikTok/cmd/video/dal/tos"
	"miniTikTok/kitex_gen/video"
	"miniTikTok/middleware"
	"miniTikTok/pkg/constants"
)

type VideoPublishService struct {
	ctx context.Context
}

func NewVideoPublishService(ctx context.Context) *VideoPublishService {
	return &VideoPublishService{
		ctx: ctx,
	}
}

func (s *VideoPublishService) PublishVideo(req *video.DouyinPublishActionRequest) error {
	title := req.Title + ".mp4"
	data := req.Data

	err := tos.Upload(data, title)
	if err != nil {
		return err
	}
	// token already verified in api
	_, claim, _ := middleware.ParseToken(req.Token)
	userId := claim.UserID
	//TODO:视频封面使用切片。。
	v := db.Video{
		PlayUrl:       constants.TosURL + title,
		CoverUrl:      "http://dn-lego-static.qbox.me/1650523002-overseabanner2880x704.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
		AuthorId:      userId,
	}
	err = db.PublishVideo(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}
