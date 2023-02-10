package service

import (
	"context"
	"fmt"
	"miniTikTok/cmd/favorite/dal/db"
	"miniTikTok/kitex_gen/favorite"
	"miniTikTok/pkg/errno"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

func (s *FavoriteListService) ListFavorite(req *favorite.DouyinFavoriteListRequest) ([]*favorite.Video, error) {
	// query user favorite
	videoIds, err := db.QueryFavoriteById(s.ctx, req.UserId)
	if err != nil {
		return []*favorite.Video{}, errno.ServiceErr
	}
	if len(videoIds) == 0 {
		return []*favorite.Video{}, nil
	}
	videos := make([]*favorite.Video, 0, len(videoIds))
	for _, vid := range videoIds {
		tmpVideo, err := db.QueryVideoByVId(context.Background(), vid)
		if err != nil {
			return videos, errno.ServiceErr
		}
		dbAuthor, err := db.QueryUserByID(context.Background(), tmpVideo.AuthorId)
		FavAuthor := &favorite.User{
			Id:            int64(dbAuthor.ID),
			Name:          dbAuthor.UserName,
			FollowCount:   dbAuthor.FollowCount,
			FollowerCount: dbAuthor.FollowerCount,
			IsFollow:      false, // 关注功能未实现
		}
		if err != nil {
			return videos, errno.ServiceErr
		}
		videos = append(videos, &favorite.Video{
			Id:            int64(tmpVideo.ID),
			Author:        FavAuthor,
			PlayUrl:       tmpVideo.PlayUrl,
			CoverUrl:      tmpVideo.CoverUrl,
			FavoriteCount: tmpVideo.FavoriteCount,
			CommentCount:  tmpVideo.CommentCount,
			IsFavorite:    true,
			Title:         tmpVideo.Title,
		})
	}
	ids := make([]int64, 0)
	for _, v := range videos {
		ids = append(ids, v.Id)
	}
	fmt.Println(ids)
	return videos, nil
}
