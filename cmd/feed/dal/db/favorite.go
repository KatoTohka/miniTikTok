package db

import (
	"context"
	"fmt"
	"miniTikTok/pkg/constants"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
	Status  bool  `json:"status"`
}

func (f *Favorite) TableName() string {
	return constants.UsersFavoriteTableName
}

// QueryFavoriteStatus query IsFavorite
func QueryFavoriteStatus(ctx context.Context, userID int64, videoID int64) (bool, error) {
	res := Favorite{}
	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", fmt.Sprint(userID), fmt.Sprint(videoID)).Limit(1).Find(&res).Error; err != nil {
		return false, nil
	}
	return res.Status, nil
}
