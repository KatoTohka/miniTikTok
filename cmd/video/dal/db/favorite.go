package db

import (
	"context"
	"fmt"
	"miniTikTok/pkg/constants"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id" structs:"user_id"`
	VideoId int64 `json:"video_id" structs:"video_id"`
	Status  bool  `json:"status" structs:"status"`
}

func (f *Favorite) TableName() string {
	return constants.UsersFavoriteTableName
}

// QueryFavoriteById query favorite video info by user id
func QueryFavoriteById(ctx context.Context, userId int64) ([]int64, error) {
	favorites := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND status = ?", fmt.Sprint(userId), "1").Find(&favorites).Error; err != nil {
		return []int64{}, err
	}
	ret := make([]int64, len(favorites))
	for i, f := range favorites {
		ret[i] = f.VideoId
	}
	return ret, nil
}
