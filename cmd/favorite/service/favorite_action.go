package service

import (
	"context"
	"errors"
	"miniTikTok/cmd/favorite/dal/db"
	"miniTikTok/kitex_gen/favorite"
	"miniTikTok/middleware"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteActionService) ActionFavorite(req *favorite.DouyinFavoriteActionRequest) error {
	_, claim, _ := middleware.ParseToken(req.Token)
	userId := claim.UserID
	favorite := db.Favorite{
		UserId:  userId,
		VideoId: req.VideoId,
	}
	videoList, _ := db.QueryFavoriteById(context.Background(), userId)
	isFavorite := false
	for _, videoid := range videoList {
		if videoid == req.VideoId {
			isFavorite = true
			break
		}
	}

	if req.ActionType == 1 {
		if isFavorite {
			return nil
		}
		if err := db.SetFavorite(context.Background(), &favorite); err != nil {
			return err
		}
		if err := db.AddFavoriteById(context.Background(), req.VideoId); err != nil {
			//没有考虑 此行为失败后对 favorite 数据的处理
			return err
		}
		return nil
	} else if req.ActionType == 2 {
		if err := db.CancelFavorite(context.Background(), &favorite); err != nil {
			return err
		}
		if err := db.DecFavoriteById(context.Background(), req.VideoId); err != nil {
			//没有考虑 此行为失败后对 favorite 数据的处理
			return err
		}
		return nil
	} else {
		return errors.New("param error")
	}
}
