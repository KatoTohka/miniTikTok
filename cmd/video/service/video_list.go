package service

import (
	"context"
	"miniTikTok/cmd/video/dal/db"
	"miniTikTok/kitex_gen/video"
	"miniTikTok/pkg/errno"
)

type VideoListService struct {
	ctx context.Context
}

func NewVideoListService(ctx context.Context) *VideoListService {
	return &VideoListService{
		ctx: ctx,
	}
}

func (s *VideoListService) ListVideo(req *video.DouyinPublishListRequest) ([]*video.Video, error) {
	// query user info
	users, err := db.QueryUserByID(s.ctx, req.UserId)
	if err != nil {
		return []*video.Video{}, errno.ServiceErr
	}
	if len(users) == 0 {
		return []*video.Video{}, errno.UserNotExistErr
	}
	var u video.User
	u.Id = int64(users[0].ID)
	u.Name = users[0].UserName
	u.FollowCount = &users[0].FollowCount
	u.FollowerCount = &users[0].FollowerCount
	u.IsFollow = false
	if err != nil {
		return []*video.Video{}, err
	}

	// query video info
	videos, err := db.QueryVideoById(s.ctx, u.Id)
	if err != nil {
		return []*video.Video{}, err
	}
	fVideo, _ := db.QueryFavoriteById(context.Background(), req.UserId)
	//TODO:效率较低
	favoriteList := make([]bool, len(videos))
	for i, v := range videos {
		if in(int64(v.ID), fVideo) {
			favoriteList[i] = true
		} else {
			favoriteList[i] = false
		}
	}
	var resp []*video.Video
	for i, v := range videos {
		vd := video.Video{
			Id:            int64(v.ID),
			Author:        &u,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    favoriteList[i],
			Title:         v.Title,
		}
		resp = append(resp, &vd)
	}
	return resp, nil
}

// util
func in(target int64, str_array []int64) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}
