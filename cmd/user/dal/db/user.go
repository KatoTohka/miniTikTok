package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"miniTikTok/pkg/constants"
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

// RegisterUser  create user in db
func RegisterUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUserByID query user info
func QueryUserByID(ctx context.Context, userID int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", fmt.Sprint(userID)).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	//if err := DB.Where("id = ?", "1").Limit(1).Find(&res).Error; err != nil {
	//	panic("failed query")
	//}
	return res, nil
}
