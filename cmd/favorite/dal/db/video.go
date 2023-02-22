package db

import (
	"context"
	"fmt"
	"miniTikTok/pkg/constants"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	PlayUrl       string `gorm:"column:play_url;type:varchar(255);not null"`
	CoverUrl      string `gorm:"column:cover_url;type:varchar(255);not null"`
	FavoriteCount int64  `gorm:"column:favorite_count;type:bigint;not null"`
	CommentCount  int64  `gorm:"column:comment_count;type:bigint;not null"`
	Title         string `gorm:"column:title;type:varchar(255);not null"`
	AuthorId      int64  `gorm:"column:author_id;type:bigint;not null"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

//// AddFavoriteById add video favorite count by vedio id
//func AddFavoriteById(ctx context.Context, videoId int64) error {
//	video := Video{}
//	video.ID = uint(videoId)
//	//Favorite Count + 1
//	if err := DB.WithContext(ctx).Model(&video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
//		return err
//	}
//	return nil
//}

// DecFavoriteById decrease video favorite count by vedio id
func DecFavoriteById(ctx context.Context, videoId int64) error {
	video := Video{}
	video.ID = uint(videoId)
	//Favorite Count - 1 (没考虑favorite count 为0时-1的情况)
	if err := DB.WithContext(ctx).Model(&video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// QueryVideoByVId query video by VideoId
func QueryVideoByVId(ctx context.Context, videoId int64) (*Video, error) {
	video := Video{}
	if err := DB.WithContext(ctx).Where("id = ?", fmt.Sprint(videoId)).First(&video).Error; err != nil {
		return &video, err
	}
	return &video, nil
}
