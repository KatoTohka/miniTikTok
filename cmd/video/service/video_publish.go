package service

import (
	"context"
	"fmt"
	"log"
	"miniTikTok/cmd/video/dal/cache"
	"miniTikTok/cmd/video/dal/db"
	"miniTikTok/cmd/video/dal/tos"
	"miniTikTok/kitex_gen/video"
	"miniTikTok/middleware"
	"miniTikTok/pkg/constants"
	"miniTikTok/pkg/snapshot"
	"miniTikTok/pkg/snowflake"
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(2)
	data := req.Data
	snow, _ := snowflake.NewSnowflake(0)
	videoId := snow.Generate()
	tosKey := fmt.Sprintf("%v.mp4", videoId)
	var err error
	go func(err error) {
		// 协程处理tos上传video
		defer wg.Done()
		err = tos.Upload(data, tosKey)
	}(err)
	playUrl := constants.TosURL + tosKey
	go func(err error) {
		// 协程处理视频帧
		defer wg.Done()
		err = cache.Cache(data, int64(videoId))
		if err != nil {
			return
		}
		imgData, err := snapshot.GetSnapshot(constants.TempCachePlace+fmt.Sprintf("%v.mp4", videoId), 1, fmt.Sprintf("%v", videoId))
		if err != nil {
			return
		}
		imgTosKey := fmt.Sprintf("%v.png", videoId)
		log.Println(imgTosKey)
		err = tos.Upload(imgData, imgTosKey)
		if err != nil {
			return
		}
	}(err)
	imgTosKey := fmt.Sprintf("%v.png", videoId)
	imgUrl := constants.TosURL + imgTosKey
	wg.Wait()
	// token already verified in api
	_, claim, _ := middleware.ParseToken(req.Token)
	userId := claim.UserID
	v := db.Video{
		PlayUrl:       playUrl,
		CoverUrl:      imgUrl,
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
