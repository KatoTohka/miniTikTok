package service

import (
	"context"
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

	if req.ActionType == 1 {
		if err := db.SetFavorite(context.Background(), &favorite); err != nil {
			return err
		}
		if err := db.AddFavoriteById(context.Background(), req.VideoId); err != nil {
			//没有考虑 此行为失败后对 favorite 数据的处理
			return err
		}
		return nil
	} else {
		if err := db.CancelFavorite(context.Background(), &favorite); err != nil {
			return err
		}
		if err := db.DecFavoriteById(context.Background(), req.VideoId); err != nil {
			//没有考虑 此行为失败后对 favorite 数据的处理
			return err
		}
		return nil
	}
}
