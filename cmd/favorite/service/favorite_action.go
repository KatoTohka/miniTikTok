package service

import (
	"context"
	"errors"
	"log"
	"miniTikTok/cmd/favorite/dal/db"
	"miniTikTok/cmd/favorite/dal/redisDb"
	"miniTikTok/kitex_gen/favorite"
	"miniTikTok/middleware"
	"time"
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
		if err := FavoriteVideo(&favorite, req); err != nil {
			return err
		}
		return nil
	} else if req.ActionType == 2 {
		if err := CancelFavoriteVideo(&favorite, req.VideoId); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("param error")
	}
}

func FavoriteVideo(favorite *db.Favorite, req *favorite.DouyinFavoriteActionRequest) error {
	_, err := redisDb.RDBSetFavorite(req.VideoId, favorite.UserId)
	if err != nil {
		log.Println("redis save failure")
	}
	// 先写mysql
	// 开启一个事务保证原子性,点赞+数量
	if err = db.TXFavorite(context.Background(), req.VideoId, favorite); err != nil {
		return err
	}
	// 再写redis
	go func() {
		time.Sleep(time.Duration(500) * time.Millisecond)
		_, err = redisDb.RDBSetFavorite(req.VideoId, favorite.UserId)
		if err != nil {
			log.Println("redis save failure")
		}
	}()
	return nil
}

func CancelFavoriteVideo(favorite *db.Favorite, Id int64) error {
	// 延迟双删
	// 先删一次redis
	_, err := redisDb.RDBCancelFavorite(Id, favorite.UserId)
	if err != nil {
		log.Println("redis rem failure")
	}
	// 删mysql
	// 开启一个事务保证原子性,取消点赞+减数量
	if err := db.TXCancelFavorite(context.Background(), Id, favorite); err != nil {
		return err
	}
	// TODO：可以参考raft的replicator
	go func() {
		time.Sleep(time.Duration(500) * time.Millisecond)
		// 再删redis
		_, err = redisDb.RDBCancelFavorite(Id, favorite.UserId)
		if err != nil {
			log.Println("redis rem failure")
		}
		log.Println("第二次删redis")
	}()
	return nil
}
