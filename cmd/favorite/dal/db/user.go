package db

import (
	"context"
	"fmt"
	"miniTikTok/pkg/constants"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string `gorm:"column:username;unique;type:varchar(30);not null"`
	PassWord      string `gorm:"column:password;type:varchar(60);not null"`
	FollowCount   int64  `gorm:"column:follow_count;type:uint;not null"`   // 关注总数
	FollowerCount int64  `gorm:"column:follower_count;type:uint;not null"` // 粉丝总数
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// QueryUserByID query author Id
func QueryUserByID(ctx context.Context, authorID int64) (*User, error) {
	res := User{}
	if err := DB.WithContext(ctx).Where("id = ?", fmt.Sprint(authorID)).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	//if err := DB.Where("id = ?", "1").Limit(1).Find(&res).Error; err != nil {
	//	panic("failed query")
	//}
	return &res, nil
}
