package db

import (
	"context"
	"gorm.io/gorm"
)

// TXFavorite 开始事务点赞
func TXFavorite(ctx context.Context, videoId int64, favorite *Favorite) error {
	update := map[string]interface{}{
		"status": 1,
	}
	video := Video{}
	video.ID = uint(videoId)
	res := make([]Favorite, 0)
	if err := DB.Model(&favorite).Where(Favorite{UserId: favorite.UserId, VideoId: favorite.VideoId}).Find(&res).Error; err != nil {
		return err
	}
	// 点过赞了
	if len(res) != 0 && res[0].Status {
		return nil
	}
	// 需要点赞
	tx := DB.Begin()
	if err := tx.Model(&favorite).Where(Favorite{UserId: favorite.UserId, VideoId: favorite.VideoId}).Assign(update).FirstOrCreate(&Favorite{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//Favorite Count + 1
	if err := tx.WithContext(ctx).Model(&video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// TXCancelFavorite 开启事务取消点赞
func TXCancelFavorite(ctx context.Context, videoId int64, favorite *Favorite) error {
	update := map[string]interface{}{
		"status": 0,
	}
	video := Video{}
	video.ID = uint(videoId)
	res := make([]Favorite, 0)
	if err := DB.Model(&favorite).Where(Favorite{UserId: favorite.UserId, VideoId: favorite.VideoId}).Find(&res).Error; err != nil {
		return err
	}
	// 没赞过
	if len(res) == 0 || !res[0].Status {
		return nil
	}
	// 需要取消点赞
	tx := DB.Begin()
	if err := tx.Model(&favorite).Where(Favorite{UserId: favorite.UserId, VideoId: favorite.VideoId}).Assign(update).FirstOrCreate(&Favorite{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := DB.WithContext(ctx).Model(&video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
