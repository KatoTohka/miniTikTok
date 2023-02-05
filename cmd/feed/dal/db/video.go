package db

import (
	"context"
	"gorm.io/gorm"
	"miniTikTok/pkg/constants"
	"time"
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

// QueryVideoBeforeTime query video info by author id
func QueryVideoBeforeTime(ctx context.Context, t time.Time, limit int) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("created_at < ?", t).Order("created_at DESC").Limit(limit).Find(&res).Error; err != nil {
		return []*Video{}, err
	}
	return res, nil
}
